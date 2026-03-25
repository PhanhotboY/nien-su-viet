package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type IncrementPostViewsCommand struct {
	*dto.IncrementPostViewsRequest
}

func NewIncrementPostViewsCommand(
	req *dto.IncrementPostViewsRequest,
) *IncrementPostViewsCommand {
	return &IncrementPostViewsCommand{
		IncrementPostViewsRequest: req,
	}
}

func NewIncrementPostViewsCommandWithValidation(
	req any,
) (*IncrementPostViewsCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.IncrementPostViewsRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewIncrementPostViewsCommand(typedReq), nil
}
