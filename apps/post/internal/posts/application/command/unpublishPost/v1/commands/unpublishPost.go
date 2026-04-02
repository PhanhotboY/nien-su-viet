package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/unpublishPost/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UnpublishPostCommand struct {
	*dto.UnpublishPostRequest
}

func NewUnpublishPostCommand(
	req *dto.UnpublishPostRequest,
) *UnpublishPostCommand {
	return &UnpublishPostCommand{
		UnpublishPostRequest: req,
	}
}

func NewUnpublishPostCommandWithValidation(
	req any,
) (*UnpublishPostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.UnpublishPostRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewUnpublishPostCommand(typedReq), nil
}
