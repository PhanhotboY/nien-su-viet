package adto

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
)

type UpdatePaymentAttemptStatusResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, any]
}

type updatePaymentAttemptStatusResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUpdatePaymentAttemptStatusResDto(id string, success bool, message string) UpdatePaymentAttemptStatusResDto {
	return &updatePaymentAttemptStatusResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}

// this command isn't exposed to gRPC
func (dto *updatePaymentAttemptStatusResDto) ToGrpcResponse() any {
	return nil
}

func (dto *updatePaymentAttemptStatusResDto) GetData() utils.OperationResponse {
	return dto.Data
}
