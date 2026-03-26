package repository

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
)

type PostPagination struct {
	Limit  uint32
	Offset uint32
}

type PostQuery struct {
	Published     *bool
	Search        string
	SortBy        string
	SortOrder     string
	AuthorID      string
	CreatedAtFrom *time.Time
	CreatedAtTo   *time.Time
	UpdatedAtFrom *time.Time
	UpdatedAtTo   *time.Time
}

type PostRepository interface {
	GetPosts(ctx context.Context, query PostQuery, pagination PostPagination) ([]*entity.PostBrief, error)
	GetPostByID(ctx context.Context, postId string) (*entity.Post, error)
	GetPostBySlug(ctx context.Context, slug string) (*entity.Post, error)
	CountPosts(ctx context.Context, query PostQuery) (uint32, error)

	CreatePost(ctx context.Context, post *entity.Post) (string, error)
	UpdatePost(ctx context.Context, postId string, post *entity.Post) (string, error)
	DeletePost(ctx context.Context, postId string) (string, error)
}
