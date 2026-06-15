package drepo

import (
	"context"

	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
)

type SubscriptionDbRepo interface {
	CreateSubscription(ctx context.Context, subscription *entity.Subscription) (string, error)
	UpdateSubscription(ctx context.Context, subscription *entity.Subscription) (string, error)
	GetSubscriptionByID(ctx context.Context, id string) (*entity.Subscription, error)
	GetSubscriptionsByUserID(ctx context.Context, userID string) ([]*entity.Subscription, error)
	ListSubscriptions(ctx context.Context, filter *sdto.ListQueryRequest) ([]*entity.Subscription, error)
}
