package commands

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type IUpdateEventHandler interface {
	consumer.ConsumerHandler
}

type updateEventHandler struct {
	logger       logger.Logger
	heventEsRepo repository.HistoricalEventEsRepository
	cacheRepo    repository.HistoricalEventCacheRepository
	tracer       trace.Tracer
}

func NewUpdateEventHandler(
	l logger.Logger,
	heventEsRepo repository.HistoricalEventEsRepository,
	cacheRepo repository.HistoricalEventCacheRepository,
	tracer trace.Tracer,
) IUpdateEventHandler {
	return &updateEventHandler{
		logger:       l,
		heventEsRepo: heventEsRepo,
		cacheRepo:    cacheRepo,
		tracer:       tracer,
	}
}

func (h *updateEventHandler) Handle(ctx context.Context, msgCtx types.MessageConsumeContext) error {
	command, err := NewUpdateEventCommand(msgCtx.Message().GetData())
	if err != nil {
		h.logger.Errorf("failed to parse UpdateEventCommand: %v", err)
		return err
	}
	h.logger.Debugf("handling RMQ message: %+v", command)

	if command.Id == "" {
		return nil // Or return an error indicating that the event ID is required
	}
	foundEvent, err := h.heventEsRepo.GetHistoricalEventById(ctx, command.Id)
	if err != nil {
		return err
	}
	if foundEvent == nil {
		return nil // Or return an error indicating that the event was not found
	}
	command.MapToEntity(foundEvent)
	if err = h.heventEsRepo.UpdateHistoricalEvent(ctx, command.Id, foundEvent); err != nil {
		h.logger.Errorf("failed to update historical event in ES: %v", err)
		return err
	}

	if err = h.cacheRepo.DeleteAllHistoricalEvents(ctx); err != nil {
		h.logger.Errorf("failed to invalidate historical events cache: %v", err)
		return err
	}

	return nil
}
