package infrastructure

import (
	"github.com/go-playground/validator"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/

var Module = fx.Module(
	"infrastructureModule",
	// Modules
	// core.Module,
	// customEcho.Module,
	grpc.Module,
	// mongodb.Module,
	// elasticsearch.Module,
	// eventstroredb.ModuleFunc(
	// 	func(params params.OrderProjectionParams) eventstroredb.ProjectionBuilderFuc {
	// 		return func(builder eventstroredb.ProjectionsBuilder) {
	// 			builder.AddProjections(params.Projections)
	// 		}
	// 	},
	// ),
	// rabbitmq.ModuleFunc(
	// 	func() configurations.RabbitMQConfigurationBuilderFuc {
	// 		return func(builder configurations.RabbitMQConfigurationBuilder) {
	// 			rabbitmq2.ConfigOrdersRabbitMQ(builder)
	// 		}
	// 	},
	// ),
	// health.Module,
	// tracing.Module,
	// metrics.Module,

	// Other provides
	fx.Provide(validator.New),
)
