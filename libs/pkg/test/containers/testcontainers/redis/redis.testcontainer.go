package redis_testcontainer

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	coptions "github.com/phanhotboy/nien-su-viet/libs/pkg/config/options"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type RedisTestContainers struct {
	container testcontainers.Container
	logger    logger.Logger
	cfg       coptions.RedisOptions
}

func NewRedisTestContainers(c config.Config, l logger.Logger) *RedisTestContainers {
	return &RedisTestContainers{
		logger: l,
		cfg:    c.GetRedisOptions(),
	}
}

func (g *RedisTestContainers) Start(ctx context.Context, t *testing.T) (coptions.RedisOptions, error) {
	g.logger.Info("Starting Redis test container...")
	containerReq := g.getRunOptions()
	dbContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return g.cfg, err
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		return g.cfg, err
	}
	port, err := dbContainer.MappedPort(ctx, nat.Port(strconv.Itoa(g.cfg.Port)))
	if err != nil {
		return g.cfg, err
	}

	cfg := coptions.RedisOptions{
		Host:          host,
		Port:          port.Int(),
		Username:      g.cfg.Username,
		Password:      g.cfg.Password,
		Database:      g.cfg.Database,
		PoolSize:      g.cfg.PoolSize,
		EnableTracing: g.cfg.EnableTracing,
	}

	isConnectable := isConnectable(ctx, g.logger, cfg)
	if !isConnectable {
		return g.cfg, fmt.Errorf("redis container is not connectable on host: %s:%d", host, cfg.Port)
	}

	g.container = dbContainer

	return cfg, nil
}

func (g *RedisTestContainers) getRunOptions() testcontainers.ContainerRequest {
	containerReq := testcontainers.ContainerRequest{
		Image:        "redis:alpine",
		ExposedPorts: []string{strconv.Itoa(g.cfg.Port)},
		WaitingFor: wait.ForListeningPort(nat.Port(strconv.Itoa(g.cfg.Port))).
			WithPollInterval(2 * time.Second),
		Hostname: g.cfg.Host,
		Env:      map[string]string{},
	}

	return containerReq
}

func isConnectable(
	ctx context.Context,
	logger logger.Logger,
	options coptions.RedisOptions,
) bool {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", options.Host, options.Port),
	})

	defer redisClient.Close()

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		// we should not use `t.Error` or `t.Errorf` for logging errors because it will `fail` our test at the end and, we just should use logs without error like log.Error (not log.Fatal)
		logger.Errorf(
			"Error in creating redis connection with %s:%d",
			options.Host,
			options.Port,
		)

		return false
	}

	logger.Infof(
		"Opened redis connection on host: %s:%d",
		options.Host,
		options.Port,
	)

	return true
}
