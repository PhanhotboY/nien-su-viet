package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreateProcessedEventCommand struct {
	*adto.CreateProcessedEventReqDto
}

func NewCreateProcessedEventCommand(req any) (*CreateProcessedEventCommand, error) {
	typedReq := new(adto.CreateProcessedEventReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreateProcessedEventCommand")
	}

	return &CreateProcessedEventCommand{
		CreateProcessedEventReqDto: typedReq,
	}, nil
}
