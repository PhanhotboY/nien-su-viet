package adto

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	planhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	pb "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

type GetPlanByIdResDto interface {
	cdto.ApplicationResponse[GetPlanByIdResData, *pb.GetPlanResponse]
}

type GetPlanByIdResData struct {
	Id              uuid.UUID              `json:"id,omitempty"`
	Code            string                 `json:"code,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Price           sdto.MoneyDto          `json:"price"`
	BillingInterval entity.BillingInterval `json:"billing_interval,omitempty"`
	IsActive        bool                   `json:"is_active,omitempty"`
	CreatedAt       time.Time              `json:"created_at"`
}

type getPlanByIdResDto struct {
	Data GetPlanByIdResData `json:"data"`
}

func NewGetPlanByIdResDto(plan *entity.Plan) GetPlanByIdResDto {
	return getPlanByIdResDto{
		Data: GetPlanByIdResData{
			Id:              plan.ID,
			Code:            plan.Code,
			Name:            plan.Name,
			Price:           *sdto.NewMoneyDto(plan.Price, plan.Currency),
			BillingInterval: plan.BillingInterval,
			IsActive:        plan.IsActive,
			CreatedAt:       plan.CreatedAt,
		},
	}
}

func (dto getPlanByIdResDto) ToGrpcResponse() *pb.GetPlanResponse {
	fmt.Printf("Converting GetPlanByIdResDto to gRPC response: %+v\n", dto)
	return &pb.GetPlanResponse{
		Data: &pb.Plan{
			Id:              dto.Data.Id.String(),
			Code:            dto.Data.Code,
			Name:            dto.Data.Name,
			Price:           dto.Data.Price.ToGrpcMoney(),
			BillingInterval: planhelper.ToGrpcInterval(dto.Data.BillingInterval),
			IsActive:        dto.Data.IsActive,
			CreatedAt:       grpcUtils.TimeToTimestamp(dto.Data.CreatedAt),
		},
	}
}

func (dto getPlanByIdResDto) GetData() GetPlanByIdResData {
	return dto.Data
}
