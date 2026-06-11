package aqueries

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/listPlans/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type ListPlansQuery struct {
	*adto.ListPlansReqDto
}

func NewListPlansQuery(req any) (*ListPlansQuery, error) {
	typedReq := new(adto.ListPlansReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewListPlansQuery")
	}

	return &ListPlansQuery{
		ListPlansReqDto: typedReq,
	}, nil
}
