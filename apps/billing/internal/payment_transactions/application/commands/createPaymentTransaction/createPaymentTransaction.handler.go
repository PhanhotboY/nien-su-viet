package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePaymentTransactionHandler interface {
	types.GrpcHandler[CreatePaymentTransactionCmd, adto.CreatePaymentTransactionResDto]
}

type createPaymentTransactionHandler struct {
	logger    logger.Logger
	dbRepo    drepo.PaymentTransactionDBRepo
	cacheRepo drepo.PaymentTransactionCacheRepo
}

func NewCreatePaymentTransactionHandler(l logger.Logger, dbrepo drepo.PaymentTransactionDBRepo, cacherepo drepo.PaymentTransactionCacheRepo) CreatePaymentTransactionHandler {
	return &createPaymentTransactionHandler{l, dbrepo, cacherepo}
}

func (h *createPaymentTransactionHandler) Handle(ctx context.Context, cmd CreatePaymentTransactionCmd) (adto.CreatePaymentTransactionResDto, error) {
	ptId, err := h.dbRepo.CreatePaymentTransaction(ctx, cmd.MapToEntity())
	if err != nil {
		h.logger.Error("Failed to create payment transaction", "error", err)
		return nil, err
	}

	return adto.NewCreatePaymentTransactionResDto(ptId, true, "Payment transaction created successfully"), nil
}
