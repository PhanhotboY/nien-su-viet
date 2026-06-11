package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePaymentAttemptHandler interface {
	grpcTypes.GrpcHandler[*CreatePaymentAttemptCommand, adto.CreatePaymentAttemptResDto]
}

type createPaymentAttemptHandler struct {
	logger logger.Logger
	cache  drepo.PaymentAttemptCacheRepo
	db     drepo.PaymentAttemptDBRepo
}

func NewCreatePaymentAttemptHandler(
	l logger.Logger,
	cache drepo.PaymentAttemptCacheRepo,
	db drepo.PaymentAttemptDBRepo,
) CreatePaymentAttemptHandler {
	return &createPaymentAttemptHandler{l, cache, db}
}

func (h *createPaymentAttemptHandler) Handle(ctx context.Context, command *CreatePaymentAttemptCommand) (adto.CreatePaymentAttemptResDto, error) {
	// Create payment attempt in database
	paymentAttemptId, err := h.db.CreatePaymentAttempt(ctx, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to create payment attempt: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to create payment attempt", "NewCreatePaymentAttemptHandler")
	}

	h.logger.Infof("Payment attempt created with ID: %s", paymentAttemptId)

	// Invalidate cache of all payment attempts after creating new one
	err = h.cache.DeleteAllPaymentAttempts(ctx)
	if err != nil {
		// Log warning but don't fail the operation
		h.logger.Warnf("Failed to delete all payment attempts cache after creating payment attempt: %v", err)
	}

	// Build response DTO
	return adto.NewCreatePaymentAttemptResDto(paymentAttemptId, true, "Payment attempt created successfully"), nil
}
