package contracts

import (
	"go.opentelemetry.io/otel/metric"
)

type PostsMetrics struct {
	RequestsTotal    metric.Int64Counter
	RequestsDuration metric.Float64Histogram
}
