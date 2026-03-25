package dto

import "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"

type PostBriefDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Slug        string `json:"slug"`
	AuthorID    string `json:"author_id"`
	PublishedAt string `json:"published_at,omitempty"`
}

func (p *PostBriefDto) FromEntity(post *entity.PostBrief) {
	p.ID = post.Id.String()
	p.Title = post.Title
	if post.Summary != nil {
		p.Summary = *post.Summary
	}
	p.Slug = post.Slug
	p.AuthorID = post.AuthorId
}
