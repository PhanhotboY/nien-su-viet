package prepo // persistence repository

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/repository"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type inBoxEventDbRepo struct {
	db dbcontracts.TxContextDb
}

func NewInBoxEventDbRepo(db dbcontracts.TxContextDb) drepo.InBoxEventDbRepo {
	return &inBoxEventDbRepo{db: db}
}

func (r *inBoxEventDbRepo) Insert(ctx context.Context, event *entity.InboxEvent) (string, error) {
	err := r.db.WithTxIfExists(ctx).DB().Create(event).Error
	if err != nil {
		return "", err
	}
	return event.ID.String(), nil
}

func (r *inBoxEventDbRepo) UpdateStatus(ctx context.Context, id string, status entity.InboxEventStatus) error {
	return r.db.WithTxIfExists(ctx).DB().Model(&entity.InboxEvent{}).Where("id = ?", id).Update("status", status).Error
}

func (r *inBoxEventDbRepo) FindByEventType(ctx context.Context, eventType string) ([]*entity.InboxEvent, error) {
	var event []*entity.InboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().Model(&entity.InboxEvent{}).Where("event_type = ?", eventType).Find(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *inBoxEventDbRepo) FindByID(ctx context.Context, id string) (*entity.InboxEvent, error) {
	var event entity.InboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().First(&event, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *inBoxEventDbRepo) FetchPending(ctx context.Context, limit int) ([]*entity.InboxEvent, error) {
	var events []*entity.InboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().Raw(`
		SELECT *
    FROM inbox_events
    WHERE status = ? OR (status = ? AND next_retry_at <= now())
    ORDER BY created_at
    LIMIT ?
    FOR UPDATE SKIP LOCKED
	`, entity.INBOX_EVENT_STATUS_PENDING, entity.INBOX_EVENT_STATUS_FAILED, limit).Find(&events).Error; err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, nil
	}
	return events, nil
}

func (r *inBoxEventDbRepo) MarkPublished(
	ctx context.Context,
	id string,
) error {
	return r.db.WithTxIfExists(ctx).DB().
		Model(&entity.InboxEvent{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"status":       entity.INBOX_EVENT_STATUS_PUBLISHED,
			"published_at": time.Now(),
		}).Error
}

const MaxRetryCount = 10

func (r *inBoxEventDbRepo) MarkFailed(
	ctx context.Context,
	id string,
	publishErr error,
) error {
	return r.db.WithTxIfExists(ctx).DB().
		Transaction(func(tx *gorm.DB) error {

			var event entity.InboxEvent

			if err := tx.
				Clauses(clause.Locking{
					Strength: "UPDATE",
				}).
				First(&event, "id = ?", id).Error; err != nil {
				return err
			}

			retryCount := event.RetryCount + 1

			status := entity.INBOX_EVENT_STATUS_FAILED

			var nextRetryAt *time.Time

			if retryCount >= MaxRetryCount {
				status = entity.INBOX_EVENT_STATUS_DEAD
			} else {
				t := calculateNextRetry(retryCount)
				nextRetryAt = &t
			}

			updates := map[string]any{
				"status":        status,
				"retry_count":   retryCount,
				"last_error":    publishErr.Error(),
				"next_retry_at": nextRetryAt,
			}

			return tx.Model(&entity.InboxEvent{}).
				Where("id = ?", id).
				Updates(updates).Error
		})

}

// retry backoff
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
