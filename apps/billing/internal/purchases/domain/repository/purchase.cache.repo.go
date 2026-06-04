package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
)

type PurchaseCacheRepo interface {
	PutPurchase(ctx context.Context, key string, Purchase *entity.Purchase) error
	GetPurchase(ctx context.Context, key string) (*entity.Purchase, error)
	DeletePurchase(ctx context.Context, key string) error
	DeleteAllPurchases(ctx context.Context) error
}
