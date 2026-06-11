package adto

import (
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	pb "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/common"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type CreatePlanResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, *billing_service.CreatePlanResponse]
}

type createPlanResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreatePlanResDto(id string, success bool, message string) CreatePlanResDto {
	return createPlanResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}

func (dto createPlanResDto) ToGrpcResponse() *billing_service.CreatePlanResponse {
	return &billing_service.CreatePlanResponse{
		Data: &pb.OperationMetadata{
			Id:      dto.Data.ID,
			Success: dto.Data.Success,
		},
	}
}

func (dto createPlanResDto) GetData() utils.OperationResponse {
	return dto.Data
}
