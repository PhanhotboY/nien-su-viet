package adto

import (
	"github.com/google/uuid"
	purhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/helper"
)

type UpdatePurchaseReqDto struct {
	ID     string `json:"id" validate:"required,uuid4"`
	Status *int32 `json:"status" validate:"required"`

	SubscriptionID *string `json:"subscription_id,omitempty" validate:"omitempty,uuid4"`

	PlanID *uuid.UUID `json:"plan_id,omitempty" validate:"omitempty,uuid4"`

	Amount   *int64  `json:"amount,omitempty" validate:"omitempty,gt=0"`
	Currency *string `json:"currency,omitempty" validate:"omitempty,len=3"`
}

func (dto *UpdatePurchaseReqDto) MapToEntity() map[string]interface{} {
	updates := map[string]interface{}{
		"id": uuid.MustParse(dto.ID),
	}
	if dto.Status != nil {
		updates["status"] = purhelper.ToEntityStatus(dto.Status)
	}
	if dto.SubscriptionID != nil {
		updates["subscription_id"] = *dto.SubscriptionID
	}
	if dto.PlanID != nil {
		updates["plan_id"] = *dto.PlanID
	}
	if dto.Amount != nil {
		updates["amount"] = *dto.Amount
	}
	if dto.Currency != nil {
		updates["currency"] = *dto.Currency
	}
	return updates
}
