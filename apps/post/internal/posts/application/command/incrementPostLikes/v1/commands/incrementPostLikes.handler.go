package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type IncrementPostLikesHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IIncrementPostLikesHandler interface {
	grpcTypes.GrpcHandler[*IncrementPostLikesCommand, *dto.IncrementPostLikesResponse]
}

func NewIncrementPostLikesHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) IncrementPostLikesHandler {
	return IncrementPostLikesHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h IncrementPostLikesHandler) Handle(
	ctx context.Context,
	cmd *IncrementPostLikesCommand,
) (*dto.IncrementPostLikesResponse, error) {
	h.log.Info("handling increment post likes command", "post_id", cmd.IncrementPostLikesRequest.ID)

	// Get existing post
	post, err := h.postRepo.GetPostByID(ctx, cmd.IncrementPostLikesRequest.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, err
	}

	// Increment likes
	post.Likes++

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.IncrementPostLikesRequest.ID, post)
	if err != nil {
		h.log.Errorf("failed to increment post likes: %v", err)
		return nil, err
	}

	h.log.Infof("post likes incremented successfully with id: %s, likes: %d", cmd.IncrementPostLikesRequest.ID, post.Likes)

	return dto.NewIncrementPostLikesResponse(id, true, "Post likes incremented successfully", post.Likes), nil
}
