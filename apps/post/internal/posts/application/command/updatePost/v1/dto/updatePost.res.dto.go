package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

// UpdatePostResponse is the response DTO for updating a post
type UpdatePostResponse struct {
	Data utils.OperationResponse `json:"data"`
}

func NewUpdatePostResponse(id string, success bool, message string) *UpdatePostResponse {
	return &UpdatePostResponse{
		Data: *utils.NewOperationResponse(id, success, message),
	}
}
