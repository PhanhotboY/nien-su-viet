package prepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"gorm.io/gorm"
)

type paymentAttemptDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewPaymentAttemptDbRepo(logger logger.Logger, db dbcontracts.TxContextDb) repository.PaymentAttemptDBRepo {
	return paymentAttemptDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r paymentAttemptDbRepo) CreatePaymentAttempt(ctx context.Context, paymentAttempt *entity.PaymentAttempt) (string, error) {
	r.logger.Info("Creating payment attempt", "paymentAttempt", paymentAttempt)

	if err := r.db.WithTxIfExists(ctx).DB().Create(paymentAttempt).Error; err != nil {
		r.logger.Error("Failed to create payment attempt", "error", err)
		return "", err
	}

	r.logger.Info("Payment attempt created successfully")
	return paymentAttempt.ID.String(), nil
}

func (r paymentAttemptDbRepo) UpdatePaymentAttempt(ctx context.Context, paymentAttemptId string, updates map[string]any) (string, error) {
	r.logger.Info("Updating payment attempt", "paymentAttemptId", paymentAttemptId, "updates", updates)

	var existingPaymentAttempt entity.PaymentAttempt
	if err := r.db.WithTxIfExists(ctx).DB().First(&existingPaymentAttempt, "id = ?", paymentAttemptId).Error; err != nil {
		r.logger.Error("Failed to find payment attempt for update", "paymentAttemptId", paymentAttemptId, "error", err)
		return "", err
	}

	if err := r.db.WithTxIfExists(ctx).DB().Model(&existingPaymentAttempt).Updates(updates).Error; err != nil {
		r.logger.Error("Failed to update payment attempt", "paymentAttemptId", paymentAttemptId, "error", err)
		return "", err
	}

	r.logger.Info("Payment attempt updated successfully", "paymentAttemptId", paymentAttemptId)
	return paymentAttemptId, nil
}

func (r paymentAttemptDbRepo) DeletePaymentAttempt(ctx context.Context, paymentAttemptId string) (string, error) {
	r.logger.Info("Deleting payment attempt", "paymentAttemptId", paymentAttemptId)

	if err := r.db.WithTxIfExists(ctx).DB().Delete(&entity.PaymentAttempt{}, "id = ?", paymentAttemptId).Error; err != nil {
		r.logger.Error("Failed to delete payment attempt", "paymentAttemptId", paymentAttemptId, "error", err)
		return "", err
	}

	r.logger.Info("Payment attempt deleted successfully", "paymentAttemptId", paymentAttemptId)
	return paymentAttemptId, nil
}

func (r paymentAttemptDbRepo) GetPaymentAttemptById(ctx context.Context, paymentAttemptId string) (*entity.PaymentAttempt, error) {
	r.logger.Info("Getting payment attempt", "paymentAttemptId", paymentAttemptId)

	var paymentAttempt = new(entity.PaymentAttempt)
	if err := r.db.WithTxIfExists(ctx).DB().First(paymentAttempt, "id = ?", paymentAttemptId).Error; err != nil {
		r.logger.Error("Failed to get payment attempt", "paymentAttemptId", paymentAttemptId, "error", err)
		return nil, err
	}

	r.logger.Info("Payment attempt retrieved successfully", "paymentAttemptId", paymentAttemptId)
	return paymentAttempt, nil
}

func (r paymentAttemptDbRepo) GetPaymentAttemptByProviderTransactionID(ctx context.Context, providerTransactionID string) (*entity.PaymentAttempt, error) {
	r.logger.Info("Getting payment attempt by provider transaction ID", "providerTransactionID", providerTransactionID)

	var paymentAttempt = new(entity.PaymentAttempt)
	if err := r.db.WithTxIfExists(ctx).DB().First(paymentAttempt, "provider_transaction_id = ?", providerTransactionID).Error; err != nil {
		r.logger.Error("Failed to get payment attempt by provider transaction ID", "providerTransactionID", providerTransactionID, "error", err)
		if err == gorm.ErrRecordNotFound {
			r.logger.Info("No payment attempt found for provider transaction ID", "providerTransactionID", providerTransactionID)
			return nil, nil
		}
		return nil, err
	}

	r.logger.Info("Payment attempt retrieved successfully by provider transaction ID", "providerTransactionID", providerTransactionID)
	return paymentAttempt, nil
}
