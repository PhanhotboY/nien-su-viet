package prepo

import (
	"context"

	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

type subscriptionDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewSubscriptionDbRepo(logger logger.Logger, db dbcontracts.TxContextDb) drepo.SubscriptionDbRepo {
	return subscriptionDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r subscriptionDbRepo) CreateSubscription(ctx context.Context, subscription *entity.Subscription) (string, error) {
	r.logger.Info("Creating subscription", "subscription", subscription)

	if err := r.db.WithTxIfExists(ctx).DB().Create(subscription).Error; err != nil {
		r.logger.Error("Failed to create subscription", "error", err)
		return "", err
	}

	r.logger.Info("Subscription created successfully", "id", subscription.ID)
	return subscription.ID.String(), nil
}

func (r subscriptionDbRepo) UpdateSubscription(ctx context.Context, id string, updates map[string]any) (string, error) {
	r.logger.Infof("Updating subscription: %s", id)

	var existingSubscription entity.Subscription
	if err := r.db.WithTxIfExists(ctx).DB().First(&existingSubscription, "id = ?", id).Error; err != nil {
		r.logger.Errorf("Failed to find Subscription for update: %s, error: %v", id, err)
		return "", err
	}

	if err := r.db.WithTxIfExists(ctx).DB().Model(&existingSubscription).Updates(updates).Error; err != nil {
		r.logger.Errorf("Failed to update Subscription: %s, error: %v", id, err)
		return "", err
	}

	r.logger.Infof("Subscription updated successfully: %s", id)
	return id, nil
}

func (r subscriptionDbRepo) GetSubscriptionByID(ctx context.Context, id string) (*entity.Subscription, error) {
	r.logger.Info("Getting subscription", "id", id)

	var subscription = new(entity.Subscription)
	if err := r.db.WithTxIfExists(ctx).DB().First(subscription, "id = ?", id).Error; err != nil {
		r.logger.Error("Failed to get subscription", "id", id, "error", err)
		return nil, err
	}

	r.logger.Info("Subscription retrieved successfully", "id", id)
	return subscription, nil
}

func (r subscriptionDbRepo) GetSubscriptionsByUserID(ctx context.Context, userID string) ([]*entity.Subscription, error) {
	r.logger.Info("Getting subscriptions by user ID", "userID", userID)

	var subscriptions []*entity.Subscription
	if err := r.db.WithTxIfExists(ctx).DB().Where("user_id = ?", userID).Find(&subscriptions).Error; err != nil {
		r.logger.Error("Failed to get subscriptions by user ID", "userID", userID, "error", err)
		return nil, err
	}

	r.logger.Info("Subscriptions retrieved successfully", "userID", userID, "count", len(subscriptions))
	return subscriptions, nil
}

func (r subscriptionDbRepo) ListSubscriptions(ctx context.Context, filter *sdto.ListQueryRequest) ([]*entity.Subscription, error) {
	r.logger.Info("Listing subscriptions", "filter", filter)

	var subscriptions []*entity.Subscription
	query := r.db.WithTxIfExists(ctx).DB()

	// Implement basic filtering if needed. Assuming simple GORM chain for now.
	if err := query.Limit(int(filter.Limit)).Offset(int((filter.Page - 1) * (filter.Limit))).Find(&subscriptions).Error; err != nil {
		r.logger.Error("Failed to list subscriptions", "error", err)
		return nil, err
	}

	r.logger.Info("Subscriptions listed successfully", "count", len(subscriptions))
	return subscriptions, nil
}
