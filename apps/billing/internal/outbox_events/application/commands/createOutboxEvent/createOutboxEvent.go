package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/application/commands/createOutboxEvent/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreateOutboxEventCommand struct {
	*adto.CreateOutboxEventReqDto
}

func NewCreateOutboxEventCommand(req any) (*CreateOutboxEventCommand, error) {
	typedReq := new(adto.CreateOutboxEventReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreateOutboxEventCommand")
	}

	return &CreateOutboxEventCommand{
		CreateOutboxEventReqDto: typedReq,
	}, nil
}
