package prepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"gorm.io/gorm"
)

type purchaseDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewPurchaseDbRepo(logger logger.Logger, db dbcontracts.TxContextDb) drepo.PurchaseDBRepo {
	return purchaseDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r purchaseDbRepo) CreatePurchase(ctx context.Context, purchase *entity.Purchase) (string, error) {
	r.logger.Info("Creating purchase", "purchase", purchase)

	if err := r.db.WithTxIfExists(ctx).DB().Create(purchase).Error; err != nil {
		r.logger.Error("Failed to create purchase", "error", err)
		return "", err
	}

	r.logger.Info("Purchase created successfully")
	return purchase.ID.String(), nil
}

func (r purchaseDbRepo) UpdatePurchase(ctx context.Context, purchaseId string, updates map[string]any) (string, error) {
	r.logger.Infof("Updating purchase: %s", purchaseId)

	var existingPurchase entity.Purchase
	if err := r.db.WithTxIfExists(ctx).DB().First(&existingPurchase, "id = ?", purchaseId).Error; err != nil {
		r.logger.Errorf("Failed to find purchase for update: %s, error: %v", purchaseId, err)
		return "", err
	}

	if err := r.db.WithTxIfExists(ctx).DB().Model(&existingPurchase).Updates(updates).Error; err != nil {
		r.logger.Errorf("Failed to update purchase: %s, error: %v", purchaseId, err)
		return "", err
	}

	r.logger.Infof("Purchase updated successfully: %s", purchaseId)
	return purchaseId, nil
}

func (r purchaseDbRepo) DeletePurchase(ctx context.Context, purchaseId string) (string, error) {
	r.logger.Info("Deleting purchase", "purchaseId", purchaseId)

	if err := r.db.WithTxIfExists(ctx).DB().Delete(&entity.Purchase{}, "id = ?", purchaseId).Error; err != nil {
		r.logger.Error("Failed to delete purchase", "purchaseId", purchaseId, "error", err)
		return "", err
	}

	r.logger.Info("Purchase deleted successfully", "purchaseId", purchaseId)
	return purchaseId, nil
}

func (r purchaseDbRepo) GetPurchaseById(ctx context.Context, purchaseId string) (*entity.Purchase, error) {
	r.logger.Info("Getting purchase", "purchaseId", purchaseId)

	var purchase = new(entity.Purchase)
	if err := r.db.WithTxIfExists(ctx).DB().First(purchase, "id = ?", purchaseId).Error; err != nil {
		r.logger.Error("Failed to get purchase", "purchaseId", purchaseId, "error", err)
		return nil, err
	}

	r.logger.Info("Purchase retrieved successfully", "purchaseId", purchaseId)
	return purchase, nil
}

func (r purchaseDbRepo) GetPurchaseByIdempotencyKey(ctx context.Context, idempotencyKey string) (*entity.Purchase, error) {
	r.logger.Info("Getting purchase by idempotency key", "idempotencyKey", idempotencyKey)

	var purchase = new(entity.Purchase)
	if err := r.db.WithTxIfExists(ctx).DB().First(purchase, "idempotency_key = ?", idempotencyKey).Error; err != nil {
		r.logger.Error("Failed to get purchase by idempotency key", "idempotencyKey", idempotencyKey, "error", err)
		if err == gorm.ErrRecordNotFound {
			r.logger.Info("No purchase found for idempotency key", "idempotencyKey", idempotencyKey)
			return nil, nil
		}
		return nil, err
	}

	r.logger.Info("Purchase retrieved successfully by idempotency key", "idempotencyKey", idempotencyKey)
	return purchase, nil
}
