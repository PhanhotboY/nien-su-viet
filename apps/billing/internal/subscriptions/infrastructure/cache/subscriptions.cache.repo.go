package crepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
)

const (
	SUBSCRIPTION_CACHE_PREFIX = "billing_service.Subscription"
)

type subscriptionCacheRepo struct {
	logger logger.Logger
	redis  redis.RedisClientWithExpire
}

func NewSubscriptionCacheRepo(
	logger logger.Logger,
	redis redis.RedisClientWithExpire,
) drepo.SubscriptionCacheRepo {
	return subscriptionCacheRepo{
		logger: logger,
		redis:  redis,
	}
}

func (r subscriptionCacheRepo) PutSubscription(ctx context.Context, key string, Subscription *entity.Subscription) error {
	r.logger.Debug("Putting Subscription in cache", "key", key)
	return r.redis.HSet(ctx, SUBSCRIPTION_CACHE_PREFIX, key, Subscription) // Set with default expiration
}

func (r subscriptionCacheRepo) GetSubscription(ctx context.Context, key string) (*entity.Subscription, error) {
	r.logger.Debug("Getting Subscription from cache", "key", key)
	var res *entity.Subscription
	if err := r.redis.HGet(ctx, SUBSCRIPTION_CACHE_PREFIX, key, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r subscriptionCacheRepo) DeleteSubscription(ctx context.Context, key string) error {
	r.logger.Debug("Deleting Subscription from cache", "key", key)
	if _, err := r.redis.HDel(ctx, SUBSCRIPTION_CACHE_PREFIX, key); err != nil {
		return err
	}
	return nil
}

func (r subscriptionCacheRepo) DeleteAllSubscriptions(ctx context.Context) error {
	r.logger.Debug("Deleting all Subscriptions from cache")
	if _, err := r.redis.MDel(ctx, SUBSCRIPTION_CACHE_PREFIX); err != nil {
		return err
	}
	return nil
}
