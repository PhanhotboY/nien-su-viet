package crepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
)

const (
	PAYMENT_ATTEMPT_CACHE_PREFIX = "billing_service.PaymentAttempt"
)

type paymentAttemptCacheRepo struct {
	logger logger.Logger
	redis  redis.RedisClientWithExpire
}

func NewPaymentAttemptCacheRepo(
	logger logger.Logger,
	redis redis.RedisClientWithExpire,
) repository.PaymentAttemptCacheRepo {
	return paymentAttemptCacheRepo{
		logger: logger,
		redis:  redis,
	}
}

func (r paymentAttemptCacheRepo) PutPaymentAttempt(ctx context.Context, key string, paymentAttempt *entity.PaymentAttempt) error {
	r.logger.Debug("Putting payment attempt in cache", "key", key)
	return r.redis.HSet(ctx, PAYMENT_ATTEMPT_CACHE_PREFIX, key, paymentAttempt)
}

func (r paymentAttemptCacheRepo) GetPaymentAttempt(ctx context.Context, key string) (*entity.PaymentAttempt, error) {
	r.logger.Debug("Getting payment attempt from cache", "key", key)
	var res *entity.PaymentAttempt
	if err := r.redis.HGet(ctx, PAYMENT_ATTEMPT_CACHE_PREFIX, key, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r paymentAttemptCacheRepo) DeletePaymentAttempt(ctx context.Context, key string) error {
	r.logger.Debug("Deleting payment attempt from cache", "key", key)
	if _, err := r.redis.HDel(ctx, PAYMENT_ATTEMPT_CACHE_PREFIX, key); err != nil {
		return err
	}
	return nil
}

func (r paymentAttemptCacheRepo) DeleteAllPaymentAttempts(ctx context.Context) error {
	r.logger.Debug("Deleting all payment attempts from cache")
	if _, err := r.redis.MDel(ctx, PAYMENT_ATTEMPT_CACHE_PREFIX); err != nil {
		return err
	}
	return nil
}
