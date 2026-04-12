package infrastructure

import (
	"context"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/go-playground/validator"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	postgres "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
	postgres_testcontainer "github.com/phanhotboy/nien-su-viet/libs/pkg/test/containers/testcontainers/postgres"
	rabbitmq_testcontainer "github.com/phanhotboy/nien-su-viet/libs/pkg/test/containers/testcontainers/rabbitmq"
	redis_testcontainer "github.com/phanhotboy/nien-su-viet/libs/pkg/test/containers/testcontainers/redis"

	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/
var TestModuleFunc = func(t *testing.T) fx.Option {
	ctx := context.Background()

	return fx.Module(
		"infrastructureTestModule",

		fx.Provide(postgres_testcontainer.NewPostgresTestContainers),
		fx.Provide(rabbitmq_testcontainer.NewRabbitMQTestContainers),
		fx.Provide(redis_testcontainer.NewRedisTestContainers),

		fx.Decorate(func(pg *postgres_testcontainer.PostgresTestContainers,
			rmq *rabbitmq_testcontainer.RabbitMQTestContainers,
			redis *redis_testcontainer.RedisTestContainers,
			cfg settings.Config,
		) settings.Config {
			pgCfg, err := pg.Start(ctx, t)
			if err != nil {
				t.Fatalf("failed to start postgres test container: %s", err)
			}
			rmqCfg, err := rmq.Start(ctx, t)
			if err != nil {
				t.Fatalf("failed to start rabbitmq test container: %s", err)
			}
			redisCfg, err := redis.Start(ctx, t)
			if err != nil {
				t.Fatalf("failed to start redis test container: %s", err)
			}

			return settings.Config{
				Server: cfg.Server,
				Grpc: settings.GrpcConfig{
					Port: nat.Port("").Port(),
				},
				Logger:     cfg.Logger,
				Postgresql: pgCfg,
				Rmq:        rmqCfg,
				Redis:      redisCfg,
			}
		}),

		// Modules
		// provide core dependencies, e.g., logging, tracing, etc.
		core.Module,
		grpc.Module,
		rabbitmq.Module,
		postgres.Module,
		redis.Module,

		// Other provides
		fx.Provide(validator.New),
	)
}
