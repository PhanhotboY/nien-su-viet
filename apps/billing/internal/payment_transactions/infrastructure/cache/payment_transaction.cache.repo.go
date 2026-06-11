package crepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
)

const (
	PAYMENT_TRANSACTION_CACHE_PREFIX = "billing_service.PaymentTransaction"
)

type paymentTransactionCacheRepo struct {
	logger logger.Logger
	redis  redis.RedisClientWithExpire
}

func NewPaymentTransactionCacheRepo(
	logger logger.Logger,
	redis redis.RedisClientWithExpire,
) drepo.PaymentTransactionCacheRepo {
	return paymentTransactionCacheRepo{
		logger: logger,
		redis:  redis,
	}
}

func (r paymentTransactionCacheRepo) PutPaymentTransaction(ctx context.Context, key string, transaction *entity.PaymentTransaction) error {
	r.logger.Debug("Putting payment transaction in cache", "key", key)
	return r.redis.HSet(ctx, PAYMENT_TRANSACTION_CACHE_PREFIX, key, transaction)
}

func (r paymentTransactionCacheRepo) GetPaymentTransaction(ctx context.Context, key string) (*entity.PaymentTransaction, error) {
	r.logger.Debug("Getting payment transaction from cache", "key", key)
	var res = new(entity.PaymentTransaction)
	if err := r.redis.HGet(ctx, PAYMENT_TRANSACTION_CACHE_PREFIX, key, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r paymentTransactionCacheRepo) DeletePaymentTransaction(ctx context.Context, key string) error {
	r.logger.Debug("Deleting payment transaction from cache", "key", key)
	if _, err := r.redis.HDel(ctx, PAYMENT_TRANSACTION_CACHE_PREFIX, key); err != nil {
		return err
	}
	return nil
}

func (r paymentTransactionCacheRepo) DeleteAllPaymentTransactions(ctx context.Context) error {
	r.logger.Debug("Deleting all payment transactions from cache")
	if _, err := r.redis.MDel(ctx, PAYMENT_TRANSACTION_CACHE_PREFIX); err != nil {
		return err
	}
	return nil
}
