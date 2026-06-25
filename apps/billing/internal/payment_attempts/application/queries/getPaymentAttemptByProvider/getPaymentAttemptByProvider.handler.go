package aquery

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttemptByProvider/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPaymentAttemptByProviderHandler interface {
	grpcTypes.GrpcHandler[*GetPaymentAttemptByProviderQuery, adto.GetPaymentAttemptByProviderResDto]
}

type getPaymentAttemptByProviderHandler struct {
	logger logger.Logger
	cache  drepo.PaymentAttemptCacheRepo
	db     drepo.PaymentAttemptDBRepo
}

func NewGetPaymentAttemptByProviderHandler(l logger.Logger, c drepo.PaymentAttemptCacheRepo, db drepo.PaymentAttemptDBRepo) GetPaymentAttemptByProviderHandler {
	return getPaymentAttemptByProviderHandler{l, c, db}
}

func (h getPaymentAttemptByProviderHandler) Handle(ctx context.Context, query *GetPaymentAttemptByProviderQuery) (adto.GetPaymentAttemptByProviderResDto, error) {
	// If cache miss, get payment attempt from database
	paymentAttempt, err := h.db.GetPaymentAttemptByProviderTransactionID(ctx, query.Provider, query.ProviderTransactionId)
	if err != nil {
		h.logger.Errorf("Failed to get payment attempt from database for provider transaction ID %s: %v", query.ProviderTransactionId, err)
		return nil, err
	}

	return adto.NewGetPaymentAttemptByProviderResDto(*paymentAttempt), nil
}
