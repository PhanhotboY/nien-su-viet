package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type PublishPostCommand struct {
	*dto.PublishPostRequest
}

func NewPublishPostCommand(
	req any,
) (*PublishPostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.PublishPostRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewPublishPostCommand")
	}

	return &PublishPostCommand{
		PublishPostRequest: typedReq,
	}, nil
}
