package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchaseStatus/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdatePurchaseStatusCommand struct {
	*adto.UpdatePurchaseStatusReqDto
}

func NewUpdatePurchaseStatusCommand(req any) (*UpdatePurchaseStatusCommand, error) {
	typedReq := new(adto.UpdatePurchaseStatusReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdatePurchaseCommand")
	}

	return &UpdatePurchaseStatusCommand{
		UpdatePurchaseStatusReqDto: typedReq,
	}, nil
}
