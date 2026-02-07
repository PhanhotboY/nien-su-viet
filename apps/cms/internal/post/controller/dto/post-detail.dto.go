package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

type PostDetail struct {
	PostBrief
	Content string `json:"content" doc:"Post content"` // Post content
}

func (dto *PostDetail) FromEntity(entity *entity.Post) {
	if dto == nil || entity == nil {
		return
	}

	dto.ID = entity.Id
	dto.Title = entity.Title
	dto.Slug = entity.Slug
	dto.CategoryId = entity.CategoryId
	dto.Thumbnail = entity.Thumbnail
	dto.AuthorId = entity.AuthorId
	dto.Content = entity.Content
	dto.Summary = entity.Summary
	dto.Views = entity.Views
	dto.Likes = entity.Likes
	dto.Published = entity.Published
	dto.CreatedAt = entity.CreatedAt
	dto.UpdatedAt = entity.UpdatedAt
}
