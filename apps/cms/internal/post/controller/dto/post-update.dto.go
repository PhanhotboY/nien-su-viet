package dto

import (
	"github.com/google/uuid"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

type PostUpdateReq struct {
	Title      *string    `json:"title,omitempty" binding:"min=1,max=255" example:"Vietnam history timeline" doc:"Post title"`            // Post title
	Slug       *string    `json:"slug,omitempty" binding:"lowercase" example:"vietnam-history-timeline" doc:"URL slug"`                   // URL slug
	CategoryId *uuid.UUID `json:"category_id,omitempty" example:"7c2bcb4b-5a26-4f08-9dd0-4cbe797a1d2c" doc:"Category ID"`                 // Foreign key to category
	Thumbnail  *string    `json:"thumbnail,omitempty" binding:"url" example:"https://example.com/thumbnail.jpg" doc:"Thumbnail URL"`      // Thumbnail URL
	AuthorId   *string    `json:"author_id,omitempty" binding:"omitempty" example:"5d2d6baf-3f1c-4c67-9a73-3aaf0d5a2b1c" doc:"Author ID"` // Foreign key to author
	Content    *string    `json:"content,omitempty" doc:"Post content"`                                                                   // Post content
	Summary    *string    `json:"summary,omitempty" binding:"max=5000" doc:"Post summary"`                                                // Post summary
	Published  *bool      `json:"published,omitempty" doc:"Publication status"`                                                           // Publication status
}

func (dto *PostUpdateReq) MapToEntity() *entity.Post {
	if dto == nil {
		return nil
	}

	ent := &entity.Post{}
	if dto.Title != nil {
		ent.Title = *dto.Title
	}
	if dto.Slug != nil {
		ent.Slug = *dto.Slug
	}
	if dto.CategoryId != nil {
		ent.CategoryId = dto.CategoryId
	}
	if dto.Thumbnail != nil {
		ent.Thumbnail = dto.Thumbnail
	}
	if dto.Content != nil {
		ent.Content = *dto.Content
	}
	if dto.Summary != nil {
		ent.Summary = dto.Summary
	}
	if dto.Published != nil {
		ent.Published = *dto.Published
	}
	if dto.AuthorId != nil {
		ent.AuthorId = *dto.AuthorId
	}
	return ent
}
