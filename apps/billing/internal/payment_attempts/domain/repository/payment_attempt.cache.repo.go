package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
)

type PaymentAttemptCacheRepo interface {
	PutPaymentAttempt(ctx context.Context, key string, paymentAttempt *entity.PaymentAttempt) error
	GetPaymentAttempt(ctx context.Context, key string) (*entity.PaymentAttempt, error)
	DeletePaymentAttempt(ctx context.Context, key string) error
	DeleteAllPaymentAttempts(ctx context.Context) error
}
