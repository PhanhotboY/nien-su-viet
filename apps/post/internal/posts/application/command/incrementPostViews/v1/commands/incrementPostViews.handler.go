package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type IncrementPostViewsHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IIncrementPostViewsHandler interface {
	grpcTypes.GrpcHandler[*IncrementPostViewsCommand, *dto.IncrementPostViewsResponse]
}

func NewIncrementPostViewsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) IncrementPostViewsHandler {
	return IncrementPostViewsHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h IncrementPostViewsHandler) Handle(
	ctx context.Context,
	cmd *IncrementPostViewsCommand,
) (*dto.IncrementPostViewsResponse, error) {
	h.log.Info("handling increment post views command", "post_id", cmd.IncrementPostViewsRequest.ID)

	// Get existing post
	post, err := h.postRepo.GetPostByID(ctx, cmd.IncrementPostViewsRequest.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, err
	}

	// Increment views
	post.Views++

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.IncrementPostViewsRequest.ID, post)
	if err != nil {
		h.log.Errorf("failed to increment post views: %v", err)
		return nil, err
	}

	h.log.Infof("post views incremented successfully with id: %s, views: %d", cmd.IncrementPostViewsRequest.ID, post.Views)

	return dto.NewIncrementPostViewsResponse(id, true, "Post views incremented successfully", post.Views), nil
}
