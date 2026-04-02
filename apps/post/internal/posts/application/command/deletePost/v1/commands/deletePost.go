package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

// DeletePostCommand for deleting a single post
type DeletePostCommand struct {
	*dto.DeletePostRequest
}

func NewDeletePostCommand(
	req *dto.DeletePostRequest,
) *DeletePostCommand {
	return &DeletePostCommand{
		DeletePostRequest: req,
	}
}

func NewDeletePostCommandWithValidation(
	req any,
) (*DeletePostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.DeletePostRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewDeletePostCommand(typedReq), nil
}
