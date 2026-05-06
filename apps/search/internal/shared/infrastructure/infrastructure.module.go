package infrastructure

import (
	"github.com/go-playground/validator"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/elasticsearch"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/otel/metrics"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/otel/tracing"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/
var Module = fx.Module(
	"infrastructureModule",

	// Modules
	// provide core dependencies, e.g., logging, tracing, etc.
	core.Module,
	grpc.Module,
	rabbitmq.Module,
	redis.Module,
	metrics.Module,
	tracing.Module,
	elasticsearch.Module,

	// Other provides
	fx.Provide(validator.New),
)
