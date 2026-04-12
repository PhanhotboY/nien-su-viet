package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type IncrementPostViewsCommand struct {
	*dto.IncrementPostViewsRequest
}

func NewIncrementPostViewsCommand(
	req any,
) (*IncrementPostViewsCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.IncrementPostViewsRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewIncrementPostViewsCommand")
	}

	return &IncrementPostViewsCommand{
		IncrementPostViewsRequest: typedReq,
	}, nil
}
