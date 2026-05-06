package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/deleteEvent/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type DeleteEventCommand struct {
	dto.DeleteEventDataDto
}

func NewDeleteEventCommand(
	req any,
) (*DeleteEventCommand, error) {
	typedCommand := new(dto.DeleteEventDataDto)
	if err := dtoUtil.ValidateStruct(req, typedCommand); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewDeleteEventCommand")
	}

	return &DeleteEventCommand{
		DeleteEventDataDto: *typedCommand,
	}, nil
}
