package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/unpublishPost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UnpublishPostCommand struct {
	*dto.UnpublishPostRequest
}

func NewUnpublishPostCommand(
	req any,
) (*UnpublishPostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.UnpublishPostRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUnpublishPostCommand")
	}

	return &UnpublishPostCommand{
		UnpublishPostRequest: typedReq,
	}, nil
}
