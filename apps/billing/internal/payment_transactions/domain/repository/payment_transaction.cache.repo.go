package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
)

type PaymentTransactionCacheRepo interface {
	PutPaymentTransaction(ctx context.Context, key string, transaction *entity.PaymentTransaction) error
	GetPaymentTransaction(ctx context.Context, key string) (*entity.PaymentTransaction, error)
	DeletePaymentTransaction(ctx context.Context, key string) error
	DeleteAllPaymentTransactions(ctx context.Context) error
}
