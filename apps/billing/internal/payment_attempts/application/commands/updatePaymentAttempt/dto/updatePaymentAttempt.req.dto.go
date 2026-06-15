package adto

import (
	"time"

	pahelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
)

type UpdatePaymentAttemptReqDto struct {
	PurchaseID string         `json:"purchase_id" validate:"required,uuid4"`
	Provider   *string        `json:"provider,omitempty"`
	Amount     *sdto.MoneyDto `json:"amount,omitempty"`
	ExpiresAt  *time.Time     `json:"expires_at,omitempty"`
	Status     *int32         `json:"status,omitempty"`
}

func (dto *UpdatePaymentAttemptReqDto) MapToEntity() map[string]any {
	updates := map[string]any{}

	if dto.Provider != nil {
		updates["provider"] = *dto.Provider
	}
	if dto.Amount != nil && dto.Amount.Amount > 0 {
		updates["amount"] = (*dto.Amount).Amount
	}
	if dto.Status != nil {
		updates["status"] = pahelper.ToEntityStatus(dto.Status)
	}
	if dto.ExpiresAt != nil {
		updates["expiresAt"] = *dto.ExpiresAt
	}
	return updates
}
