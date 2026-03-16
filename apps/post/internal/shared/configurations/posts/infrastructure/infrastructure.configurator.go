package infrastructure

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp/contracts"
)

type InfrastructureConfigurator struct {
	contracts.Application
}

func NewInfrastructureConfigurator(
	app contracts.Application,
) *InfrastructureConfigurator {
	return &InfrastructureConfigurator{
		Application: app,
	}
}

func (ic *InfrastructureConfigurator) ConfigInfrastructures() {
	// ic.ResolveFunc(
	// 	func(l logger.Logger, tracer tracing.AppTracer, metrics metrics.AppMetrics) error {
	// 		err := mediatr.RegisterRequestPipelineBehaviors(
	// 			loggingpipelines.NewMediatorLoggingPipeline(l),
	// 			tracingpipelines.NewMediatorTracingPipeline(
	// 				tracer,
	// 				tracingpipelines.WithLogger(l),
	// 			),
	// 			metricspipelines.NewMediatorMetricsPipeline(
	// 				metrics,
	// 				metricspipelines.WithLogger(l),
	// 			),
	// 		)

	// 		return err
	// 	},
	// )
}
