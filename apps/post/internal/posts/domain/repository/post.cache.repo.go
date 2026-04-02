package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type PostCacheRepository interface {
	PutPost(ctx context.Context, key string, Post *entity.Post) error
	PutPosts(ctx context.Context, key string, Posts *utils.PaginatedResponse[entity.PostBrief]) error
	GetPost(ctx context.Context, key string) (*entity.Post, error)
	GetPosts(ctx context.Context, key string) (*utils.PaginatedResponse[entity.PostBrief], error)
	DeletePost(ctx context.Context, key string) error
	DeleteAllPosts(ctx context.Context) error
}
