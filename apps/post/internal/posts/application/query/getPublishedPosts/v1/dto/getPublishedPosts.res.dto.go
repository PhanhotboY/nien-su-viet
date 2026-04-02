package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type GetPublishedPostsRes struct {
	Data       []dto.PostBriefDto `json:"data"`
	Pagination utils.Pagination   `json:"pagination"`
}

func NewGetPublishedPostsRes(posts []dto.PostBriefDto, pagination utils.Pagination) *GetPublishedPostsRes {
	return &GetPublishedPostsRes{
		Data:       posts,
		Pagination: pagination,
	}
}
