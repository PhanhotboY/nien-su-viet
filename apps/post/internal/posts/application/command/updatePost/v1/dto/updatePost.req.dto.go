package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
)

// UpdatePostRequest is the request DTO for updating a post
type UpdatePostRequest struct {
	ID          string     `json:"id" validate:"required,uuid"`
	Title       *string    `json:"title" validate:"omitempty,min=1,max=255"`
	Slug        *string    `json:"slug" validate:"omitempty,min=1,max=255"`
	Content     *string    `json:"content" validate:"omitempty,min=1"`
	Summary     *string    `json:"summary" validate:"omitempty,max=500"`
	Thumbnail   *string    `json:"thumbnail" validate:"omitempty,url"`
	CategoryID  *string    `json:"category_id" validate:"omitempty,uuid"`
	Published   *bool      `json:"published"`
	PublishedAt *time.Time `json:"published_at"`
}

func (r *UpdatePostRequest) MapToEntity() *entity.Post {
	post := &entity.Post{}
	// Always set ID
	if postID, err := uuid.Parse(r.ID); err == nil {
		post.Id = postID
	}

	// Set optional fields if provided
	if r.Title != nil {
		post.Title = *r.Title
	}

	if r.Slug != nil {
		post.Slug = *r.Slug
	}

	if r.Content != nil {
		post.Content = *r.Content
	}

	if r.Summary != nil {
		post.Summary = r.Summary
	}

	if r.Thumbnail != nil {
		post.Thumbnail = r.Thumbnail
	}

	if r.CategoryID != nil {
		if categoryID, err := uuid.Parse(*r.CategoryID); err == nil {
			post.CategoryId = &categoryID
		}
	}

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
