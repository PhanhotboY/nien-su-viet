package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type ICreatePostHandler interface {
	grpcTypes.GrpcHandler[*CreatePostCommand, *dto.CreatePostResponse]
}

func NewCreatePostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) CreatePostHandler {
	return CreatePostHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (h CreatePostHandler) Handle(
	ctx context.Context,
	cmd *CreatePostCommand,
) (*dto.CreatePostResponse, error) {
	// Save to repository
	id, err := h.postRepo.CreatePost(ctx, cmd.MapToEntity())
	if err != nil {
		h.log.Errorf("failed to create post: %v", err)
		return nil, err
	}

	return dto.NewCreatePostResponse(id, true, "Post created successfully"), nil
}
