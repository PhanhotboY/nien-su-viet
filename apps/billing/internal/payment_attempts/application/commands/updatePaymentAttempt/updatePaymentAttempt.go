package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttempt/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdatePaymentAttemptCommand struct {
	*adto.UpdatePaymentAttemptReqDto
}

func NewUpdatePaymentAttemptCommand(req any) (*UpdatePaymentAttemptCommand, error) {
	typedReq := new(adto.UpdatePaymentAttemptReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdatePaymentAttemptCommand")
	}

	return &UpdatePaymentAttemptCommand{
		UpdatePaymentAttemptReqDto: typedReq,
	}, nil
}
