package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type ICreateEventHandler interface {
	consumer.ConsumerHandler
}

type createEventHandler struct {
	logger       logger.Logger
	heventEsRepo repository.HistoricalEventEsRepository
	cacheRepo    repository.HistoricalEventCacheRepository
	tracer       trace.Tracer
}

func NewCreateEventHandler(
	l logger.Logger,
	heventEsRepo repository.HistoricalEventEsRepository,
	cacheRepo repository.HistoricalEventCacheRepository,
	tracer trace.Tracer,
) ICreateEventHandler {
	return &createEventHandler{
		logger:       l,
		heventEsRepo: heventEsRepo,
		cacheRepo:    cacheRepo,
		tracer:       tracer,
	}
}

func (h *createEventHandler) Handle(ctx context.Context, msgCtx types.MessageConsumeContext) error {
	command, err := NewCreateEventCommand(msgCtx.Message().GetData())
	if err != nil {
		h.logger.Errorf("failed to parse CreateEventCommand: %v", err)
		return err
	}
	h.logger.Debugf("handling RMQ message: %+v", command)

	if err = h.heventEsRepo.IndexHistoricalEvent(ctx, command.HistoricalEvent); err != nil {
		h.logger.Errorf("failed to index historical event in ES: %v", err)
		return err
	}

	if err = h.cacheRepo.DeleteAllHistoricalEvents(ctx); err != nil {
		h.logger.Errorf("failed to invalidate historical events cache: %v", err)
		return err
	}

	return nil
}
