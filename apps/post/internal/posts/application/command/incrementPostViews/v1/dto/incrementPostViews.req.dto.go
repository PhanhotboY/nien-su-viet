package dto

// IncrementPostViewsRequest is the request DTO for incrementing post views
type IncrementPostViewsRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}
