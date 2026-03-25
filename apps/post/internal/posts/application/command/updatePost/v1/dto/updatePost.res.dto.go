package dto

import (
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// UpdatePostResponse is the response DTO for updating a post
type UpdatePostResponse struct {
	Data sharedDto.OperationResponse `json:"data"`
}

func NewUpdatePostResponse(id string, success bool, message string) *UpdatePostResponse {
	return &UpdatePostResponse{
		Data: *sharedDto.NewOperationResponse(id, success, message),
	}
}
