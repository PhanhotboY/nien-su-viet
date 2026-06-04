package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePurchaseHandler interface {
	grpcTypes.GrpcHandler[*CreatePurchaseCommand, *adto.CreatePurchaseResDto]
}

type createPurchaseHandler struct {
	logger logger.Logger
	cache  drepo.PurchaseCacheRepo
	db     drepo.PurchaseDBRepo
}

func NewCreatePurchaseHandler(l logger.Logger, cache drepo.PurchaseCacheRepo, db drepo.PurchaseDBRepo) CreatePurchaseHandler {
	return &createPurchaseHandler{l, cache, db}
}

func (h *createPurchaseHandler) Handle(ctx context.Context, command *CreatePurchaseCommand) (*adto.CreatePurchaseResDto, error) {
	purchaseId, err := h.db.CreatePurchase(ctx, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to create purchase: %v", err)
		return nil, err
	}

	// Invalidate cache after creating purchase
	err = h.cache.DeleteAllPurchases(ctx)
	if err != nil {
		// Todo: Handle cache deletion failure
		h.logger.Warnf("Failed to delete all purchases cache after creating purchase: %v", err)
	}

	return adto.NewCreatePurchaseResDto(purchaseId, true, "Purchase created successfully"), nil
}
