package adto

import (
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type CreateSubscriptionResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, any]
}

type createSubscriptionResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreateSubscriptionResDto(id string, success bool, message string) CreateSubscriptionResDto {
	return &createSubscriptionResDto{
		Data: utils.OperationResponse{
			ID:      id,
			Success: success,
			Message: message,
		},
	}
}

func (d *createSubscriptionResDto) GetData() utils.OperationResponse {
	return d.Data
}

func (d *createSubscriptionResDto) ToGrpcResponse() any {
	return nil
}
