package worker

import (
	"context"

	cworker "github.com/phanhotboy/nien-su-viet/libs/pkg/core/worker"
	"go.uber.org/fx"
)

type WorkersParams struct {
	fx.In

	Workers []cworker.Worker `group:"workers"`
}

var Module = fx.Module(
	"workerModule",

	fx.Provide(fx.Annotate(NewInboxEventsWorker,
		fx.ResultTags(`group:"workers"`))),

	fx.Invoke(func(lc fx.Lifecycle, params WorkersParams) {
		runner := cworker.NewWorkersRunner(params.Workers)

		lc.Append(
			fx.Hook{
				OnStart: func(ctx context.Context) error {
					runner.Start(context.Background())
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return runner.Stop(ctx)
				},
			},
		)

	}),
)
