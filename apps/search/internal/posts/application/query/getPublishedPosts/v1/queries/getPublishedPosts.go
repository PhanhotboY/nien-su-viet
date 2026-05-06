package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPublishedPosts/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPublishedPostsQuery struct {
	dto.GetPublicPostsQueryReq
}

func NewGetPublishedPostsQuery(
	req any,
) (*GetPublishedPostsQuery, error) {
	typedReq := new(dto.GetPublicPostsQueryReq)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPublishedPostsQuery")
	}

	return &GetPublishedPostsQuery{
		GetPublicPostsQueryReq: *typedReq,
	}, nil
}
