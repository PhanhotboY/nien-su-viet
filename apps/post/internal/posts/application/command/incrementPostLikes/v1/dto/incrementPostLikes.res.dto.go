package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

type Data struct {
	Likes                       int `json:"likes"`
	sharedDto.OperationResponse `json:"operation_response"`
}

// IncrementPostLikesResponse is the response DTO for incrementing post likes
type IncrementPostLikesResponse struct {
	Data `json:"data"`
}

func NewIncrementPostLikesResponse(id string, success bool, message string, likes int) *IncrementPostLikesResponse {
	return &IncrementPostLikesResponse{
		Data: Data{
			Likes:             likes,
			OperationResponse: *sharedDto.NewOperationResponse(id, success, message),
		},
	}
}
