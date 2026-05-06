package commands

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

// UpdatePostCommand for updating a post
type UpdatePostCommand struct {
	*dto.UpdatePostRequest
}

func NewUpdatePostCommand(
	req any,
) (*UpdatePostCommand, error) {
	typedReq := new(dto.UpdatePostRequest)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdatePostCommand")
	}

	return &UpdatePostCommand{
		UpdatePostRequest: typedReq,
	}, nil
}
