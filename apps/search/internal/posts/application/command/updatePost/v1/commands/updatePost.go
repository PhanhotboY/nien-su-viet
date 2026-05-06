package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/updatePost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdatePostCommand struct {
	dto.UpdatePostDataDto
}

func NewUpdatePostCommand(
	req any,
) (*UpdatePostCommand, error) {
	typedCommand := new(dto.UpdatePostDataDto)
	if err := dtoUtil.ValidateStruct(req, typedCommand); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdatePostCommand")
	}

	return &UpdatePostCommand{
		UpdatePostDataDto: *typedCommand,
	}, nil
}
