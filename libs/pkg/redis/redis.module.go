package redis

import (
	"context"
	"fmt"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/health/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"

	"go.uber.org/fx"
)

var (
	// Module provided to fxlog
	// https://uber-go.github.io/fx/modules.html
	Module = fx.Module(
		"redisModule",
		redisProviders,
		redisInvokes,
	) //nolint:gochecknoglobals

	redisProviders = fx.Options(fx.Provide( //nolint:gochecknoglobals
		NewRedisClient,
		fx.Annotate(
			NewRedisHealthChecker,
			fx.As(new(contracts.Health)),
			fx.ResultTags(fmt.Sprintf(`group:"%s"`, "healths")),
		),
	))

	redisInvokes = fx.Options(
		fx.Invoke(registerHooks),
	) //nolint:gochecknoglobals
)

func registerHooks(
	lc fx.Lifecycle,
	client RedisClientWithExpire,
	logger logger.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := client.Client.Ping(ctx).Err(); err != nil {
				logger.Errorf("error in pinging redis: %v", err)
				return err
			}

			logger.Info("redis connected successfully")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := client.Close(); err != nil {
				logger.Errorf("error in closing redis: %v", err)
			} else {
				logger.Info("redis closed gracefully")
			}

			return nil
		},
	})
}
