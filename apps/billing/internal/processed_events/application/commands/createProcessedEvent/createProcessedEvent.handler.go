package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreateProcessedEventHandler interface {
	grpcTypes.GrpcHandler[*CreateProcessedEventCommand, adto.CreateProcessedEventResDto]
}

type createProcessedEventHandler struct {
	logger             logger.Logger
	processedEventRepo drepo.ProcessedEventDbRepo
}

func NewCreateProcessedEventHandler(l logger.Logger, processedEventRepo drepo.ProcessedEventDbRepo) CreateProcessedEventHandler {
	return &createProcessedEventHandler{l, processedEventRepo}
}

func (h *createProcessedEventHandler) Handle(ctx context.Context, command *CreateProcessedEventCommand) (adto.CreateProcessedEventResDto, error) {
	h.logger.Info("Handling CreateProcessedEventCommand", "command", command.CreateProcessedEventReqDto)

	event := command.MapToEntity()

	err := h.processedEventRepo.CreateProcessedEvent(ctx, event)
	if err != nil {
		h.logger.Errorf("Failed to create processed event: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to create processed event", "CreateProcessedEventHandler")
	}

	h.logger.Infof("Processed event created successfully with ID: %s", event.ID)

	return adto.NewCreateProcessedEventResDto(event.ID.String(), true, "Processed event created successfully"), nil
}
