package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type ICreatePostHandler interface {
	consumer.ConsumerHandler
}

type createPostHandler struct {
	logger     logger.Logger
	postEsRepo repository.PostEsRepository
	cacheRepo  repository.PostCacheRepository
	tracer     trace.Tracer
}

func NewCreatePostHandler(
	l logger.Logger,
	postEsRepo repository.PostEsRepository,
	cacheRepo repository.PostCacheRepository,
	tracer trace.Tracer,
) ICreatePostHandler {
	return &createPostHandler{
		logger:     l,
		postEsRepo: postEsRepo,
		cacheRepo:  cacheRepo,
		tracer:     tracer,
	}
}

func (h *createPostHandler) Handle(ctx context.Context, msgCtx types.MessageConsumeContext) error {
	command, err := NewCreatePostCommand(msgCtx.Message().GetData())
	if err != nil {
		h.logger.Errorf("failed to parse CreatePostCommand: %v", err)
		return err
	}
	h.logger.Infof("handling RMQ message: %+v", command)

	if err = h.postEsRepo.IndexPost(ctx, command.Post); err != nil {
		h.logger.Errorf("failed to index post in ES: %v", err)
		return err
	}

	if err = h.cacheRepo.DeleteAllPosts(ctx); err != nil {
		h.logger.Errorf("failed to invalidate posts cache: %v", err)
		return err
	}

	return nil
}
