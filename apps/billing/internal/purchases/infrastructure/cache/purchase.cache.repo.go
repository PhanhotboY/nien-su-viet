package crepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
)

const (
	PURCHASE_CACHE_PREFIX = "billing_service.Purchase"
)

type purchaseCacheRepo struct {
	logger logger.Logger
	redis  redis.RedisClientWithExpire
}

func NewPurchaseCacheRepo(
	logger logger.Logger,
	redis redis.RedisClientWithExpire,
) drepo.PurchaseCacheRepo {
	return purchaseCacheRepo{
		logger: logger,
		redis:  redis,
	}
}

func (r purchaseCacheRepo) PutPurchase(ctx context.Context, key string, Purchase *entity.Purchase) error {
	r.logger.Debug("Putting purchase in cache", "key", key)
	return r.redis.HSet(ctx, PURCHASE_CACHE_PREFIX, key, Purchase) // Set with default expiration
}

func (r purchaseCacheRepo) GetPurchase(ctx context.Context, key string) (*entity.Purchase, error) {
	r.logger.Debug("Getting purchase from cache", "key", key)
	var res *entity.Purchase
	if err := r.redis.HGet(ctx, PURCHASE_CACHE_PREFIX, key, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r purchaseCacheRepo) DeletePurchase(ctx context.Context, key string) error {
	r.logger.Debug("Deleting purchase from cache", "key", key)
	if _, err := r.redis.HDel(ctx, PURCHASE_CACHE_PREFIX, key); err != nil {
		return err
	}
	return nil
}

func (r purchaseCacheRepo) DeleteAllPurchases(ctx context.Context) error {
	r.logger.Debug("Deleting all purchases from cache")
	if _, err := r.redis.MDel(ctx, PURCHASE_CACHE_PREFIX); err != nil {
		return err
	}
	return nil
}
