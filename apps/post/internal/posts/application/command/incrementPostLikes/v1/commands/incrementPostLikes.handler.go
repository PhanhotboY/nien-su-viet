package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type IncrementPostLikesHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IIncrementPostLikesHandler interface {
	grpcTypes.GrpcHandler[*IncrementPostLikesCommand, *dto.IncrementPostLikesResponse]
}

func NewIncrementPostLikesHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) IncrementPostLikesHandler {
	return IncrementPostLikesHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (h IncrementPostLikesHandler) Handle(
	ctx context.Context,
	cmd *IncrementPostLikesCommand,
) (*dto.IncrementPostLikesResponse, error) {
	// Get existing post
	post, err := h.postRepo.GetPostByID(ctx, cmd.IncrementPostLikesRequest.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	// Increment likes
	post.Likes += 1
	updates := map[string]interface{}{
		"likes": post.Likes,
	}

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.IncrementPostLikesRequest.ID, updates)
	if err != nil {
		h.log.Errorf("failed to increment post likes: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	err = h.cacheRepo.DeleteAllPosts(ctx)
	if err != nil {
		h.log.Warnf("failed to delete all posts cache after incrementing post likes: %v", err)
	}

	return dto.NewIncrementPostLikesResponse(id, true, "Post likes incremented successfully", post.Likes), nil
}
