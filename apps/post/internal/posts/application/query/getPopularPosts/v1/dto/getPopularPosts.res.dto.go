package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type GetPopularPostsRes struct {
	Data       []dto.PostBriefDto `json:"data"`
	Pagination utils.Pagination   `json:"pagination,omitempty"`
}

func NewGetPopularPostsRes(posts []dto.PostBriefDto, pagination utils.Pagination) *GetPopularPostsRes {
	return &GetPopularPostsRes{
		Data:       posts,
		Pagination: pagination,
	}
}
