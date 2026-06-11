package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
)

type CreatePurchaseReqDto struct {
	UserID string `json:"user_id" validate:"required"`
	PlanID string `json:"plan_id" validate:"required,uuid4"`

	// Idempotent purchase creation (critical)
	IdempotencyKey string `json:"idempotency_key" validate:"required,uuid4"`

	CreatedAt   *time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at" validate:"omitempty"`
}

func (dto *CreatePurchaseReqDto) MapToEntity() *entity.Purchase {
	createdAt := time.Now()
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	// Subscription is is not created yet, update later on payment succeeded
	return &entity.Purchase{
		UserID: dto.UserID,
		PlanID: uuid.MustParse(dto.PlanID),
		// New Purchase must be PENDING
		Status:         entity.PURCHASE_STATUS_PENDING,
		IdempotencyKey: dto.IdempotencyKey,
		CreatedAt:      createdAt,
		CompletedAt:    dto.CompletedAt,
	}
}
