package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"
	event "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/events"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/bus"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

// ============================================================
// UpdatePostHandler
// ============================================================

type UpdatePostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
	bus      bus.Bus
}

type IUpdatePostHandler interface {
	grpcTypes.GrpcHandler[*UpdatePostCommand, *dto.UpdatePostResponse]
}

func NewUpdatePostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	bus bus.Bus,
) UpdatePostHandler {
	return UpdatePostHandler{
		log:      log,
		postRepo: postRepo,
		bus:      bus,
	}
}

func (h UpdatePostHandler) Handle(
	ctx context.Context,
	cmd *UpdatePostCommand,
) (*dto.UpdatePostResponse, error) {
	// Update in repository
	id, err := h.postRepo.UpdatePost(ctx, cmd.ID, cmd.ToUpdateMap())
	if err != nil {
		h.log.Errorf("failed to update post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}

	post, err := h.postRepo.GetPostByID(ctx, id)
	if err != nil {
		h.log.Errorf("failed to get updated post: %v", err)
		return nil, grpcerrors.ParseError(err)
	}
	if postUpdatedEvent, err := event.NewPostUpdatedEvent(*post); err != nil {
		h.log.Errorf("failed to create post updated event: %v", err)
		return nil, grpcerrors.ParseError(err)
	} else {
		h.bus.PublishMessage(ctx, postUpdatedEvent)
	}

	return dto.NewUpdatePostResponse(id, true, "Post updated successfully"), nil
}
