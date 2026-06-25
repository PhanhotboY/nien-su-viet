package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
)

type PaymentAttemptDBRepo interface {
	CreatePaymentAttempt(ctx context.Context, paymentAttempt *entity.PaymentAttempt) (string, error)
	UpdatePaymentAttempt(ctx context.Context, paymentAttemptId string, updates map[string]any) (string, error)
	DeletePaymentAttempt(ctx context.Context, paymentAttemptId string) (string, error)
	GetPaymentAttemptById(ctx context.Context, paymentAttemptId string) (*entity.PaymentAttempt, error)
	GetPaymentAttemptByProviderTransactionID(ctx context.Context, provider string, providerTransactionID string) (*entity.PaymentAttempt, error)
}
