package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/dto"
	event "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/events"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/bus"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

// ============================================================
// DeletePostHandler - Single Delete
// ============================================================

type DeletePostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
	bus      bus.Bus
}

type IDeletePostHandler interface {
	grpcTypes.GrpcHandler[*DeletePostCommand, *dto.DeletePostResponse]
}

func NewDeletePostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	bus bus.Bus,
) DeletePostHandler {
	return DeletePostHandler{
		log:      log,
		postRepo: postRepo,
		bus:      bus,
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
		return nil, grpcerrors.ParseError(err)
	}

	if postDeletedEvent, err := event.NewPostDeletedEvent(id); err != nil {
		h.log.Errorf("failed to create post deleted event: %v", err)
	} else {
		h.bus.PublishMessage(ctx, postDeletedEvent)
	}

	return dto.NewDeletePostResponse(id, true, "Post deleted successfully"), nil
}
