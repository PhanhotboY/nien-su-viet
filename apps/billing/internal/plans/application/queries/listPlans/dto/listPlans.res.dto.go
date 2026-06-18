package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	planhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

type Plan struct {
	ID uuid.UUID `json:"id"`

	Code string `json:"code"`
	Name string `json:"name"`

	Price sdto.MoneyDto `json:"price"`

	BillingInterval entity.BillingInterval `json:"billing_interval"`

	IsActive bool `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
}

type ListPlansResDto interface {
	cdto.ApplicationResponse[[]Plan, *billing_service.ListPlansResponse]
}

type listPlansResDto struct {
	Data []Plan `json:"data"`
}

func NewListPlansResDto(plans []*entity.Plan) ListPlansResDto {
	res := make([]Plan, len(plans))
	for i, plan := range plans {
		res[i] = Plan{
			ID:   plan.ID,
			Code: plan.Code,
			Name: plan.Name,
			Price: sdto.MoneyDto{
				Amount:   plan.Price,
				Currency: plan.Currency},
			BillingInterval: plan.BillingInterval,
			IsActive:        plan.IsActive,
			CreatedAt:       plan.CreatedAt,
		}
	}
	return listPlansResDto{Data: res}
}

func (dto listPlansResDto) ToGrpcResponse() *billing_service.ListPlansResponse {
	res := make([]*billing_service.Plan, len(dto.Data))
	for i, plan := range dto.Data {
		res[i] = &billing_service.Plan{
			Id:              plan.ID.String(),
			Code:            plan.Code,
			Name:            plan.Name,
			Price:           plan.Price.ToGrpcMoney(),
			BillingInterval: planhelper.ToGrpcInterval(plan.BillingInterval),
			IsActive:        plan.IsActive,
			CreatedAt:       grpcUtils.TimeToTimestamp(&plan.CreatedAt),
		}
	}
	return &billing_service.ListPlansResponse{Data: res}
}

func (dto listPlansResDto) GetData() []Plan {
	return dto.Data
}
