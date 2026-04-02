package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
)

// GetPostRes represents the response DTO for getting a single post
type GetPostRes struct {
	Data dto.PostDetailDto `json:"data"`
}

func (p *GetPostRes) FromEntity(post *entity.Post) {
	p.Data.FromEntity(post)
}
