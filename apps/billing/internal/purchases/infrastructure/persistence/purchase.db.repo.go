package prepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"gorm.io/gorm"
)

type purchaseDbRepo struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewPurchaseDbRepo(logger logger.Logger, db *gorm.DB) drepo.PurchaseDBRepo {
	return purchaseDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r purchaseDbRepo) CreatePurchase(ctx context.Context, purchase *entity.Purchase) (string, error) {
	r.logger.Info("Creating purchase", "purchase", purchase)

	if err := r.db.WithContext(ctx).Create(purchase).Error; err != nil {
		r.logger.Error("Failed to create purchase", "error", err)
		return "", err
	}

	r.logger.Info("Purchase created successfully", "purchaseId", purchase.ID)
	return purchase.ID.String(), nil
}

func (r purchaseDbRepo) UpdatePurchase(ctx context.Context, purchaseId string, updates map[string]any) (string, error) {
	r.logger.Info("Updating purchase", "purchaseId", purchaseId, "updates", updates)

	var existingPurchase entity.Purchase
	if err := r.db.WithContext(ctx).First(&existingPurchase, "id = ?", purchaseId).Error; err != nil {
		r.logger.Error("Failed to find purchase for update", "purchaseId", purchaseId, "error", err)
		return "", err
	}

	if err := r.db.WithContext(ctx).Model(&existingPurchase).Updates(updates).Error; err != nil {
		r.logger.Error("Failed to update purchase", "purchaseId", purchaseId, "error", err)
		return "", err
	}

	r.logger.Info("Purchase updated successfully", "purchaseId", purchaseId)
	return purchaseId, nil
}

func (r purchaseDbRepo) DeletePurchase(ctx context.Context, purchaseId string) (string, error) {
	r.logger.Info("Deleting purchase", "purchaseId", purchaseId)

	if err := r.db.WithContext(ctx).Delete(&entity.Purchase{}, "id = ?", purchaseId).Error; err != nil {
		r.logger.Error("Failed to delete purchase", "purchaseId", purchaseId, "error", err)
		return "", err
	}

	r.logger.Info("Purchase deleted successfully", "purchaseId", purchaseId)
	return purchaseId, nil
}

func (r purchaseDbRepo) GetPurchase(ctx context.Context, purchaseId string) (*entity.Purchase, error) {
	r.logger.Info("Getting purchase", "purchaseId", purchaseId)

	var purchase = new(entity.Purchase)
	if err := r.db.WithContext(ctx).First(purchase, "id = ?", purchaseId).Error; err != nil {
		r.logger.Error("Failed to get purchase", "purchaseId", purchaseId, "error", err)
		return nil, err
	}

	r.logger.Info("Purchase retrieved successfully", "purchaseId", purchaseId)
	return purchase, nil
}
