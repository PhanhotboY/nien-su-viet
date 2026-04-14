package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type IncrementPostLikesCommand struct {
	*dto.IncrementPostLikesRequest
}

func NewIncrementPostLikesCommand(
	req any,
) (*IncrementPostLikesCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.IncrementPostLikesRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewIncrementPostLikesCommand")
	}

	return &IncrementPostLikesCommand{
		IncrementPostLikesRequest: typedReq,
	}, nil
}
