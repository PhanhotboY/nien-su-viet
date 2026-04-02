package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePostCommand struct {
	*dto.CreatePostRequest
}

func NewCreatePostCommand(
	req *dto.CreatePostRequest,
) *CreatePostCommand {
	return &CreatePostCommand{
		CreatePostRequest: req,
	}
}

func NewCreatePostCommandWithValidation(
	req any,
) (*CreatePostCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.CreatePostRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewCreatePostCommand(typedReq), nil
}
