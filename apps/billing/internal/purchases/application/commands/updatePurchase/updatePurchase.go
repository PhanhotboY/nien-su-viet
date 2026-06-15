package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdatePurchaseCommand struct {
	*adto.UpdatePurchaseReqDto
}

func NewUpdatePurchaseCommand(req any) (*UpdatePurchaseCommand, error) {
	typedReq := new(adto.UpdatePurchaseReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdatePurchaseCommand")
	}

	return &UpdatePurchaseCommand{
		UpdatePurchaseReqDto: typedReq,
	}, nil
}
