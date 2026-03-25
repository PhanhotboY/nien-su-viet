package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type DeletePostsHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IDeletePostsHandler interface {
	grpcTypes.GrpcHandler[*DeletePostsCommand, *dto.DeletePostsResponse]
}

func NewDeletePostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) DeletePostsHandler {
	return DeletePostsHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h DeletePostsHandler) Handle(
	ctx context.Context,
	cmd *DeletePostsCommand,
) (*dto.DeletePostsResponse, error) {
	h.log.Info("handling delete posts command", "post_ids_count", len(cmd.PostIds))

	if len(cmd.PostIds) == 0 {
		return dto.NewDeletePostsResponse("", false, "No posts to delete"), nil
	}

	resultID := ""
	if len(cmd.PostIds) > 0 {
		resultID = cmd.PostIds[0]
	}

	// Delete the posts from repository
	id, err := h.postRepo.DeletePost(ctx, resultID)
	if err != nil {
		h.log.Errorf("failed to delete posts: %v", err)
		return nil, err
	}

	h.log.Infof("posts deleted successfully, total: %d", len(cmd.PostIds))

	return dto.NewDeletePostsResponse(id, true, "Posts deleted successfully"), nil
}
