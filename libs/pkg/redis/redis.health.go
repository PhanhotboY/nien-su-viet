package redis

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/health/contracts"
	"github.com/redis/go-redis/v9"
)

type RedisHealthChecker struct {
	client *redis.Client
}

func NewRedisHealthChecker(client RedisClientWithExpire) contracts.Health {
	return &RedisHealthChecker{client: client.Client}
}

func (healthChecker *RedisHealthChecker) CheckHealth(ctx context.Context) error {
	return healthChecker.client.Ping(ctx).Err()
}

func (healthChecker *RedisHealthChecker) GetHealthName() string {
	return "redis"
}
