package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type IUpdatePostHandler interface {
	consumer.ConsumerHandler
}

type updatePostHandler struct {
	logger     logger.Logger
	postEsRepo repository.PostEsRepository
	cacheRepo  repository.PostCacheRepository
	tracer     trace.Tracer
}

func NewUpdatePostHandler(
	l logger.Logger,
	postEsRepo repository.PostEsRepository,
	cacheRepo repository.PostCacheRepository,
	tracer trace.Tracer,
) IUpdatePostHandler {
	return &updatePostHandler{
		logger:     l,
		postEsRepo: postEsRepo,
		cacheRepo:  cacheRepo,
		tracer:     tracer,
	}
}

func (h *updatePostHandler) Handle(ctx context.Context, msgCtx types.MessageConsumeContext) error {
	command, err := NewUpdatePostCommand(msgCtx.Message().GetData())
	if err != nil {
		h.logger.Errorf("failed to parse UpdatePostCommand: %v", err)
		return err
	}
	h.logger.Debugf("handling RMQ message: %+v", command)

	if command.Id == "" {
		return nil // Or return an error indicating that the Post ID is required
	}
	foundPost, err := h.postEsRepo.GetPostByID(ctx, command.Id)
	if err != nil {
		return err
	}
	if foundPost == nil {
		return nil // Or return an error indicating that the post was not found
	}
	command.MapToEntity(foundPost)
	if err = h.postEsRepo.UpdatePost(ctx, command.Id, foundPost); err != nil {
		h.logger.Errorf("failed to update post in ES: %v", err)
		return err
	}

	if err = h.cacheRepo.DeleteAllPosts(ctx); err != nil {
		h.logger.Errorf("failed to invalidate post cache: %v", err)
		return err
	}

	return nil
}
