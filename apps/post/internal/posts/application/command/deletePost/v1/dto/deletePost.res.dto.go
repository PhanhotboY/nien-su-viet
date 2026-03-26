package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// DeletePostResponse is the response DTO for deleting a post (single or multiple)
type DeletePostResponse struct {
	Data sharedDto.OperationResponse `json:"data"`
}

func NewDeletePostResponse(id string, success bool, message string) *DeletePostResponse {
	return &DeletePostResponse{
		Data: *sharedDto.NewOperationResponse(id, success, message),
	}
}
