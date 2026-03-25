package dto

// IncrementPostLikesRequest is the request DTO for incrementing post likes
type IncrementPostLikesRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}
