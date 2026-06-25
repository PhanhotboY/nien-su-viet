package adto

import (
	"github.com/google/uuid"
	subhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/helper"
)

type UpdateSubscriptionStatusReqDto struct {
	ID     string `json:"id" validate:"required,uuid4"`
	Status int32  `json:"status" validate:"required"`
}

func (dto *UpdateSubscriptionStatusReqDto) MapToEntity() map[string]interface{} {
	return map[string]interface{}{
		"id":     uuid.MustParse(dto.ID),
		"status": subhelper.ToEntityStatus(&dto.Status),
	}
}
