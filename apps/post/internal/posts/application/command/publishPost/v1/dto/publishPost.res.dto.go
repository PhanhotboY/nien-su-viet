package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

// PublishPostResponse is the response DTO for publishing a post
type PublishPostResponse struct {
	Data utils.OperationResponse `json:"data"`
}

func NewPublishPostResponse(id string, success bool, message string) *PublishPostResponse {
	return &PublishPostResponse{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
