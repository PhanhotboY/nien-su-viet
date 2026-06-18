package prepo

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type outboxEventDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewOutboxEventDbRepo(logger logger.Logger, db dbcontracts.TxContextDb) drepo.OutboxEventDbRepo {
	return outboxEventDbRepo{
		logger: logger,
		db:     db,
	}
}

func (r outboxEventDbRepo) CreateOutboxEvent(ctx context.Context, event *entity.OutboxEvent) error {
	r.logger.Info("Creating outbox event", "event", event)

	if err := r.db.WithTxIfExists(ctx).DB().Create(event).Error; err != nil {
		r.logger.Error("Failed to create outbox event", "error", err)
		return err
	}

	r.logger.Info("Outbox event created successfully", "id", event.ID)
	return nil
}

func (r outboxEventDbRepo) UpdateOutboxEvent(ctx context.Context, event *entity.OutboxEvent) error {
	r.logger.Info("Updating outbox event", "id", event.ID)

	if err := r.db.WithTxIfExists(ctx).DB().Save(event).Error; err != nil {
		r.logger.Error("Failed to update outbox event", "id", event.ID, "error", err)
		return err
	}

	r.logger.Info("Outbox event updated successfully", "id", event.ID)
	return nil
}

func (r outboxEventDbRepo) GetOutboxEventByID(ctx context.Context, id string) (*entity.OutboxEvent, error) {
	r.logger.Info("Getting outbox event", "id", id)

	var event = new(entity.OutboxEvent)
	if err := r.db.WithTxIfExists(ctx).DB().First(event, "id = ?", id).Error; err != nil {
		r.logger.Error("Failed to get outbox event", "id", id, "error", err)
		return nil, err
	}

	r.logger.Info("Outbox event retrieved successfully", "id", id)
	return event, nil
}

func (r outboxEventDbRepo) GetOutboxEventsByStatus(ctx context.Context, status entity.OutboxEventStatus, limit int) ([]*entity.OutboxEvent, error) {
	r.logger.Info("Getting outbox events by status", "status", status, "limit", limit)

	var events []*entity.OutboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().Where("status = ?", status).Limit(limit).Find(&events).Error; err != nil {
		r.logger.Error("Failed to get outbox events by status", "status", status, "error", err)
		return nil, err
	}

	r.logger.Info("Outbox events retrieved successfully", "status", status, "count", len(events))
	return events, nil
}

func (r outboxEventDbRepo) FindByEventType(ctx context.Context, eventType string) ([]*entity.OutboxEvent, error) {
	var event []*entity.OutboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().Model(&entity.OutboxEvent{}).Where("event_type = ?", eventType).Find(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r outboxEventDbRepo) FetchPending(ctx context.Context, limit int) ([]*entity.OutboxEvent, error) {
	r.logger.Info("Fetching pending outbox events", "limit", limit)

	var events []*entity.OutboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().Raw(`
		SELECT *
		FROM outbox_events
		WHERE status = ? OR (status = ? AND next_retry_at <= now())
		ORDER BY created_at
		LIMIT ?
		FOR UPDATE SKIP LOCKED
	`, entity.OUTBOX_EVENT_STATUS_PENDING, entity.OUTBOX_EVENT_STATUS_FAILED, limit).Scan(&events).Error; err != nil {
		r.logger.Error("Failed to fetch pending outbox events", "error", err)
		return nil, err
	}

	return events, nil
}

func (r outboxEventDbRepo) MarkPublished(ctx context.Context, id string) error {
	r.logger.Info("Marking outbox event as published", "id", id)

	return r.db.WithTxIfExists(ctx).DB().
		Model(&entity.OutboxEvent{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":       entity.OUTBOX_EVENT_STATUS_PUBLISHED,
			"published_at": time.Now(),
		}).Error
}

const MaxRetryCount = 10

func (r outboxEventDbRepo) MarkFailed(ctx context.Context, id string, publishErr error) error {
	r.logger.Info("Marking outbox event as failed", "id", id, "error", publishErr)

	return r.db.WithTxIfExists(ctx).DB().Transaction(func(tx *gorm.DB) error {
		var event entity.OutboxEvent

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&event, "id = ?", id).Error; err != nil {
			return err
		}

		retryCount := event.RetryCount + 1
		status := entity.OUTBOX_EVENT_STATUS_FAILED
		var nextRetryAt *time.Time

		if retryCount >= MaxRetryCount {
			status = entity.OUTBOX_EVENT_STATUS_DEAD
		} else {
			t := calculateNextRetry(retryCount)
			nextRetryAt = &t
		}

		updates := map[string]any{
			"status":        status,
			"retry_count":   retryCount,
			"next_retry_at": nextRetryAt,
		}

		return tx.Model(&entity.OutboxEvent{}).Where("id = ?", id).Updates(updates).Error
	})
}

func calculateNextRetry(retryCount int32) time.Time {
	now := time.Now()
	switch retryCount {
	case 1:
		return now.Add(5 * time.Second)
	case 2:
		return now.Add(30 * time.Second)
	case 3:
		return now.Add(2 * time.Minute)
	case 4:
		return now.Add(10 * time.Minute)
	default:
		return now.Add(1 * time.Hour)
	}
}
