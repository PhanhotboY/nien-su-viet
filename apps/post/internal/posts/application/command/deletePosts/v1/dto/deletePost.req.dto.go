package dto

// DeletePostsRequest is the request DTO for deleting multiple posts
type DeletePostsRequest struct {
	PostIds []string `json:"post_ids" validate:"required,min=1,dive,uuid"`
}
