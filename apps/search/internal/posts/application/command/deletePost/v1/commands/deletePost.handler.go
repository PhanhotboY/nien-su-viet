package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type IDeletePostHandler interface {
	consumer.ConsumerHandler
}

type deletePostHandler struct {
	logger     logger.Logger
	postEsRepo repository.PostEsRepository
	cacheRepo  repository.PostCacheRepository
	tracer     trace.Tracer
}

func NewDeletePostHandler(
	l logger.Logger,
	postEsRepo repository.PostEsRepository,
	cacheRepo repository.PostCacheRepository,
	tracer trace.Tracer,
) IDeletePostHandler {
	return &deletePostHandler{
		logger:     l,
		postEsRepo: postEsRepo,
		cacheRepo:  cacheRepo,
		tracer:     tracer,
	}
}

func (h *deletePostHandler) Handle(ctx context.Context, msgCtx types.MessageConsumeContext) error {
	command, err := NewDeletePostCommand(msgCtx.Message().GetData())
	if err != nil {
		h.logger.Errorf("failed to parse DeletePostCommand: %v", err)
		return err
	}
	h.logger.Debugf("handling RMQ message: %+v", command)

	if err = h.postEsRepo.DeletePost(ctx, command.Id); err != nil {
		h.logger.Errorf("failed to delete historical event in ES: %v", err)
		return err
	}

	if err = h.cacheRepo.DeleteAllPosts(ctx); err != nil {
		h.logger.Errorf("failed to invalidate posts cache: %v", err)
		return err
	}

	return nil
}
