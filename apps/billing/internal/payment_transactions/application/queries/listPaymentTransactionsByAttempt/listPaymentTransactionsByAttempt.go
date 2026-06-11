package aqueries

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/queries/listPaymentTransactionsByAttempt/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type ListPaymentTransactionsByAttemptQuery struct {
	*adto.ListPaymentTransactionsByAttemptReqDto
}

func NewListPaymentTransactionsByAttemptQuery(req any) (*ListPaymentTransactionsByAttemptQuery, error) {
	var typedReq = new(adto.ListPaymentTransactionsByAttemptReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewListPaymentTransactionsByAttemptQuery")
	}

	return &ListPaymentTransactionsByAttemptQuery{typedReq}, nil
}
