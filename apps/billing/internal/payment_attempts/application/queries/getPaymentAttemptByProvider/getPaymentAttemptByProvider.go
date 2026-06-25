package aquery

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttemptByProvider/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPaymentAttemptByProviderQuery struct {
	*adto.GetPaymentAttemptByProviderReqDto
}

func NewGetPaymentAttemptByProviderQuery(req any) (*GetPaymentAttemptByProviderQuery, error) {
	var typedReq = new(adto.GetPaymentAttemptByProviderReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPaymentAttemptByProviderQuery")
	}

	return &GetPaymentAttemptByProviderQuery{typedReq}, nil
}
