package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
)

type PurchaseDBRepo interface {
	CreatePurchase(ctx context.Context, purchase *entity.Purchase) (string, error)
	UpdatePurchase(ctx context.Context, purchaseId string, updates map[string]any) (string, error)
	DeletePurchase(ctx context.Context, purchaseId string) (string, error)
	GetPurchase(ctx context.Context, purchaseId string) (*entity.Purchase, error)
}
