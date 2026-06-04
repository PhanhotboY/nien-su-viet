package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	purhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/helper"
)

type CreatePurchaseReqDto struct {
	UserID         string `json:"user_id" validate:"required"`
	SubscriptionID string `json:"subscription_id" validate:"required,uuid4"`
	PlanID         string `json:"plan_id" validate:"required,uuid4"`

	Amount   int64  `json:"amount" validate:"required,gt=0"`
	Currency string `json:"currency" validate:"required"`

	Status *int32 `json:"status"`
	// Idempotent purchase creation (critical)
	IdempotencyKey string `json:"idempotency_key" validate:"required,len=128"`

	CreatedAt   time.Time  `json:"created_at" validate:"required"`
	CompletedAt *time.Time `json:"completed_at" validate:"omitempty"`
}

func (dto *CreatePurchaseReqDto) MapToEntity() *entity.Purchase {
	return &entity.Purchase{
		UserID:         dto.UserID,
		SubscriptionID: uuid.MustParse(dto.SubscriptionID),
		PlanID:         uuid.MustParse(dto.PlanID),
		Amount:         dto.Amount,
		Currency:       dto.Currency,
		Status:         purhelper.ToEntityStatus(dto.Status),
		IdempotencyKey: dto.IdempotencyKey,
		CreatedAt:      dto.CreatedAt,
		CompletedAt:    dto.CompletedAt,
	}
}
