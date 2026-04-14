package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePosts/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

// DeletePostsCommand for deleting multiple posts
type DeletePostsCommand struct {
	*dto.DeletePostsRequest
}

func NewDeletePostsCommand(
	req any,
) (*DeletePostsCommand, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.DeletePostsRequest{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewDeletePostsCommand")
	}

	return &DeletePostsCommand{
		DeletePostsRequest: typedReq,
	}, nil
}
