package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchaseStatus/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type UpdatePurchaseStatusHandler interface {
	grpcTypes.GrpcHandler[*UpdatePurchaseStatusCommand, *adto.UpdatePurchaseStatusResDto]
}

type updatePurchaseStatusHandler struct {
	logger logger.Logger
	cache  drepo.PurchaseCacheRepo
	db     drepo.PurchaseDBRepo
}

func NewUpdatePurchaseStatusHandler(l logger.Logger, cache drepo.PurchaseCacheRepo, db drepo.PurchaseDBRepo) UpdatePurchaseStatusHandler {
	return &updatePurchaseStatusHandler{l, cache, db}
}

func (h *updatePurchaseStatusHandler) Handle(ctx context.Context, command *UpdatePurchaseStatusCommand) (*adto.UpdatePurchaseStatusResDto, error) {
	purchaseId, err := h.db.UpdatePurchase(ctx, command.ID, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to update purchase status: %v", err)
		return nil, err
	}

	// Invalidate cache after creating purchase
	err = h.cache.DeleteAllPurchases(ctx)
	if err != nil {
		// Todo: Handle cache deletion failure
		h.logger.Warnf("Failed to delete all purchases cache after updating purchase status: %v", err)
	}

	return adto.NewUpdatePurchaseStatusResDto(purchaseId, true, "Purchase status updated successfully"), nil
}
