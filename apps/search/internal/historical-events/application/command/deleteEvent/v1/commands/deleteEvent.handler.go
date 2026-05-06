package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type IDeleteEventHandler interface {
	consumer.ConsumerHandler
}

type deleteEventHandler struct {
	logger       logger.Logger
	heventEsRepo repository.HistoricalEventEsRepository
	cacheRepo    repository.HistoricalEventCacheRepository
	tracer       trace.Tracer
}

func NewDeleteEventHandler(
	l logger.Logger,
	heventEsRepo repository.HistoricalEventEsRepository,
	cacheRepo repository.HistoricalEventCacheRepository,
	tracer trace.Tracer,
) IDeleteEventHandler {
	return &deleteEventHandler{
		logger:       l,
		heventEsRepo: heventEsRepo,
		cacheRepo:    cacheRepo,
		tracer:       tracer,
	}
}

func (h *deleteEventHandler) Handle(ctx context.Context, msgCtx types.MessageConsumeContext) error {
	command, err := NewDeleteEventCommand(msgCtx.Message().GetData())
	if err != nil {
		h.logger.Errorf("failed to parse DeleteEventCommand: %v", err)
		return err
	}
	h.logger.Debugf("handling RMQ message: %+v", command)

	if err = h.heventEsRepo.DeleteHistoricalEvent(ctx, command.Id); err != nil {
		h.logger.Errorf("failed to delete historical event in ES: %v", err)
		return err
	}

	if err = h.cacheRepo.DeleteAllHistoricalEvents(ctx); err != nil {
		h.logger.Errorf("failed to invalidate historical events cache: %v", err)
		return err
	}

	return nil
}
