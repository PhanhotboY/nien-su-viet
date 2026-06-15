package acmd

import (
	"context"

	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/application/commands/createOutboxEvent/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreateOutboxEventHandler interface {
	grpcTypes.GrpcHandler[*CreateOutboxEventCommand, adto.CreateOutboxEventResDto]
}

type createOutboxEventHandler struct {
	logger          logger.Logger
	outboxEventRepo drepo.OutboxEventDbRepo
}

func NewCreateOutboxEventHandler(l logger.Logger, outboxEventRepo drepo.OutboxEventDbRepo) CreateOutboxEventHandler {
	return &createOutboxEventHandler{l, outboxEventRepo}
}

func (h *createOutboxEventHandler) Handle(ctx context.Context, command *CreateOutboxEventCommand) (adto.CreateOutboxEventResDto, error) {
	h.logger.Info("Handling CreateOutboxEventCommand", "command", command.CreateOutboxEventReqDto)

	event := command.MapToEntity()

	err := h.outboxEventRepo.CreateOutboxEvent(ctx, event)
	if err != nil {
		h.logger.Errorf("Failed to create outbox event: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to create outbox event", "CreateOutboxEventHandler")
	}

	h.logger.Infof("Outbox event created successfully with ID: %s", event.ID)

	return adto.NewCreateOutboxEventResDto(event.ID.String(), true, "Outbox event created successfully"), nil
}
