package dto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/utils"

type Data struct {
	Likes                   int `json:"likes"`
	utils.OperationResponse `json:"operation_response"`
}

// IncrementPostLikesResponse is the response DTO for incrementing post likes
type IncrementPostLikesResponse struct {
	Data `json:"data"`
}

func NewIncrementPostLikesResponse(id string, success bool, message string, likes int) *IncrementPostLikesResponse {
	return &IncrementPostLikesResponse{
		Data: Data{
			Likes:             likes,
			OperationResponse: *utils.NewOperationResponse(id, success, message),
		},
	}
}
