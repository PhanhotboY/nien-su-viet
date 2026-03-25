package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

type GetPopularPostsRes struct {
	Data       []dto.PostBriefDto      `json:"data"`
	Pagination sharedDto.PaginationDto `json:"pagination,omitempty"`
}

func NewGetPopularPostsRes(posts []dto.PostBriefDto, pagination sharedDto.PaginationDto) *GetPopularPostsRes {
	return &GetPopularPostsRes{
		Data:       posts,
		Pagination: pagination,
	}
}
