package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPublishedPostsQuery struct {
	*dto.GetPublicPostsQueryReq
}

func NewGetPublishedPostsQuery(
	req *dto.GetPublicPostsQueryReq,
) *GetPublishedPostsQuery {
	return &GetPublishedPostsQuery{
		GetPublicPostsQueryReq: req,
	}
}

func NewGetPublishedPostsQueryWithValidation(
	req any,
) (*GetPublishedPostsQuery, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.GetPublicPostsQueryReq{}, nil)
	if err != nil {
		return nil, err
	}

	return NewGetPublishedPostsQuery(typedReq), nil
}
