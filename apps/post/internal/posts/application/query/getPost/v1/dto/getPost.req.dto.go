package dto

// GetPostQueryReq represents the request DTO for getting a single post by ID or Slug
type GetPostQueryReq struct {
	IDOrSlug string `json:"id" validate:"required"`
}
