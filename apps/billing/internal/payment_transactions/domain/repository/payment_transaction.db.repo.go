package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
)

type PaymentTransactionDBRepo interface {
	CreatePaymentTransaction(ctx context.Context, transaction *entity.PaymentTransaction) (string, error)
	UpdatePaymentTransaction(ctx context.Context, transactionId string, updates map[string]any) (string, error)
	DeletePaymentTransaction(ctx context.Context, transactionId string) (string, error)
	GetPaymentTransactionById(ctx context.Context, transactionId string) (*entity.PaymentTransaction, error)
	GetPaymentTransactionsByPaymentAttemptId(ctx context.Context, paymentAttemptId string) ([]*entity.PaymentTransaction, error)
}
