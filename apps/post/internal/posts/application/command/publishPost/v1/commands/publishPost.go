package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type PublishPostCommand struct {
	*dto.PublishPostRequest
}

func NewPublishPostCommand(
	req *dto.PublishPostRequest,
) *PublishPostCommand {
	return &PublishPostCommand{
		PublishPostRequest: req,
	}
}

func NewPublishPostCommandWithValidation(
	req any,
) (*PublishPostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.PublishPostRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewPublishPostCommand(typedReq), nil
}
