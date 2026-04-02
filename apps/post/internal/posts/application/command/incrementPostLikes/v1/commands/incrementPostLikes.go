package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type IncrementPostLikesCommand struct {
	*dto.IncrementPostLikesRequest
}

func NewIncrementPostLikesCommand(
	req *dto.IncrementPostLikesRequest,
) *IncrementPostLikesCommand {
	return &IncrementPostLikesCommand{
		IncrementPostLikesRequest: req,
	}
}

func NewIncrementPostLikesCommandWithValidation(
	req any,
) (*IncrementPostLikesCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.IncrementPostLikesRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewIncrementPostLikesCommand(typedReq), nil
}
