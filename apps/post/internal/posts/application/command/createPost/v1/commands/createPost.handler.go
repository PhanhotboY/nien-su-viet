package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"
	event "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/events"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/bus"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
	msgBus   bus.Bus
}

type ICreatePostHandler interface {
	grpcTypes.GrpcHandler[*CreatePostCommand, *dto.CreatePostResponse]
}

func NewCreatePostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	msgBus bus.Bus,
) CreatePostHandler {
	return CreatePostHandler{
		log:      log,
		postRepo: postRepo,
		msgBus:   msgBus,
	}
}

func (h CreatePostHandler) Handle(
	ctx context.Context,
	cmd *CreatePostCommand,
) (*dto.CreatePostResponse, error) {
	// Save to repository
	post, err := h.postRepo.CreatePost(ctx, cmd.MapToEntity())
	if err != nil {
		h.log.Errorf("failed to create post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	if postCreatedEvent, err := event.NewPostCreatedEvent(post); err != nil {
		h.log.Errorf("failed to create post created event: %v", err)
	} else {
		h.msgBus.PublishMessage(ctx, postCreatedEvent)
	}

	return dto.NewCreatePostResponse(post.Id.String(), true, "Post created successfully"), nil
}
