package adto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

type UpdatePurchaseResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUpdatePurchaseResDto(id string, success bool, message string) *UpdatePurchaseResDto {
	return &UpdatePurchaseResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
