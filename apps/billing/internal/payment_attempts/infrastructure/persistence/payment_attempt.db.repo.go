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
	r.logger.Info("Creating payment attempt", paymentAttempt)

	if err := r.db.WithTxIfExists(ctx).DB().Create(paymentAttempt).Error; err != nil {
		r.logger.Error("Failed to create payment attempt", err)
		return "", err
	}

	r.logger.Info("Payment attempt created successfully")
	return paymentAttempt.ID.String(), nil
}

func (r paymentAttemptDbRepo) UpdatePaymentAttempt(ctx context.Context, paymentAttemptId string, updates map[string]any) (string, error) {
	r.logger.Info("Updating payment attempt", paymentAttemptId, updates)

	var existingPaymentAttempt entity.PaymentAttempt
	if err := r.db.WithTxIfExists(ctx).DB().First(&existingPaymentAttempt, "id = ?", paymentAttemptId).Error; err != nil {
		r.logger.Error("Failed to find payment attempt for update", paymentAttemptId, err)
		return "", err
	}

	if err := r.db.WithTxIfExists(ctx).DB().Model(&existingPaymentAttempt).Updates(updates).Error; err != nil {
		r.logger.Errorf("Failed to update payment attempt %s, error: %v", paymentAttemptId, err)
		return "", err
	}

	r.logger.Info("Payment attempt updated successfully %s", paymentAttemptId)
	return paymentAttemptId, nil
}

func (r paymentAttemptDbRepo) DeletePaymentAttempt(ctx context.Context, paymentAttemptId string) (string, error) {
	r.logger.Info("Deleting payment attempt ", paymentAttemptId)

	if err := r.db.WithTxIfExists(ctx).DB().Delete(&entity.PaymentAttempt{}, "id = ?", paymentAttemptId).Error; err != nil {
		r.logger.Error("Failed to delete payment attempt ", paymentAttemptId, err)
		return "", err
	}

	r.logger.Info("Payment attempt deleted successfully ", paymentAttemptId)
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

func (r paymentAttemptDbRepo) GetPaymentAttemptByProviderTransactionID(ctx context.Context, provider string, providerTransactionID string) (*entity.PaymentAttempt, error) {
	r.logger.Infof("Getting payment attempt by provider transaction ID: %s", providerTransactionID)

	var paymentAttempt entity.PaymentAttempt
	if err := r.db.WithTxIfExists(ctx).DB().First(&paymentAttempt, "provider_transaction_id = ? AND provider = ?", providerTransactionID, provider).Error; err != nil {
		r.logger.Errorf("Failed to get payment attempt by provider transaction ID: %s, error: %v", providerTransactionID, err)
		if err == gorm.ErrRecordNotFound {
			r.logger.Infof("No payment attempt found for provider transaction ID: %s", providerTransactionID)
			return nil, nil
		}
		return nil, err
	}

	r.logger.Infof("Payment attempt retrieved successfully by provider transaction ID: %s", providerTransactionID)
	return &paymentAttempt, nil
}
