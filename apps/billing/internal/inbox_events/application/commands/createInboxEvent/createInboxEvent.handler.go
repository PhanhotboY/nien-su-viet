package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/createInboxEvent/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreateInboxEventCmdHandler interface {
	types.GrpcHandler[CreateInboxEventCmd, adto.CreateInboxEventResDto]
}

type createInboxEventCmdHandler struct {
	logger logger.Logger
	repo   drepo.InBoxEventDbRepo
}

func NewCreateInboxEventCmdHandler(l logger.Logger, r drepo.InBoxEventDbRepo) CreateInboxEventCmdHandler {
	return &createInboxEventCmdHandler{
		logger: l,
		repo:   r,
	}
}

func (h *createInboxEventCmdHandler) Handle(
	ctx context.Context, cmd CreateInboxEventCmd,
) (adto.CreateInboxEventResDto, error) {
	id, err := h.repo.Insert(ctx, cmd.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to insert inbox event: %v", err)
		return nil, err
	}

	return adto.NewCreateInboxEventResDto(id), nil
}
