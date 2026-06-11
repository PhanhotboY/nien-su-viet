package infrastructure

import (
	"github.com/go-playground/validator"
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/otel/metrics"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/otel/tracing"
	postgres "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"

	billingConfig "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/config"
)

// https://pmihaylov.com/shared-components-go-microservices/
var Module = fx.Module(
	"infrastructureModule",

	fx.Provide(
		fx.Annotate(
			func(cfg config.Config) billingConfig.BillingConfig {
				return cfg.(billingConfig.BillingConfig)
			},
			fx.As(new(billingConfig.BillingConfig)),
		)),

	// Modules
	// provide core dependencies, e.g., metrics, tracing, etc.
	core.Module,
	grpc.Module,
	rabbitmq.Module,
	postgres.Module,
	redis.Module,
	metrics.Module,
	tracing.Module,

	// Other provides
	fx.Provide(validator.New),
)
