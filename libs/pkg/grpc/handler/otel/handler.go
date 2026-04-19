package otel

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

// ref: https://github.com/bakins/otel-grpc-statshandler/blob/main/statshandler.go
// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/google.golang.org/grpc/otelgrpc/config.go
// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/google.golang.org/grpc/otelgrpc/stats_handler.go
// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go#L52

// https://opentelemetry.io/docs/specs/otel/metrics/semantic_conventions/rpc-metrics/

// ServerHandler implements https://pkg.go.dev/google.golang.org/grpc/stats#ServerHandler
// It records OpenTelemetry metrics and traces.
type ServerHandler struct {
	handler
}

// ClientHandler implements https://pkg.go.dev/google.golang.org/grpc/stats#ServerHandler
// It records OpenTelemetry metrics and traces.
type ClientHandler struct {
	handler
}

type gRPCContextKey struct{}

type gRPCContext struct {
	messagesReceived int64
	messagesSent     int64
	attributes       []attribute.KeyValue
	startTime        time.Time
}

type handler struct {
	tracer     trace.Tracer
	meter      metric.Meter
	propagator propagation.TextMapPropagator
	// rpcDuration        metric.Int64Histogram
	// rpcRequestSize     metric.Int64Histogram
	// rpcResponseSize    metric.Int64Histogram
	// rpcRequestsPerRPC  metric.Int64Histogram
	// rpcResponsesPerRPC metric.Int64Histogram
	// rpcTotalFailed     metric.Int64Counter
	// rpcTotalSuccess    metric.Int64Counter
	requestsTotal      metric.Int64Counter
	requestsDurationMs metric.Float64Histogram
	spanKind           trace.SpanKind
	config             config
}

func newHandler(spanKind trace.SpanKind, options []Option) (handler, error) {
	c := defualtConfig

	for _, o := range options {
		o.apply(&c)
	}

	meter := c.metricsProvider.Meter(c.instrumentationName)

	requestsTotal, err := meter.Int64Counter(
		"requests_total",
	)
	if err != nil {
		return handler{}, err
	}

	requestsDurationMs, err := meter.Float64Histogram(
		"request_duration_ms",
		metric.WithUnit("ms"),
	)

	h := handler{
		tracer:             c.tracerProvider.Tracer(c.instrumentationName),
		meter:              meter,
		spanKind:           spanKind,
		config:             c,
		requestsTotal:      requestsTotal,
		requestsDurationMs: requestsDurationMs,
		// rpcDuration:        rpcServerDuration,
		// rpcRequestSize:     rpcRequestSize,
		// rpcResponseSize:    rpcResponseSize,
		// rpcRequestsPerRPC:  rpcRequestsPerRPC,
		// rpcResponsesPerRPC: rpcResponsesPerRPC,
		// rpcTotalFailed:     rpcTotalFailed,
		// rpcTotalSuccess:    rpcTotalSuccess,
	}

	return h, nil
}

func (h *handler) tagRPC(
	ctx context.Context,
	info *stats.RPCTagInfo,
) context.Context {
	ctx = extract(ctx, h.config.propagator)

	var attributes []attribute.KeyValue
	attributes = append(attributes, semconv.RPCSystemGRPC)

	parts := strings.Split(info.FullMethodName, "/")
	if len(parts) == 3 {
		attributes = append(attributes, semconv.RPCServiceKey.String(parts[1]))
		attributes = append(attributes, semconv.RPCMethodKey.String(parts[2]))
	}

	gctx := gRPCContext{attributes: attributes, startTime: time.Now()}

	return inject(
		context.WithValue(ctx, gRPCContextKey{}, &gctx),
		h.config.propagator,
	)
}

func (h *handler) handleRPC(ctx context.Context, rs stats.RPCStats) {
	_ = trace.SpanFromContext(ctx)
	gctx, _ := ctx.Value(gRPCContextKey{}).(*gRPCContext)

	switch rs := rs.(type) {
	case *stats.Begin:
	// case *stats.InPayload:
	// 	if gctx != nil {
	// 		// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go#L52
	// 		opt := metric.WithAttributes(gctx.attributes...)
	// 		h.rpcRequestSize.Record(ctx, int64(rs.Length), opt)
	// 	}

	// case *stats.OutPayload:
	// 	if gctx != nil {
	// 		// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go#L52
	// 		opt := metric.WithAttributes(gctx.attributes...)
	// 		h.rpcResponseSize.Record(ctx, int64(rs.Length), opt)
	// 	}
	case *stats.End:
		if rs.Error != nil {
			s, _ := status.FromError(rs.Error)
			gctx.attributes = append(gctx.attributes, statusCodeAttr(s.Code()))
		} else {
			gctx.attributes = append(gctx.attributes, statusCodeAttr(codes.OK))

		}
		opt := metric.WithAttributes(gctx.attributes...)
		h.requestsTotal.Add(ctx, 1, opt)

		if gctx != nil {
			duration := time.Since(gctx.startTime).Microseconds()
			opt := metric.WithAttributes(gctx.attributes...)

			h.requestsDurationMs.Record(
				ctx,
				float64(duration)/1000.0, // convert to ms
				opt,
			)
		}

	default:
		return
	}
}

func statusCodeAttr(c codes.Code) attribute.KeyValue {
	return attribute.Key("status_class").String(strconv.Itoa(int(math.Floor(float64(mapGrpcCodeToHttpStatusCode(c))/100))) + "xx")
}

func NewServerHandler(options ...Option) stats.Handler {
	h, err := newHandler(trace.SpanKindServer, options)
	if err != nil {
		otel.Handle(err)
	}

	s := &ServerHandler{
		handler: h,
	}

	return s
}

func NewClientHandler(options ...Option) stats.Handler {
	h, err := newHandler(trace.SpanKindClient, options)
	if err != nil {
		otel.Handle(err)
	}

	c := &ClientHandler{
		handler: h,
	}

	return c
}

func (s *ServerHandler) TagRPC(
	ctx context.Context,
	info *stats.RPCTagInfo,
) context.Context {
	return s.handler.tagRPC(ctx, info)
}

// HandleRPC processes the RPC stats.
func (s *ServerHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
	s.handler.handleRPC(ctx, rs)
}

// TagConn can attach some information to the given context.
func (s *ServerHandler) TagConn(
	ctx context.Context,
	_ *stats.ConnTagInfo,
) context.Context {
	// no-op
	return ctx
}

// HandleConn processes the Conn stats.
func (s *ServerHandler) HandleConn(_ context.Context, _ stats.ConnStats) {
	// no-op
}

func (c *ClientHandler) TagRPC(
	ctx context.Context,
	info *stats.RPCTagInfo,
) context.Context {
	return c.handler.tagRPC(ctx, info)
}

func (c *ClientHandler) HandleRPC(
	ctx context.Context,
	rpcStats stats.RPCStats,
) {
	c.handler.handleRPC(ctx, rpcStats)
}

func (c *ClientHandler) TagConn(
	ctx context.Context,
	_ *stats.ConnTagInfo,
) context.Context {
	// no-op
	return ctx
}

func (c *ClientHandler) HandleConn(
	_ context.Context,
	_ stats.ConnStats,
) {
	// no-op
}
