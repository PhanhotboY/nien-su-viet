package acmd

import (
	"context"

	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePaymentTransactionHandler interface {
	types.GrpcHandler[CreatePaymentTransactionCmd, adto.CreatePaymentTransactionResDto]
}

type createPaymentTransactionHandler struct {
	logger    logger.Logger
	dbRepo    drepo.PaymentAttemptDBRepo
	cacheRepo drepo.PaymentAttemptCacheRepo
}

func NewCreatePaymentTransactionHandler(l logger.Logger, dbrepo drepo.PaymentAttemptDBRepo, cacherepo drepo.PaymentAttemptCacheRepo) CreatePaymentTransactionHandler {
	return &createPaymentTransactionHandler{l, dbrepo, cacherepo}
}

func (h *createPaymentTransactionHandler) Handle(ctx context.Context, cmd CreatePaymentTransactionCmd) (adto.CreatePaymentTransactionResDto, error) {
	ptId, err := h.dbRepo.CreatePaymentAttempt(ctx, cmd.MapToEntity())
	if err != nil {
		h.logger.Error("Failed to create payment attempt", "error", err)
		return nil, err
	}

	return adto.NewCreatePaymentTransactionResDto(ptId, true, "Payment transaction created successfully"), nil
}
