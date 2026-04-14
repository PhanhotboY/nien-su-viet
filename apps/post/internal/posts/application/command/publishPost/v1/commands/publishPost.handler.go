package commands

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type PublishPostHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IPublishPostHandler interface {
	grpcTypes.GrpcHandler[*PublishPostCommand, *dto.PublishPostResponse]
}

func NewPublishPostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) PublishPostHandler {
	return PublishPostHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (h PublishPostHandler) Handle(
	ctx context.Context,
	cmd *PublishPostCommand,
) (*dto.PublishPostResponse, error) {
	// Get existing post
	_, err := h.postRepo.GetPostByID(ctx, cmd.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	// Update published status and timestamp
	now := time.Now()
	updates := map[string]interface{}{
		"published":    true,
		"published_at": &now,
	}

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.ID, updates)
	if err != nil {
		h.log.Errorf("failed to publish post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	err = h.cacheRepo.DeleteAllPosts(ctx)
	if err != nil {
		h.log.Warnf("failed to delete all posts cache after publishing post: %v", err)
	}

	return dto.NewPublishPostResponse(id, true, "Post published successfully"), nil
}
