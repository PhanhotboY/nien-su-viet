package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
)

type CreateInboxEventCmdHandler interface {
	types.GrpcHandler[CreateInboxEventCmd, *adto.CreateInboxEventResDto]
}

type createInboxEventCmdHandler struct {
	repo drepo.InBoxEventDbRepo
}

func NewCreateInboxEventCmdHandler(r drepo.InBoxEventDbRepo) CreateInboxEventCmdHandler {
	return &createInboxEventCmdHandler{
		repo: r,
	}
}

func (h *createInboxEventCmdHandler) Handle(
	ctx context.Context, cmd CreateInboxEventCmd,
) (*adto.CreateInboxEventResDto, error) {
	return nil, nil
}
