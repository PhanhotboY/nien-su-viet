package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

// UpdatePostCommand for updating a post
type UpdatePostCommand struct {
	*dto.UpdatePostRequest
}

func NewUpdatePostCommand(
	req *dto.UpdatePostRequest,
) *UpdatePostCommand {
	return &UpdatePostCommand{
		UpdatePostRequest: req,
	}
}

func NewUpdatePostCommandWithValidation(
	req any,
) (*UpdatePostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.UpdatePostRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewUpdatePostCommand(typedReq), nil
}
