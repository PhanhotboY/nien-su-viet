package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

// ============================================================
// UpdatePostHandler
// ============================================================

type UpdatePostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IUpdatePostHandler interface {
	grpcTypes.GrpcHandler[*UpdatePostCommand, *dto.UpdatePostResponse]
}

func NewUpdatePostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) UpdatePostHandler {
	return UpdatePostHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h UpdatePostHandler) Handle(
	ctx context.Context,
	cmd *UpdatePostCommand,
) (*dto.UpdatePostResponse, error) {
	h.log.Info("handling update post command", "post_id", cmd.ID)

	// Update in repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.ID, cmd.MapToEntity())
	if err != nil {
		h.log.Errorf("failed to update post: %v", err)
		return nil, err
	}

	h.log.Infof("post updated successfully with id: %s", id)

	return dto.NewUpdatePostResponse(id, true, "Post updated successfully"), nil
}
