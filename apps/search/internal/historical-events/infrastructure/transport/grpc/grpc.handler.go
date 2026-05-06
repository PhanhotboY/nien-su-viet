package grpc

import (
	"context"

	"github.com/go-playground/validator"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	getAllEventsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getAllEvents/v1/queries"
	getEventQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getEvent/v1/queries"
	pb "github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/grpc/genproto"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type HistoricalEventsGrpcServerHandler struct {
	logger    logger.Logger
	validator *validator.Validate

	getEventHandler     getEventQuery.IGetEventHandler
	getAllEventsHandler getAllEventsQuery.IGetAllEventsHandler
}

func NewHistoricalEventsGrpcServerHandler(
	logger logger.Logger,
	validator *validator.Validate,

	getEventHandler getEventQuery.IGetEventHandler,
	getAllEventsHandler getAllEventsQuery.IGetAllEventsHandler,
) pb.HistoricalEventServiceServer {
	return &HistoricalEventsGrpcServerHandler{
		logger:    logger,
		validator: validator,

		getEventHandler:     getEventHandler,
		getAllEventsHandler: getAllEventsHandler,
	}
}

// ============================================================
// QUERY HANDLERS
// ============================================================

func (p *HistoricalEventsGrpcServerHandler) GetEvent(
	ctx context.Context,
	req *pb.GetHistoricalEventRequest,
) (*pb.GetHistoricalEventResponse, error) {
	p.logger.Infof("[HistoricalEventservice] Handle get historical event by id: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetEvent"))

	query, err := getEventQuery.NewGetEventQuery(req)
	if err != nil {
		p.logger.Error("[HistoricalEventservice] Invalid get event query", "error", err)
		return nil, err
	}

	data, err := p.getEventHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Errorf("[HistoricalEventservice] Failed to handle get event query: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(data, &pb.GetHistoricalEventResponse{}, p.logger)
}

func (p *HistoricalEventsGrpcServerHandler) GetEventPreview(
	ctx context.Context,
	req *pb.GetHistoricalEventPreviewRequest,
) (*pb.GetHistoricalEventPreviewResponse, error) {
	p.logger.Infof("[HistoricalEventservice] Handle get historical event preview: %+v", req)
	// span := trace.SpanFromContext(ctx)
	// span.SetAttributes(attribute.String("rpc.method", "GetEvent"))

	// query, err := .NewGetEventQuery(req)
	// if err != nil {
	// 	p.logger.Error("[HistoricalEventservice] Invalid get Event query", "error", err)
	// 	return nil, err
	// }

	// data, err := p.getEventHandler.Handle(ctx, query)
	// if err != nil {
	// 	p.logger.Errorf("[HistoricalEventservice] Failed to handle get Event query: %s", err.Error())
	// 	return nil, err
	// }

	// return grpcUtils.UnmarshalProtoMessage(data, &pb.GetEventResponse{}, p.logger)
	return nil, nil
}

func (p *HistoricalEventsGrpcServerHandler) GetAllEvents(
	ctx context.Context,
	req *pb.GetAllHistoricalEventsRequest,
) (*pb.GetAllHistoricalEventsResponse, error) {
	p.logger.Infof("[HistoricalEventService] Handle get published HistoricalEvents query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPublishedHistoricalEvents"))

	query, err := getAllEventsQuery.NewGetAllEventsQuery(req)
	if err != nil {
		p.logger.Errorf("[HistoricalEventService] Invalid get published HistoricalEvents query: %s", err.Error())
		return nil, err
	}
	data, err := p.getAllEventsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Errorf("[HistoricalEventservice] Failed to handle get published HistoricalEvents query: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(data, &pb.GetAllHistoricalEventsResponse{}, p.logger)
}
