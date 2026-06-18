package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
)

type OutboxEventDbRepo interface {
	CreateOutboxEvent(ctx context.Context, event *entity.OutboxEvent) error
	UpdateOutboxEvent(ctx context.Context, event *entity.OutboxEvent) error
	GetOutboxEventByID(ctx context.Context, id string) (*entity.OutboxEvent, error)
	GetOutboxEventsByStatus(ctx context.Context, status entity.OutboxEventStatus, limit int) ([]*entity.OutboxEvent, error)
	FindByEventType(ctx context.Context, eventType string) ([]*entity.OutboxEvent, error)

	FetchPending(ctx context.Context, limit int) ([]*entity.OutboxEvent, error)
	MarkPublished(ctx context.Context, id string) error
	MarkFailed(ctx context.Context, id string, err error) error
}
