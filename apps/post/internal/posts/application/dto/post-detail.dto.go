package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetPostRes represents the response DTO for getting a single post
type PostDetailDto struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Slug        string                 `json:"slug"`
	Content     string                 `json:"content"`
	Summary     *string                `json:"summary,omitempty"`
	Thumbnail   *string                `json:"thumbnail,omitempty"`
	AuthorID    string                 `json:"author_id"`
	CategoryID  *string                `json:"category_id,omitempty"`
	Views       int                    `json:"views"`
	Likes       int                    `json:"likes"`
	Published   bool                   `json:"published"`
	PublishedAt *timestamppb.Timestamp `json:"published_at,omitempty"`
	CreatedAt   timestamppb.Timestamp  `json:"created_at"`
	UpdatedAt   timestamppb.Timestamp  `json:"updated_at"`
}

func (p *PostDetailDto) FromEntity(post *entity.Post) {
	if p == nil || post == nil {
		return
	}

	p.ID = post.Id.String()
	p.Title = post.Title
	p.Slug = post.Slug
	p.Content = post.Content
	p.Summary = post.Summary
	p.Thumbnail = post.Thumbnail
	p.AuthorID = post.AuthorId

	if post.CategoryId != nil {
		categoryIDStr := post.CategoryId.String()
		p.CategoryID = &categoryIDStr
	}

	p.Views = post.Views
	p.Likes = post.Likes
	p.Published = post.Published

	if post.PublishedAt != nil {
		p.PublishedAt = timestamppb.New(*post.PublishedAt)
	}
	p.CreatedAt = *timestamppb.New(post.CreatedAt)
	p.UpdatedAt = *timestamppb.New(post.UpdatedAt)
}
