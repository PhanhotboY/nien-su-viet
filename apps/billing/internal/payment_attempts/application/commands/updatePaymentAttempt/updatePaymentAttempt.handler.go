package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttempt/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type UpdatePaymentAttemptHandler interface {
	grpcTypes.GrpcHandler[*UpdatePaymentAttemptCommand, adto.UpdatePaymentAttemptResDto]
}

type updatePaymentAttemptHandler struct {
	logger logger.Logger
	cache  drepo.PaymentAttemptCacheRepo
	db     drepo.PaymentAttemptDBRepo
}

func NewUpdatePaymentAttemptHandler(
	l logger.Logger,
	cache drepo.PaymentAttemptCacheRepo,
	db drepo.PaymentAttemptDBRepo,
) UpdatePaymentAttemptHandler {
	return &updatePaymentAttemptHandler{l, cache, db}
}

func (h *updatePaymentAttemptHandler) Handle(ctx context.Context, command *UpdatePaymentAttemptCommand) (adto.UpdatePaymentAttemptResDto, error) {
	// Update payment attempt in database
	paymentAttemptId, err := h.db.UpdatePaymentAttempt(ctx, command.PurchaseID, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to update payment attempt: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to update payment attempt", "NewUpdatePaymentAttemptHandler")
	}

	h.logger.Infof("Payment attempt updated with ID: %s", paymentAttemptId)

	// Invalidate cache of all payment attempts after updating new one
	err = h.cache.DeleteAllPaymentAttempts(ctx)
	if err != nil {
		// Log warning but don't fail the operation
		h.logger.Warnf("Failed to delete all payment attempts cache after updating payment attempt: %v", err)
	}

	// Build response DTO
	return adto.NewUpdatePaymentAttemptResDto(paymentAttemptId, true, "Payment attempt updated successfully"), nil
}
