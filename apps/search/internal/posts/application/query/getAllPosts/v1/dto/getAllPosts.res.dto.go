package dto

import (
	postDto "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

// GetAllPostsResponse is the response DTO for getting all posts
type GetAllPostsRes struct {
	Data       []postDto.PostBriefDto `json:"data"`
	Pagination utils.Pagination       `json:"pagination"`
}

func NewGetAllPostsRes(posts []postDto.PostBriefDto, pagination utils.Pagination) *GetAllPostsRes {
	return &GetAllPostsRes{
		Data:       posts,
		Pagination: pagination,
	}
}
