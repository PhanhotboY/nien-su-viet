package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getAllPosts/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetAllPostsQuery struct {
	*dto.GetAllPostsQueryReq
}

func NewGetAllPostsQuery(
	req *dto.GetAllPostsQueryReq,
) *GetAllPostsQuery {
	return &GetAllPostsQuery{
		GetAllPostsQueryReq: req,
	}
}

func NewGetAllPostsQueryWithValidation(
	req any,
) (*GetAllPostsQuery, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.GetAllPostsQueryReq{}, nil)
	if err != nil {
		return nil, err
	}

	return NewGetAllPostsQuery(typedReq), nil
}
