package dto

import (
	"github.com/google/uuid"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

type PostCreateReq struct {
	Title      string     `json:"title" binding:"required,min=1,max=255" example:"Vietnam history timeline" doc:"Post title"`                  // Post title
	Slug       string     `json:"slug" binding:"required,lowercase" example:"vietnam-history-timeline" doc:"URL slug"`                         // URL slug
	CategoryId *uuid.UUID `json:"category_id,omitempty" binding:"omitempty" example:"7c2bcb4b-5a26-4f08-9dd0-4cbe797a1d2c" doc:"Category ID"`  // Foreign key to category
	Thumbnail  *string    `json:"thumbnail,omitempty" binding:"omitempty,url" example:"https://example.com/thumbnail.jpg" doc:"Thumbnail URL"` // Thumbnail URL
	AuthorId   string     `json:"author_id" binding:"required" example:"5d2d6baf-3f1c-4c67-9a73-3aaf0d5a2b1c" doc:"Author ID"`                 // Foreign key to author
	Content    string     `json:"content" binding:"required" doc:"Post content"`                                                               // Post content
	Summary    *string    `json:"summary,omitempty" binding:"omitempty,max=5000" doc:"Post summary"`                                           // Post summary
	Published  *bool      `json:"published,omitempty" binding:"omitempty" doc:"Publication status"`                                            // Publication status
}

func (dto *PostCreateReq) MapToEntity() *entity.Post {
	if dto == nil {
		return nil
	}

	published := false
	if dto.Published != nil {
		published = *dto.Published
	}

	return &entity.Post{
		Title:      dto.Title,
		Slug:       dto.Slug,
		CategoryId: dto.CategoryId,
		Thumbnail:  dto.Thumbnail,
		AuthorId:   dto.AuthorId,
		Content:    dto.Content,
		Summary:    dto.Summary,
		Published:  published,
	}
}
