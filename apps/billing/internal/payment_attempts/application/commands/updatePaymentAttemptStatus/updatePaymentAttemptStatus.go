package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdatePaymentAttemptStatusCommand struct {
	*adto.UpdatePaymentAttemptStatusReqDto
}

func NewUpdatePaymentAttemptStatusCommand(req any) (*UpdatePaymentAttemptStatusCommand, error) {
	typedReq := new(adto.UpdatePaymentAttemptStatusReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdatePaymentAttemptStatusCommand")
	}

	return &UpdatePaymentAttemptStatusCommand{
		UpdatePaymentAttemptStatusReqDto: typedReq,
	}, nil
}
