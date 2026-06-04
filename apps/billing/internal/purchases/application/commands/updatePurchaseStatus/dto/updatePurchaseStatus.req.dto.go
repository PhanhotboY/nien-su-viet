package adto

import (
	"github.com/google/uuid"
	purhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/helper"
)

type UpdatePurchaseStatusReqDto struct {
	ID     string `json:"id" validate:"required,uuid4"`
	Status int32  `json:"status" validate:"required"`
}

func (dto *UpdatePurchaseStatusReqDto) MapToEntity() map[string]interface{} {
	return map[string]interface{}{
		"ID":     uuid.MustParse(dto.ID),
		"Status": purhelper.ToEntityStatus(&dto.Status),
	}
}
