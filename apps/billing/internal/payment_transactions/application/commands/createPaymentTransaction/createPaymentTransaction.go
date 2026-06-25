package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePaymentTransactionCmd struct {
	adto.CreatePaymentTransactionReqDto
}

func NewCreatePaymentTransactionCmd(req any) (*CreatePaymentTransactionCmd, error) {
	typedReq := new(adto.CreatePaymentTransactionReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreatePaymentTransactionCmd")
	}
	return &CreatePaymentTransactionCmd{
		CreatePaymentTransactionReqDto: *typedReq,
	}, nil
}
