package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type HistoricalEventCacheRepository interface {
	PutHistoricalEvent(ctx context.Context, key string, event *entity.HistoricalEvent) error
	PutHistoricalEvents(ctx context.Context, key string, events *utils.PaginatedResponse[entity.HistoricalEventBrief]) error
	GetHistoricalEvent(ctx context.Context, key string) (*entity.HistoricalEvent, error)
	GetHistoricalEvents(ctx context.Context, key string) (*utils.PaginatedResponse[entity.HistoricalEventBrief], error)
	DeleteHistoricalEvent(ctx context.Context, key string) error
	DeleteAllHistoricalEvents(ctx context.Context) error
}
