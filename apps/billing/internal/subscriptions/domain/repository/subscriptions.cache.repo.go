package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
)

type SubscriptionCacheRepo interface {
	PutSubscription(ctx context.Context, key string, Subscription *entity.Subscription) error
	GetSubscription(ctx context.Context, key string) (*entity.Subscription, error)
	DeleteSubscription(ctx context.Context, key string) error
	DeleteAllSubscriptions(ctx context.Context) error
}
