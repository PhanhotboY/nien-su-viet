package dto

import (
	"time"

	"github.com/google/uuid"
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

func (r *UpdatePostRequest) ToUpdateMap() map[string]interface{} {
	updates := make(map[string]interface{})
	// Always set ID
	if postID, err := uuid.Parse(r.ID); err == nil {
		updates["id"] = postID
	}

	// Set optional fields if provided
	if r.Title != nil {
		updates["title"] = *r.Title
	}

	if r.Slug != nil {
		updates["slug"] = *r.Slug
	}

	if r.Content != nil {
		updates["content"] = *r.Content
	}

	if r.Summary != nil {
		updates["summary"] = *r.Summary
	}

	if r.Thumbnail != nil {
		updates["thumbnail"] = *r.Thumbnail
	}

	if r.CategoryID != nil {
		if categoryID, err := uuid.Parse(*r.CategoryID); err == nil {
			updates["category_id"] = categoryID
		}
	}

	// Handle published flag and timestamp
	if r.Published != nil {
		if *r.Published {
			updates["published"] = *r.Published
			if r.PublishedAt != nil {
				updates["published_at"] = r.PublishedAt
			} else {
				now := time.Now()
				updates["published_at"] = &now
			}
		} else {
			updates["published"] = false
			updates["published_at"] = nil
		}
	}

	return updates
}
