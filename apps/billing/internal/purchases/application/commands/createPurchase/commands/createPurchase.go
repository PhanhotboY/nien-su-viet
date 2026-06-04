package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePurchaseCommand struct {
	*adto.CreatePurchaseReqDto
}

func NewCreatePurchaseCommand(req any) (*CreatePurchaseCommand, error) {
	typedReq := new(adto.CreatePurchaseReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreatePurchaseCommand")
	}

	return &CreatePurchaseCommand{
		CreatePurchaseReqDto: typedReq,
	}, nil
}
