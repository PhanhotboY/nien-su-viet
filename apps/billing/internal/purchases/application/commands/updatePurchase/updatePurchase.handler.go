package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type UpdatePurchaseHandler interface {
	grpcTypes.GrpcHandler[*UpdatePurchaseCommand, *adto.UpdatePurchaseResDto]
}

type updatePurchaseHandler struct {
	logger logger.Logger
	cache  drepo.PurchaseCacheRepo
	db     drepo.PurchaseDBRepo
}

func NewUpdatePurchaseHandler(l logger.Logger, cache drepo.PurchaseCacheRepo, db drepo.PurchaseDBRepo) UpdatePurchaseHandler {
	return &updatePurchaseHandler{l, cache, db}
}

func (h *updatePurchaseHandler) Handle(ctx context.Context, command *UpdatePurchaseCommand) (*adto.UpdatePurchaseResDto, error) {
	purchaseId, err := h.db.UpdatePurchase(ctx, command.ID, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to update purchase: %v", err)
		return nil, err
	}

	// Invalidate cache after creating purchase
	err = h.cache.DeleteAllPurchases(ctx)
	if err != nil {
		// Todo: Handle cache deletion failure
		h.logger.Warnf("Failed to delete all purchases cache after updating purchase: %v", err)
	}

	return adto.NewUpdatePurchaseResDto(purchaseId, true, "Purchase updated successfully"), nil
}
