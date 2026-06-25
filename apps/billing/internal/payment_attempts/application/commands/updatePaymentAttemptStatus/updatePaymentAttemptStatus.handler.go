package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type UpdatePaymentAttemptStatusHandler interface {
	grpcTypes.GrpcHandler[*UpdatePaymentAttemptStatusCommand, adto.UpdatePaymentAttemptStatusResDto]
}

type updatePaymentAttemptStatusHandler struct {
	logger logger.Logger
	cache  drepo.PaymentAttemptCacheRepo
	db     drepo.PaymentAttemptDBRepo
}

func NewUpdatePaymentAttemptStatusHandler(
	l logger.Logger,
	cache drepo.PaymentAttemptCacheRepo,
	db drepo.PaymentAttemptDBRepo,
) UpdatePaymentAttemptStatusHandler {
	return &updatePaymentAttemptStatusHandler{l, cache, db}
}

func (h *updatePaymentAttemptStatusHandler) Handle(ctx context.Context, command *UpdatePaymentAttemptStatusCommand) (adto.UpdatePaymentAttemptStatusResDto, error) {
	// Update payment AttemptStatus in database
	paymentAttemptId, err := h.db.UpdatePaymentAttempt(ctx, command.ID, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to update payment AttemptStatus: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to update payment AttemptStatus", "NewUpdatePaymentAttemptStatusHandler")
	}

	h.logger.Infof("Payment AttemptStatus updated with ID: %s", paymentAttemptId)

	// Invalidate cache of all payment AttemptStatuss after updating new one
	err = h.cache.DeleteAllPaymentAttempts(ctx)
	if err != nil {
		// Log warning but don't fail the operation
		h.logger.Warnf("Failed to delete all payment attempts cache after updating payment attempt: %v", err)
	}

	// Build response DTO
	return adto.NewUpdatePaymentAttemptStatusResDto(paymentAttemptId, true, "Payment attempt updated successfully"), nil
}
