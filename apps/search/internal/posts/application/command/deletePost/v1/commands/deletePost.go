package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/deletePost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type DeletePostCommand struct {
	dto.DeletePostDataDto
}

func NewDeletePostCommand(
	req any,
) (*DeletePostCommand, error) {
	typedCommand := new(dto.DeletePostDataDto)
	if err := dtoUtil.ValidateStruct(req, typedCommand); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewDeletePostCommand")
	}

	return &DeletePostCommand{
		DeletePostDataDto: *typedCommand,
	}, nil
}
