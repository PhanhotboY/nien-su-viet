package adto

import (
	pahelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/helper"
)

type UpdatePaymentAttemptStatusReqDto struct {
	ID         string `json:"id" validate:"required,uuid"`
	PurchaseID string `json:"purchase_id" validate:"required,uuid"`
	Status     int32  `json:"status"`
}

func (dto *UpdatePaymentAttemptStatusReqDto) MapToEntity() map[string]any {
	return map[string]any{
		"status": pahelper.ToEntityStatus(&dto.Status),
	}
}
