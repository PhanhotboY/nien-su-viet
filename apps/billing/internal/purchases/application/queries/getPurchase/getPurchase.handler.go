package aquery

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPurchaseHandler interface {
	grpcTypes.GrpcHandler[*GetPurchaseQuery, adto.GetPurchaseResDto]
}

type getPurchaseHandler struct {
	logger logger.Logger
	cache  drepo.PurchaseCacheRepo
	db     drepo.PurchaseDBRepo
}

func NewGetPurchaseHandler(l logger.Logger, c drepo.PurchaseCacheRepo, db drepo.PurchaseDBRepo) GetPurchaseHandler {
	return getPurchaseHandler{l, c, db}
}

func (h getPurchaseHandler) Handle(ctx context.Context, query *GetPurchaseQuery) (adto.GetPurchaseResDto, error) {
	// Try to get purchase from cache first
	purchase, err := h.cache.GetPurchase(ctx, query.PurchaseId)
	if err == nil {
		h.logger.Infof("Cache hit for purchase ID %s", query.PurchaseId)
		return adto.NewGetPurchaseResDto(*purchase), nil
	}
	h.logger.Infof("Cache miss for purchase ID %s: %v", query.PurchaseId, err)

	// If cache miss, get purchase from database
	purchase, err = h.db.GetPurchaseById(ctx, query.PurchaseId)
	if err != nil {
		h.logger.Errorf("Failed to get purchase from database for ID %s: %v", query.PurchaseId, err)
		return nil, err
	}

	// Cache the purchase for future requests
	err = h.cache.PutPurchase(ctx, query.PurchaseId, purchase)
	if err != nil {
		h.logger.Warnf("Failed to cache purchase ID %s: %v", query.PurchaseId, err)
	}

	return adto.NewGetPurchaseResDto(*purchase), nil
}
