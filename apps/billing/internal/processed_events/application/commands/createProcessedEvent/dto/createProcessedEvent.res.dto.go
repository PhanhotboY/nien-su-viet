package adto

import (
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type CreateProcessedEventResDto interface {
	cdto.ApplicationResponse[utils.OperationResponse, any]
}

type createProcessedEventResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreateProcessedEventResDto(id string, success bool, message string) CreateProcessedEventResDto {
	return &createProcessedEventResDto{
		Data: utils.OperationResponse{
			ID:      id,
			Success: success,
			Message: message,
		},
	}
}

func (d *createProcessedEventResDto) GetData() utils.OperationResponse {
	return d.Data
}

func (d *createProcessedEventResDto) ToGrpcResponse() any {
	return nil
}
