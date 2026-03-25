package dto

// PublishPostRequest is the request DTO for publishing a post
type PublishPostRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}
