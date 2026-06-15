package acmd

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreateSubscriptionCommand struct {
	*adto.CreateSubscriptionReqDto
}

func NewCreateSubscriptionCommand(req any) (*CreateSubscriptionCommand, error) {
	typedReq := new(adto.CreateSubscriptionReqDto)
	err := dtoUtil.ValidateStruct(req, typedReq)
	if err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewCreateSubscriptionCommand")
	}

	return &CreateSubscriptionCommand{
		CreateSubscriptionReqDto: typedReq,
	}, nil
}
