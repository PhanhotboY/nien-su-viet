package aquery

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPaymentAttemptHandler interface {
	grpcTypes.GrpcHandler[*GetPaymentAttemptQuery, adto.GetPaymentAttemptResDto]
}

type getPaymentAttemptHandler struct {
	logger logger.Logger
	cache  drepo.PaymentAttemptCacheRepo
	db     drepo.PaymentAttemptDBRepo
}

func NewGetPaymentAttemptHandler(l logger.Logger, c drepo.PaymentAttemptCacheRepo, db drepo.PaymentAttemptDBRepo) GetPaymentAttemptHandler {
	return getPaymentAttemptHandler{l, c, db}
}

func (h getPaymentAttemptHandler) Handle(ctx context.Context, query *GetPaymentAttemptQuery) (adto.GetPaymentAttemptResDto, error) {
	// Try to get payment attempt from cache first
	paymentAttempt, err := h.cache.GetPaymentAttempt(ctx, query.PaymentAttemptId)
	if err != nil {
		h.logger.Warnf("failed to get payment attempt from cache for ID %s: %v", query.PaymentAttemptId, err)
	} else if paymentAttempt != nil {
		h.logger.Infof("Cache hit for payment attempt ID %s", query.PaymentAttemptId)
		return adto.NewGetPaymentAttemptResDto(*paymentAttempt), nil
	}
	h.logger.Infof("Cache miss for payment attempt ID %s: %v", query.PaymentAttemptId, err)

	// If cache miss, get payment attempt from database
	paymentAttempt, err = h.db.GetPaymentAttemptById(ctx, query.PaymentAttemptId)
	if err != nil {
		h.logger.Errorf("Failed to get payment attempt from database for ID %s: %v", query.PaymentAttemptId, err)
		return nil, err
	}

	// Cache the payment attempt for future requests
	err = h.cache.PutPaymentAttempt(ctx, query.PaymentAttemptId, paymentAttempt)
	if err != nil {
		h.logger.Warnf("Failed to cache payment attempt ID %s: %v", query.PaymentAttemptId, err)
	}

	return adto.NewGetPaymentAttemptResDto(*paymentAttempt), nil
}
