package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

// DeletePostCommand for deleting a single post
type DeletePostCommand struct {
	*dto.DeletePostRequest
}

func NewDeletePostCommand(
	req any,
) (*DeletePostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.DeletePostRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewDeletePostCommand")
	}

	return &DeletePostCommand{
		DeletePostRequest: typedReq,
	}, nil
}
