package dto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/entity"
)

type UpdatePostDataDto struct {
	Id          entity.PostId `json:"id"`                     // Primary key
	Title       *string       `json:"title"`                  // Post title
	Slug        *string       `json:"slug"`                   // URL slug
	Content     *string       `json:"content"`                // Post content
	Summary     *string       `json:"summary,omitempty"`      // Post summary
	Thumbnail   *string       `json:"thumbnail,omitempty"`    // Thumbnail URL
	AuthorId    *string       `json:"author_id"`              // Foreign key to author
	CategoryId  *string       `json:"category_id"`            // Foreign key to category
	Views       *int          `json:"views"`                  // Number of views
	Likes       *int          `json:"likes"`                  // Number of likes
	Published   *bool         `json:"published"`              // Publication status
	PublishedAt *time.Time    `json:"published_at,omitempty"` // Publication timestamp
	CreatedAt   *time.Time    `json:"created_at"`             // Creation timestamp
	UpdatedAt   *time.Time    `json:"updated_at"`             // Last update timestamp
}

func (d *UpdatePostDataDto) MapToEntity(existingPost *entity.Post) {
	if d.Title != nil {
		existingPost.Title = *d.Title
	}
	if d.Slug != nil {
		existingPost.Slug = *d.Slug
	}
	if d.Content != nil {
		existingPost.Content = *d.Content
	}
	if d.Summary != nil {
		existingPost.Summary = d.Summary
	}
	if d.Thumbnail != nil {
		existingPost.Thumbnail = d.Thumbnail
	}
	if d.AuthorId != nil {
		existingPost.AuthorId = *d.AuthorId
	}
	if d.CategoryId != nil {
		existingPost.CategoryId = d.CategoryId
	}
	if d.Views != nil {
		existingPost.Views = *d.Views
	}
	if d.Likes != nil {
		existingPost.Likes = *d.Likes
	}
	if d.Published != nil {
		existingPost.Published = *d.Published
	}
	if d.PublishedAt != nil {
		existingPost.PublishedAt = d.PublishedAt
	}
	if d.CreatedAt != nil {
		existingPost.CreatedAt = *d.CreatedAt
	}
	if d.UpdatedAt != nil {
		existingPost.UpdatedAt = *d.UpdatedAt
	}
}
