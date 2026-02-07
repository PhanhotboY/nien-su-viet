package service

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

type PostService interface {
	FindPosts(ctx context.Context, query *dto.FindPostsQueryReq) ([]*entity.PostBrief, error)
	GetPublishedPosts(ctx context.Context, query *dto.GetPublicPostsQueryReq) ([]*entity.PostBrief, error)
	FindPostByIdOrSlug(ctx context.Context, postId string) (*entity.Post, error)
	CreatePost(ctx context.Context, post *dto.PostCreateReq) (string, error)
	UpdatePost(ctx context.Context, postId string, post *dto.PostUpdateReq) (string, error)
	DeletePost(ctx context.Context, postId string) (string, error)
	PublishPost(ctx context.Context, postId string) (string, error)
	UnpublishPost(ctx context.Context, postId string) (string, error)
}
