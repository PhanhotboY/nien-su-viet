package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/updateSubscriptionStatus/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type UpdateSubscriptionStatusCommand struct {
	*adto.UpdateSubscriptionStatusReqDto
}

func NewUpdateSubscriptionStatusCommand(req any) (*UpdateSubscriptionStatusCommand, error) {
	typedReq := new(adto.UpdateSubscriptionStatusReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewUpdateSubscriptionCommand")
	}

	return &UpdateSubscriptionStatusCommand{
		UpdateSubscriptionStatusReqDto: typedReq,
	}, nil
}
