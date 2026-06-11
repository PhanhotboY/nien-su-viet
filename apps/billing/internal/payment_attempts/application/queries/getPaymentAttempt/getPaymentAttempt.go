package aquery

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPaymentAttemptQuery struct {
	*adto.GetPaymentAttemptReqDto
}

func NewGetPaymentAttemptQuery(req any) (*GetPaymentAttemptQuery, error) {
	var typedReq = new(adto.GetPaymentAttemptReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPaymentAttemptQuery")
	}

	return &GetPaymentAttemptQuery{typedReq}, nil
}
