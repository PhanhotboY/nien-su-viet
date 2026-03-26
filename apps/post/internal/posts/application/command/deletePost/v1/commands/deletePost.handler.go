package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

// ============================================================
// DeletePostHandler - Single Delete
// ============================================================

type DeletePostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IDeletePostHandler interface {
	grpcTypes.GrpcHandler[*DeletePostCommand, *dto.DeletePostResponse]
}

func NewDeletePostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) DeletePostHandler {
	return DeletePostHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h DeletePostHandler) Handle(
	ctx context.Context,
	cmd *DeletePostCommand,
) (*dto.DeletePostResponse, error) {
	// Delete the post from repository
	id, err := h.postRepo.DeletePost(ctx, cmd.ID)
	if err != nil {
		h.log.Errorf("failed to delete post: %v", err)
		return nil, err
	}

	return dto.NewDeletePostResponse(id, true, "Post deleted successfully"), nil
}
