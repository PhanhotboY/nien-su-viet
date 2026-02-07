package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

type PostBrief struct {
	ID         entity.PostId `json:"id" example:"0d8c0f2b-6f49-4a0c-9b7a-1f2f3c4d5e6f" doc:"Post ID"`                        // Post ID
	Title      string        `json:"title" example:"Vietnam history timeline" doc:"Post title"`                              // Post title
	Slug       string        `json:"slug" example:"vietnam-history-timeline" doc:"URL slug"`                                 // URL slug
	CategoryId *uuid.UUID    `json:"category_id,omitempty" example:"7c2bcb4b-5a26-4f08-9dd0-4cbe797a1d2c" doc:"Category ID"` // Foreign key to category
	Thumbnail  *string       `json:"thumbnail,omitempty" example:"https://example.com/thumbnail.jpg" doc:"Thumbnail URL"`    // Thumbnail URL
	AuthorId   string        `json:"author_id" example:"5d2d6baf-3f1c-4c67-9a73-3aaf0d5a2b1c" doc:"Author ID"`               // Foreign key to author
	Summary    *string       `json:"summary,omitempty" binding:"omitempty,max=5000" doc:"Post summary"`                      // Post summary
	Views      int           `json:"views" doc:"Number of views"`                                                            // Number of views
	Likes      int           `json:"likes" doc:"Number of likes"`                                                            // Number of likes
	Published  bool          `json:"published" doc:"Publication status"`                                                     // Publication status
	CreatedAt  time.Time     `json:"created_at" doc:"Creation timestamp"`                                                    // Creation timestamp
	UpdatedAt  time.Time     `json:"updated_at" doc:"Last update timestamp"`                                                 // Last update timestamp
}

func (dto *PostBrief) FromEntity(entity *entity.PostBrief) {
	if dto == nil || entity == nil {
		return
	}

	dto.ID = entity.Id
	dto.Title = entity.Title
	dto.Slug = entity.Slug
	dto.CategoryId = entity.CategoryId
	dto.Thumbnail = entity.Thumbnail
	dto.AuthorId = entity.AuthorId
	dto.Summary = entity.Summary
	dto.Views = entity.Views
	dto.Likes = entity.Likes
	dto.Published = entity.Published
	dto.CreatedAt = entity.CreatedAt
	dto.UpdatedAt = entity.UpdatedAt
}
