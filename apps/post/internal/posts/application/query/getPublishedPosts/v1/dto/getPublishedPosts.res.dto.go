package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
)

type GetPublishedPostsRes struct {
	Data       []dto.PostBriefDto      `json:"data"`
	Pagination sharedDto.PaginationDto `json:"pagination"`
}

func NewGetPublishedPostsRes(posts []dto.PostBriefDto, pagination sharedDto.PaginationDto) *GetPublishedPostsRes {
	return &GetPublishedPostsRes{
		Data:       posts,
		Pagination: pagination,
	}
}
