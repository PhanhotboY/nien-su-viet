package dto

import (
	postDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

// GetAllPostsResponse is the response DTO for getting all posts
type GetAllPostsRes struct {
	Data       []postDto.PostBriefDto  `json:"data"`
	Pagination sharedDto.PaginationDto `json:"pagination,omitempty"`
}

func NewGetAllPostsRes(posts []postDto.PostBriefDto, pagination sharedDto.PaginationDto) *GetAllPostsRes {
	return &GetAllPostsRes{
		Data:       posts,
		Pagination: pagination,
	}
}
