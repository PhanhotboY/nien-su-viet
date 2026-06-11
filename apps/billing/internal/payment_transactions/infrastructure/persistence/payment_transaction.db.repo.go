package prepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

type paymentTransactionDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewPaymentTransactionDbRepo(logger logger.Logger, db dbcontracts.TxContextDb) drepo.PaymentTransactionDBRepo {
	return paymentTransactionDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r paymentTransactionDbRepo) CreatePaymentTransaction(ctx context.Context, transaction *entity.PaymentTransaction) (string, error) {
	r.logger.Info("Creating payment transaction", "transaction", transaction)

	if err := r.db.WithTxIfExists(ctx).DB().Create(transaction).Error; err != nil {
		r.logger.Error("Failed to create payment transaction", "error", err)
		return "", err
	}

	r.logger.Info("Payment transaction created successfully")
	return transaction.ID.String(), nil
}

func (r paymentTransactionDbRepo) UpdatePaymentTransaction(ctx context.Context, transactionId string, updates map[string]any) (string, error) {
	r.logger.Info("Updating payment transaction", "transactionId", transactionId, "updates", updates)

	var existingTransaction entity.PaymentTransaction
	if err := r.db.WithTxIfExists(ctx).DB().First(&existingTransaction, "id = ?", transactionId).Error; err != nil {
		r.logger.Error("Failed to find payment transaction for update", "transactionId", transactionId, "error", err)
		return "", err
	}

	if err := r.db.WithTxIfExists(ctx).DB().Model(&existingTransaction).Updates(updates).Error; err != nil {
		r.logger.Error("Failed to update payment transaction", "transactionId", transactionId, "error", err)
		return "", err
	}

	r.logger.Info("Payment transaction updated successfully", "transactionId", transactionId)
	return transactionId, nil
}

func (r paymentTransactionDbRepo) DeletePaymentTransaction(ctx context.Context, transactionId string) (string, error) {
	r.logger.Info("Deleting payment transaction", "transactionId", transactionId)

	if err := r.db.WithTxIfExists(ctx).DB().Delete(&entity.PaymentTransaction{}, "id = ?", transactionId).Error; err != nil {
		r.logger.Error("Failed to delete payment transaction", "transactionId", transactionId, "error", err)
		return "", err
	}

	r.logger.Info("Payment transaction deleted successfully", "transactionId", transactionId)
	return transactionId, nil
}

func (r paymentTransactionDbRepo) GetPaymentTransactionById(ctx context.Context, transactionId string) (*entity.PaymentTransaction, error) {
	r.logger.Info("Getting payment transaction", "transactionId", transactionId)

	var transaction = new(entity.PaymentTransaction)
	if err := r.db.WithTxIfExists(ctx).DB().First(transaction, "id = ?", transactionId).Error; err != nil {
		r.logger.Error("Failed to get payment transaction", "transactionId", transactionId, "error", err)
		return nil, err
	}

	r.logger.Info("Payment transaction retrieved successfully", "transactionId", transactionId)
	return transaction, nil
}

func (r paymentTransactionDbRepo) GetPaymentTransactionsByPaymentAttemptId(ctx context.Context, paymentAttemptId string) ([]*entity.PaymentTransaction, error) {
	r.logger.Info("Getting payment transactions by payment attempt id", "paymentAttemptId", paymentAttemptId)

	var transactions []*entity.PaymentTransaction
	if err := r.db.WithTxIfExists(ctx).DB().Where("payment_attempt_id = ?", paymentAttemptId).Find(&transactions).Error; err != nil {
		r.logger.Error("Failed to get payment transactions by payment attempt id", "paymentAttemptId", paymentAttemptId, "error", err)
		return nil, err
	}

	r.logger.Info("Payment transactions retrieved successfully by payment attempt id", "paymentAttemptId", paymentAttemptId, "count", len(transactions))
	return transactions, nil
}
