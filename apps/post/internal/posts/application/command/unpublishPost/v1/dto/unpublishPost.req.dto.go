package dto

// UnpublishPostRequest is the request DTO for unpublishing a post
type UnpublishPostRequest struct {
ID string `json:"id" validate:"required,uuid"`
}
