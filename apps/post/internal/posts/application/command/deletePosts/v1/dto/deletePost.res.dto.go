package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

// DeletePostResponse is the response DTO for deleting a post (single or multiple)
type DeletePostsResponse struct {
	Data utils.OperationResponse `json:"data"`
}

func NewDeletePostsResponse(id string, success bool, message string) *DeletePostsResponse {
	return &DeletePostsResponse{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
