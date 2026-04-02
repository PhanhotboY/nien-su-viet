package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

// CreatePostResponse is the response DTO for creating a post
type CreatePostResponse struct {
	Data utils.OperationResponse `json:"data"`
}

func NewCreatePostResponse(id string, success bool, message string) *CreatePostResponse {
	return &CreatePostResponse{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
