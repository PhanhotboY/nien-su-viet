package aqueries

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/queries/listPaymentTransactionsByAttempt/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type ListPaymentTransactionsByAttemptHandler interface {
	grpcTypes.GrpcHandler[*ListPaymentTransactionsByAttemptQuery, adto.ListPaymentTransactionsByAttemptResDto]
}

type listPaymentTransactionsByAttemptHandler struct {
	logger logger.Logger
	db     drepo.PaymentTransactionDBRepo
}

func NewListPaymentTransactionsByAttemptHandler(l logger.Logger, db drepo.PaymentTransactionDBRepo) ListPaymentTransactionsByAttemptHandler {
	return listPaymentTransactionsByAttemptHandler{l, db}
}

func (h listPaymentTransactionsByAttemptHandler) Handle(ctx context.Context, query *ListPaymentTransactionsByAttemptQuery) (adto.ListPaymentTransactionsByAttemptResDto, error) {
	h.logger.Infof("Getting payment transactions for attempt ID %s", query.PaymentAttemptId)

	transactions, err := h.db.GetPaymentTransactionsByPaymentAttemptId(ctx, query.PaymentAttemptId)
	if err != nil {
		h.logger.Errorf("Failed to get payment transactions for attempt ID %s: %v", query.PaymentAttemptId, err)
		return nil, err
	}

	return adto.NewListPaymentTransactionsByAttemptResDto(transactions), nil
}
