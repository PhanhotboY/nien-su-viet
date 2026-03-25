package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// PublishPostResponse is the response DTO for publishing a post
type PublishPostResponse struct {
	Data sharedDto.OperationResponse `json:"data"`
}

func NewPublishPostResponse(id string, success bool, message string) *PublishPostResponse {
	return &PublishPostResponse{
		Data: *sharedDto.NewOperationResponse(id, success, message),
	}
}
