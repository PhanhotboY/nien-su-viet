package adto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

type UpdateSubscriptionStatusResDto struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUpdateSubscriptionStatusResDto(id string, success bool, message string) *UpdateSubscriptionStatusResDto {
	return &UpdateSubscriptionStatusResDto{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
