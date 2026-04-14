package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type IncrementPostViewsHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IIncrementPostViewsHandler interface {
	grpcTypes.GrpcHandler[*IncrementPostViewsCommand, *dto.IncrementPostViewsResponse]
}

func NewIncrementPostViewsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) IncrementPostViewsHandler {
	return IncrementPostViewsHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (h IncrementPostViewsHandler) Handle(
	ctx context.Context,
	cmd *IncrementPostViewsCommand,
) (*dto.IncrementPostViewsResponse, error) {
	// Get existing post
	post, err := h.postRepo.GetPostByID(ctx, cmd.IncrementPostViewsRequest.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	// Increment views
	post.Views += 1
	updates := map[string]interface{}{
		"views": post.Views,
	}

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.IncrementPostViewsRequest.ID, updates)
	if err != nil {
		h.log.Errorf("failed to increment post views: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	err = h.cacheRepo.DeleteAllPosts(ctx)
	if err != nil {
		h.log.Warnf("failed to delete all posts cache after incrementing post views: %v", err)
	}

	return dto.NewIncrementPostViewsResponse(id, true, "Post views incremented successfully", post.Views), nil
}
