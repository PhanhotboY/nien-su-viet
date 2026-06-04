package adto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	purhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/helper"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

type PurchaseDto struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	SubscriptionID string `json:"subscription_id"`
	PlanID         string `json:"plan_id"`

	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`

	Status billing_service.PurchaseStatus `json:"status"`

	// Idempotent purchase creation (critical)
	IdempotencyKey string `json:"idempotency_key"`

	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type GetPurchaseResDto struct {
	Data PurchaseDto `json:"data"`
}

func NewGetPurchaseResDto(entity entity.Purchase) *GetPurchaseResDto {
	purchase := PurchaseDto{
		ID:             entity.ID.String(),
		UserID:         entity.UserID,
		SubscriptionID: entity.SubscriptionID.String(),
		PlanID:         entity.PlanID.String(),

		Amount:   entity.Amount,
		Currency: entity.Currency,

		Status: purhelper.ToGrpcStatus(entity.Status),

		IdempotencyKey: entity.IdempotencyKey,

		CreatedAt:   entity.CreatedAt,
		CompletedAt: entity.CompletedAt,
	}

	return &GetPurchaseResDto{
		Data: purchase,
	}
}
