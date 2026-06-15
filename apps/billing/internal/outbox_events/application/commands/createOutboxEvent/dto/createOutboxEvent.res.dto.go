package adto

import (
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type CreateOutboxEventResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, any]
}

type createOutboxEventResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreateOutboxEventResDto(id string, success bool, message string) CreateOutboxEventResDto {
	return &createOutboxEventResDto{
		Data: utils.OperationResponse{
			ID:      id,
			Success: success,
			Message: message,
		},
	}
}

func (d *createOutboxEventResDto) GetData() utils.OperationResponse {
	return d.Data
}

func (d *createOutboxEventResDto) ToGrpcResponse() any {
	return nil
}
