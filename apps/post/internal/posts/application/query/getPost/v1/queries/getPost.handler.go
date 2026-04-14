package queries

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPostHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IGetPostHandler interface {
	grpcTypes.GrpcHandler[*GetPostQuery, *dto.GetPostRes]
}

func NewGetPostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) GetPostHandler {
	return GetPostHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (c GetPostHandler) Handle(
	ctx context.Context,
	query *GetPostQuery,
) (*dto.GetPostRes, error) {
	var post *entity.Post
	var err error

	// Try to get post from cache
	post, err = c.cacheRepo.GetPost(ctx, query.IDOrSlug)
	if err != nil {
		c.log.Warnf("failed to get post from cache: %v, fallback to db", err)
	}

	// If not in cache, fetch from database
	if post == nil {
		if query.IsValidUUID() {
			post, err = c.postRepo.GetPostByID(ctx, query.IDOrSlug)
		} else {
			post, err = c.postRepo.GetPostBySlug(ctx, query.IDOrSlug)
		}

		if err != nil {
			return nil, grpcerrors.ParseError(err)
		}

		// Cache the result
		if err := c.cacheRepo.PutPost(ctx, query.IDOrSlug, post); err != nil {
			c.log.Warnf("failed to cache post: %v", err)
		}
	}

	res := &dto.GetPostRes{}
	res.FromEntity(post)

	return res, nil
}
