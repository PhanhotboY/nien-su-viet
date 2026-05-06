package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPopularPosts/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPopularPostsQuery struct {
	dto.GetPopularPostsQueryReq
}

func NewGetPopularPostsQuery(
	req any,
) (*GetPopularPostsQuery, error) {
	typedReq := new(dto.GetPopularPostsQueryReq)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPopularPostsQuery")
	}

	return &GetPopularPostsQuery{
		GetPopularPostsQueryReq: *typedReq,
	}, nil
}
