package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePosts/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

// DeletePostsCommand for deleting multiple posts
type DeletePostsCommand struct {
	*dto.DeletePostsRequest
}

func NewDeletePostsCommand(
	req *dto.DeletePostsRequest,
) *DeletePostsCommand {
	return &DeletePostsCommand{
		DeletePostsRequest: req,
	}
}

func NewDeletePostsCommandWithValidation(
	req any,
) (*DeletePostsCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.DeletePostsRequest{}, nil)
	if err != nil {
		return nil, err
	}

	return NewDeletePostsCommand(typedReq), nil
}
