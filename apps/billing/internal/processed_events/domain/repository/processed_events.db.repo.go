package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/entity"
)

type ProcessedEventDbRepo interface {
	CreateProcessedEvent(ctx context.Context, event *entity.ProcessedEvent) error
	UpdateProcessedEvent(ctx context.Context, event *entity.ProcessedEvent) error
	GetProcessedEventByID(ctx context.Context, id string) (*entity.ProcessedEvent, error)
}
