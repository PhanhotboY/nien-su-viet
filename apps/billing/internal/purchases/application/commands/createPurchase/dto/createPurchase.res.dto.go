package adto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

type CreatePurchaseResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreatePurchaseResDto(id string, success bool, message string) *CreatePurchaseResDto {
	return &CreatePurchaseResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
