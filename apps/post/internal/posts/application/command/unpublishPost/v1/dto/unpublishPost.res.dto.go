package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

// UnpublishPostResponse is the response DTO for unpublishing a post
type UnpublishPostResponse struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUnpublishPostResponse(id string, success bool, message string) *UnpublishPostResponse {
	return &UnpublishPostResponse{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
