package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreatePaymentAttemptCommand struct {
	*adto.CreatePaymentAttemptReqDto
}

func NewCreatePaymentAttemptCommand(req any) (*CreatePaymentAttemptCommand, error) {
	typedReq := new(adto.CreatePaymentAttemptReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreatePaymentAttemptCommand")
	}

	return &CreatePaymentAttemptCommand{
		CreatePaymentAttemptReqDto: typedReq,
	}, nil
}
