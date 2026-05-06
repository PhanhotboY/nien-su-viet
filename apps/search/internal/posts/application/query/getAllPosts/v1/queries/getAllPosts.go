package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getAllPosts/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetAllPostsQuery struct {
	dto.GetAllPostsQueryReq
}

func NewGetAllPostsQuery(
	req any,
) (*GetAllPostsQuery, error) {
	typedReq := dto.NewGetAllPostsQueryReqWithDefaultValue()
	if err := dtoUtil.ValidateStruct(req, &typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetAllPostsQuery")
	}

	return &GetAllPostsQuery{
		GetAllPostsQueryReq: typedReq,
	}, nil
}
