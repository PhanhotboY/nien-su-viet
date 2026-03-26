package dto

import (
	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

type PostBriefDto struct {
	Id          string     `json:"id"`                     // Primary key
	Title       string     `json:"title"`                  // Post title
	Slug        string     `json:"slug"`                   // URL slug
	Summary     *string    `json:"summary,omitempty"`      // Post summary
	Thumbnail   *string    `json:"thumbnail,omitempty"`    // Thumbnail URL
	AuthorId    string     `json:"author_id"`              // Foreign key to author
	CategoryId  *uuid.UUID `json:"category_id"`            // Foreign key to category
	Views       int        `json:"views"`                  // Number of views
	Likes       int        `json:"likes"`                  // Number of likes
	Published   bool       `json:"published"`              // Publication status
	PublishedAt *string    `json:"published_at,omitempty"` // Publication timestamp
	CreatedAt   string     `json:"created_at"`             // Creation timestamp
	UpdatedAt   string     `json:"updated_at"`             // Last update timestamp
}

func (p *PostBriefDto) FromEntity(post *entity.PostBrief) {
	p.Id = post.Id.String()
	p.Title = post.Title
	p.Thumbnail = post.Thumbnail
	p.CategoryId = post.CategoryId
	if post.Summary != nil {
		p.Summary = post.Summary
	}
	p.Published = post.Published
	if post.PublishedAt != nil {
		publishedAtStr := grpcUtils.TimeToString(*post.PublishedAt)
		p.PublishedAt = &publishedAtStr
	}
	p.Slug = post.Slug
	p.AuthorId = post.AuthorId
	p.Views = post.Views
	p.Likes = post.Likes
	p.CreatedAt = grpcUtils.TimeToString(post.CreatedAt)
	p.UpdatedAt = grpcUtils.TimeToString(post.UpdatedAt)
}
