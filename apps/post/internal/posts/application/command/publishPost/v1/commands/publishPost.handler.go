package commands

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type PublishPostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IPublishPostHandler interface {
	grpcTypes.GrpcHandler[*PublishPostCommand, *dto.PublishPostResponse]
}

func NewPublishPostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) PublishPostHandler {
	return PublishPostHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h PublishPostHandler) Handle(
	ctx context.Context,
	cmd *PublishPostCommand,
) (*dto.PublishPostResponse, error) {
	h.log.Info("handling publish post command", "post_id", cmd.ID)

	// Get existing post
	post, err := h.postRepo.GetPostByID(ctx, cmd.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, err
	}

	// Update published status and timestamp
	post.Published = true
	now := time.Now()
	post.PublishedAt = &now

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.ID, post)
	if err != nil {
		h.log.Errorf("failed to publish post: %v", err)
		return nil, err
	}

	h.log.Infof("post published successfully with id: %s", cmd.ID)

	return dto.NewPublishPostResponse(id, true, "Post published successfully"), nil
}
