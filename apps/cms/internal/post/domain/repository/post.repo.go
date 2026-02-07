package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

type PostPagination struct {
	Limit  int
	Offset int
}

type PostQuery struct {
	Published bool
}

type PostRepository interface {
	GetPosts(ctx context.Context, query PostQuery, pagination PostPagination) ([]*entity.PostBrief, error)
	GetPostByID(ctx context.Context, postId string) (*entity.Post, error)
	GetPostBySlug(ctx context.Context, slug string) (*entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) (string, error)
	UpdatePost(ctx context.Context, postId string, post *entity.Post) (string, error)
	DeletePost(ctx context.Context, postId string) (string, error)
}
