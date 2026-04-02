package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
)

// CreatePostRequest is the request DTO for creating a post
type CreatePostRequest struct {
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Slug        string     `json:"slug" validate:"required,min=1,max=255"`
	Content     string     `json:"content" validate:"required,min=1"`
	Summary     *string    `json:"summary" validate:"omitempty,max=500"`
	Thumbnail   *string    `json:"thumbnail" validate:"omitempty,url"`
	AuthorId    string     `json:"author_id" validate:"required"`
	CategoryID  *string    `json:"category_id"`
	Published   *bool      `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
}

func (r *CreatePostRequest) MapToEntity() *entity.Post {
	post := &entity.Post{}

	post.Id = uuid.New()
	post.Title = r.Title
	post.Slug = r.Slug
	post.Content = r.Content
	post.Summary = r.Summary
	post.Thumbnail = r.Thumbnail
	post.AuthorId = r.AuthorId

	// Handle published flag and timestamp
	if r.Published != nil && *r.Published {
		post.Published = true
		if r.PublishedAt != nil {
			post.PublishedAt = r.PublishedAt
		} else {
			now := time.Now()
			post.PublishedAt = &now
		}
	} else if r.Published != nil && !*r.Published {
		post.Published = false
		post.PublishedAt = nil
	}

	return post
}
