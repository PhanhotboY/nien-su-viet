package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/updateEvent/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdateEventCommand struct {
	dto.UpdateEventDataDto
}

func NewUpdateEventCommand(
	req any,
) (*UpdateEventCommand, error) {
	typedCommand := new(dto.UpdateEventDataDto)
	if err := dtoUtil.ValidateStruct(req, typedCommand); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdateEventCommand")
	}

	return &UpdateEventCommand{
		UpdateEventDataDto: *typedCommand,
	}, nil
}
