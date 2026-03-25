package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// UnpublishPostResponse is the response DTO for unpublishing a post
type UnpublishPostResponse struct {
	Data sharedDto.OperationResponse `json:"data"`
}

func NewUnpublishPostResponse(id string, success bool, message string) *UnpublishPostResponse {
	return &UnpublishPostResponse{
		Data: *sharedDto.NewOperationResponse(id, success, message),
	}
}
