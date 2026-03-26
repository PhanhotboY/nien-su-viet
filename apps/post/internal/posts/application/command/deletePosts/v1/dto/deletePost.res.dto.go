package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// DeletePostResponse is the response DTO for deleting a post (single or multiple)
type DeletePostsResponse struct {
	Data sharedDto.OperationResponse `json:"data"`
}

func NewDeletePostsResponse(id string, success bool, message string) *DeletePostsResponse {
	return &DeletePostsResponse{
		Data: *sharedDto.NewOperationResponse(id, success, message),
	}
}
