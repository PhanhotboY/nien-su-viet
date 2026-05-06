package repository

import (
	"context"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/entity"
)

type PostQuery struct {
	types.QueryVariant

	Page      uint32 `json:"page,omitempty" validate:"omitempty,gte=1"`
	Limit     uint32 `json:"limit,omitempty" validate:"omitempty,gte=1,lte=1000"`
	SortBy    string `json:"sort_by,omitempty" validate:"omitempty,max=100"`
	SortOrder string `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"` // "asc" or "desc"
}

type PostSearchResponse struct {
	Posts      []entity.PostBrief `json:"posts"`
	TotalCount int64              `json:"total_count"`
}

type PostEsRepository interface {
	IndexPost(ctx context.Context, post entity.Post) error
	SearchPosts(ctx context.Context, query PostQuery) (*PostSearchResponse, error)
	GetPostByID(ctx context.Context, postId entity.PostId) (*entity.Post, error)
	GetPostBySlug(ctx context.Context, slug string) (*entity.Post, error)
	UpdatePost(ctx context.Context, postId entity.PostId, post *entity.Post) error
	DeletePost(ctx context.Context, postId entity.PostId) error
	DeleteAllPosts(ctx context.Context) error
}
