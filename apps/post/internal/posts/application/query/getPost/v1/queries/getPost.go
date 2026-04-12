package queries

import (
	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPostQuery struct {
	*dto.GetPostQueryReq
}

func NewGetPostQuery(
	req any,
) (*GetPostQuery, error) {
	typedReq, err := dtoUtil.ValidateStruct(req, dto.GetPostQueryReq{})
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPostQuery")
	}

	return &GetPostQuery{
		GetPostQueryReq: typedReq,
	}, nil
}

func (g GetPostQuery) IsValidUUID() bool {
	return uuid.Validate(g.IDOrSlug) == nil
}
