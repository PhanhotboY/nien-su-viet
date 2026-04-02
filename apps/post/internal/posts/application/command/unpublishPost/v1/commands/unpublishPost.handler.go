package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/unpublishPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type UnpublishPostHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IUnpublishPostHandler interface {
	grpcTypes.GrpcHandler[*UnpublishPostCommand, *dto.UnpublishPostResponse]
}

func NewUnpublishPostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) UnpublishPostHandler {
	return UnpublishPostHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (h UnpublishPostHandler) Handle(
	ctx context.Context,
	cmd *UnpublishPostCommand,
) (*dto.UnpublishPostResponse, error) {

	// Get existing post
	post, err := h.postRepo.GetPostByID(ctx, cmd.UnpublishPostRequest.ID)
	if err != nil {
		h.log.Errorf("failed to get post: %v", err)
		return nil, err
	}

	// Update published status and clear timestamp
	post.Published = false
	post.PublishedAt = nil

	// Save to repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.UnpublishPostRequest.ID, post)
	if err != nil {
		h.log.Errorf("failed to unpublish post: %v", err)
		return nil, err
	}

	err = h.cacheRepo.DeleteAllPosts(ctx)
	if err != nil {
		h.log.Warnf("failed to delete all posts cache after unpublishing post: %v", err)
	}

	return dto.NewUnpublishPostResponse(id, true, "Post unpublished successfully"), nil
}
