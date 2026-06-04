package adto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

type UpdatePurchaseStatusResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUpdatePurchaseStatusResDto(id string, success bool, message string) *UpdatePurchaseStatusResDto {
	return &UpdatePurchaseStatusResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
