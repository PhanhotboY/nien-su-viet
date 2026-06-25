package adto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	planhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
)

type CreatePlanReqDto struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`

	Price sdto.MoneyDto `json:"price" validate:"required"`

	BillingInterval int32 `json:"billing_interval"`

	IsActive *bool `json:"is_active" validate:""`
}

func (dto *CreatePlanReqDto) MapToEntity() *entity.Plan {
	isActive := true
	if dto.IsActive != nil {
		isActive = *dto.IsActive
	}
	return &entity.Plan{
		Code:            dto.Code,
		Name:            dto.Name,
		Price:           dto.Price.Amount,
		BillingInterval: planhelper.ToEntityInterval(&dto.BillingInterval, entity.BILLING_INTERVAL_MONTH),
		IsActive:        isActive,
		Currency:        "VND",
		CreatedAt:       time.Now(),
	}
}
