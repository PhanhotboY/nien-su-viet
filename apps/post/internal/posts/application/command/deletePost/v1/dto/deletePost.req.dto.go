package dto

// DeletePostRequest is the request DTO for deleting a single post
type DeletePostRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}
