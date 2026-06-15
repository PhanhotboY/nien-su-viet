package prepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

type processedEventDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewProcessedEventDbRepo(logger logger.Logger, db dbcontracts.TxContextDb) drepo.ProcessedEventDbRepo {
	return processedEventDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r processedEventDbRepo) CreateProcessedEvent(ctx context.Context, event *entity.ProcessedEvent) error {
	r.logger.Info("Creating processed event", "event", event)

	if err := r.db.WithTxIfExists(ctx).DB().Create(event).Error; err != nil {
		r.logger.Error("Failed to create processed event", "error", err)
		return err
	}

	r.logger.Info("Processed event created successfully", "id", event.ID)
	return nil
}

func (r processedEventDbRepo) UpdateProcessedEvent(ctx context.Context, event *entity.ProcessedEvent) error {
	r.logger.Info("Updating processed event", "id", event.ID)

	if err := r.db.WithTxIfExists(ctx).DB().Save(event).Error; err != nil {
		r.logger.Error("Failed to update processed event", "id", event.ID, "error", err)
		return err
	}

	r.logger.Info("Processed event updated successfully", "id", event.ID)
	return nil
}

func (r processedEventDbRepo) GetProcessedEventByID(ctx context.Context, id string) (*entity.ProcessedEvent, error) {
	r.logger.Info("Getting processed event", "id", id)

	var event = new(entity.ProcessedEvent)
	if err := r.db.WithTxIfExists(ctx).DB().First(event, "id = ?", id).Error; err != nil {
		r.logger.Error("Failed to get processed event", "id", id, "error", err)
		return nil, err
	}

	r.logger.Info("Processed event retrieved successfully", "id", id)
	return event, nil
}
