package queries

import (
	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPostQuery struct {
	*dto.GetPostQueryReq
}

func NewGetPostQuery(
	req *dto.GetPostQueryReq,
) *GetPostQuery {
	return &GetPostQuery{
		GetPostQueryReq: req,
	}
}

func (g GetPostQuery) IsValidUUID() bool {
	return uuid.Validate(g.IDOrSlug) == nil
}

func NewGetPostQueryWithValidation(
	req any,
) (*GetPostQuery, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.GetPostQueryReq{}, nil)
	if err != nil {
		return nil, err
	}

	return NewGetPostQuery(typedReq), nil
}
