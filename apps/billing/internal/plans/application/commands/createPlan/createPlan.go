package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePlanCmd struct {
	*adto.CreatePlanReqDto
}

func NewCreatePlanCommand(req any) (*CreatePlanCmd, error) {
	typedReq := new(adto.CreatePlanReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreatePlanCommand")
	}

	return &CreatePlanCmd{
		CreatePlanReqDto: typedReq,
	}, nil
}
