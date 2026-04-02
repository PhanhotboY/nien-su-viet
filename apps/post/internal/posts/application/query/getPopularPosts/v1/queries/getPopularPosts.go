package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPopularPosts/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPopularPostsQuery struct {
	*dto.GetPopularPostsQueryReq
}

func NewGetPopularPostsQuery(
	req *dto.GetPopularPostsQueryReq,
) *GetPopularPostsQuery {
	return &GetPopularPostsQuery{
		GetPopularPostsQueryReq: req,
	}
}

func NewGetPopularPostsQueryWithValidation(
	req any,
) (*GetPopularPostsQuery, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.GetPopularPostsQueryReq{}, nil)
	if err != nil {
		return nil, err
	}

	return NewGetPopularPostsQuery(typedReq), nil
}
