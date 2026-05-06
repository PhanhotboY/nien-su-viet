package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/createPost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePostCommand struct {
	dto.CreatePostDataDto
}

func NewCreatePostCommand(
	req any,
) (*CreatePostCommand, error) {
	typedCommand := new(dto.CreatePostDataDto)
	if err := dtoUtil.ValidateStruct(req, typedCommand); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreatePostCommand")
	}

	return &CreatePostCommand{
		CreatePostDataDto: *typedCommand,
	}, nil
}
