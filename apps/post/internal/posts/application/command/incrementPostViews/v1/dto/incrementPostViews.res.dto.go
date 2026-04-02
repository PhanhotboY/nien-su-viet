package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

type Data struct {
	Views                   int `json:"views"`
	utils.OperationResponse `json:"operation_response"`
}

// IncrementPostViewsResponse is the response DTO for incrementing post views
type IncrementPostViewsResponse struct {
	Data `json:"data"`
}

func NewIncrementPostViewsResponse(id string, success bool, message string, views int) *IncrementPostViewsResponse {
	return &IncrementPostViewsResponse{
		Data: Data{
			Views:             views,
			OperationResponse: *utils.NewOperationResponse(id, success, message),
		},
	}
}
