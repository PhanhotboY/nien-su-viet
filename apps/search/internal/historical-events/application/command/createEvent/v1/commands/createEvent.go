package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/createEvent/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreateEventCommand struct {
	dto.CreateEventDataDto
}

func NewCreateEventCommand(
	req any,
) (*CreateEventCommand, error) {
	typedCommand := new(dto.CreateEventDataDto)
	if err := dtoUtil.ValidateStruct(req, typedCommand); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreateEventCommand")
	}

	return &CreateEventCommand{
		CreateEventDataDto: *typedCommand,
	}, nil
}
