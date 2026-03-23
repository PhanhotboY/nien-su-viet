package rabbitmq

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	bus2 "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/bus"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/producer"

	// "github.com/phanhotboy/nien-su-viet/libs/pkg/health/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	// "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"
	rabbitmqconsumer "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/consumer"
	rabbitmqproducer "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/producer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"

	// "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"

	"go.uber.org/fx"
)

var (
	Module = fx.Module("rabbitmqModule",
		rabbitmqProviders,
		rabbitmqInvokes,
	)

	// - order is not important in provide
	// - provide can have parameter and will resolve if registered
	// - execute its func only if it requested
	rabbitmqProviders = fx.Options(
		fx.Provide(types.NewRabbitMQConnection),
		fx.Provide(fx.Annotate(
			bus.NewRabbitmqBus,
			fx.ParamTags(``, ``, ``, `optional:"true"`),
			fx.As(new(producer.Producer)),
			fx.As(new(bus2.Bus)),
			fx.As(new(bus.RabbitmqBus)),
		)),
		fx.Provide(rabbitmqconsumer.NewConsumerFactory),
		fx.Provide(rabbitmqproducer.NewProducerFactory),
	// fx.Provide(fx.Annotate(
	// 	NewRabbitMQHealthChecker,
	// 	fx.As(new(contracts.Health)),
	// 	fx.ResultTags(`group:"healths"`),
	// ))
	)

	rabbitmqInvokes = fx.Options(fx.Invoke(registerHooks)) //nolint:gochecknoglobals
)

func registerHooks(
	lc fx.Lifecycle,
	bus bus.RabbitmqBus,
	s settings.Config,
	logger logger.Logger,
) {
	rabbitmqOptions := s.Rmq
	if !rabbitmqOptions.AutoStart {
		return
	}

	runCtx, cancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := bus.Start(runCtx); err != nil {
					logger.Errorf(
						"(bus.Start) error in running rabbitmq server: {%v}",
						err,
					)
					return
				}
			}()
			logger.Info("rabbitmq is listening.")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			cancel()

			if err := bus.Stop(); err != nil {
				logger.Errorf("error shutting down rabbitmq server: %v", err)
			} else {
				logger.Info("rabbitmq server shutdown gracefully")
			}

			return nil
		},
	})
}
