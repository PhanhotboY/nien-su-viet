package aqueries

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPlanByIdQuery struct {
	*adto.GetPlanByIdReqDto
}

func NewGetPlanByIdQuery(req any) (*GetPlanByIdQuery, error) {
	typedReq := new(adto.GetPlanByIdReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPlanByIdQuery")
	}

	return &GetPlanByIdQuery{
		GetPlanByIdReqDto: typedReq,
	}, nil
}
