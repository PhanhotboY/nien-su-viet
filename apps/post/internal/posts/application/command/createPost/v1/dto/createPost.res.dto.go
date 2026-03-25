package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// CreatePostResponse is the response DTO for creating a post
type CreatePostResponse struct {
	Data sharedDto.OperationResponse `json:"data"`
}

func NewCreatePostResponse(id string, success bool, message string) *CreatePostResponse {
	return &CreatePostResponse{
		Data: *sharedDto.NewOperationResponse(id, success, message),
	}
}
