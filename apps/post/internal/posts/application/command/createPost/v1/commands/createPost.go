package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePostCommand struct {
	*dto.CreatePostRequest
}

func NewCreatePostCommand(
	req any,
) (*CreatePostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.CreatePostRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreatePostCommand")
	}

	return &CreatePostCommand{
		CreatePostRequest: typedReq,
	}, nil
}
