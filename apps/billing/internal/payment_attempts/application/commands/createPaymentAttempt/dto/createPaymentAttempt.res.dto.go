package adto

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/common"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
)

type CreatePaymentAttemptResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, *billing_service.CreatePaymentAttemptResponse]
}

type createPaymentAttemptResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreatePaymentAttemptResDto(id string, success bool, message string) CreatePaymentAttemptResDto {
	return &createPaymentAttemptResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}

func (dto *createPaymentAttemptResDto) ToGrpcResponse() *billing_service.CreatePaymentAttemptResponse {
	return &billing_service.CreatePaymentAttemptResponse{
		Data: &common.OperationMetadata{
			Id:      "",
			Success: true,
		},
	}
}

func (dto *createPaymentAttemptResDto) GetData() utils.OperationResponse {
	return dto.Data
}
