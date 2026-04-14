package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type DeletePostsHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IDeletePostsHandler interface {
	grpcTypes.GrpcHandler[*DeletePostsCommand, *dto.DeletePostsResponse]
}

func NewDeletePostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) DeletePostsHandler {
	return DeletePostsHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (h DeletePostsHandler) Handle(
	ctx context.Context,
	cmd *DeletePostsCommand,
) (*dto.DeletePostsResponse, error) {
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
		return nil, grpcerrors.ParseError(err)
	}

	err = h.cacheRepo.DeleteAllPosts(ctx)
	if err != nil {
		h.log.Warnf("failed to delete all posts cache after deleting posts: %v", err)
	}

	return dto.NewDeletePostsResponse(id, true, "Posts deleted successfully"), nil
}
