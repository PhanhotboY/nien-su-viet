package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

// DeletePostResponse is the response DTO for deleting a post (single or multiple)
type DeletePostResponse struct {
	Data utils.OperationResponse `json:"data"`
}

func NewDeletePostResponse(id string, success bool, message string) *DeletePostResponse {
	return &DeletePostResponse{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
