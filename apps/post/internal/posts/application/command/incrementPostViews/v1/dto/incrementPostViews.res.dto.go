package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

type Data struct {
	Views                       int `json:"views"`
	sharedDto.OperationResponse `json:"operation_response"`
}

// IncrementPostViewsResponse is the response DTO for incrementing post views
type IncrementPostViewsResponse struct {
	Data `json:"data"`
}

func NewIncrementPostViewsResponse(id string, success bool, message string, views int) *IncrementPostViewsResponse {
	return &IncrementPostViewsResponse{
		Data: Data{
			Views:             views,
			OperationResponse: *sharedDto.NewOperationResponse(id, success, message),
		},
	}
}
