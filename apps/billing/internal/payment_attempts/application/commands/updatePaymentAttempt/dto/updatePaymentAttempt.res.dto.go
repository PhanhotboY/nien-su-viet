package adto

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
)

type UpdatePaymentAttemptResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, any]
}

type updatePaymentAttemptResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUpdatePaymentAttemptResDto(id string, success bool, message string) UpdatePaymentAttemptResDto {
	return &updatePaymentAttemptResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}

// this command isn't exposed to gRPC
func (dto *updatePaymentAttemptResDto) ToGrpcResponse() any {
	return nil
}

func (dto *updatePaymentAttemptResDto) GetData() utils.OperationResponse {
	return dto.Data
}
