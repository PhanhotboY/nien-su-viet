package repository

import (
	"context"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
)

type HistoricalEventsQuery struct {
	types.QueryVariant

	Page      uint32 `json:"page,omitempty" validate:"omitempty,gte=1"`
	Limit     uint32 `json:"limit,omitempty" validate:"omitempty,gte=1,lte=1000"`
	SortBy    string `json:"sort_by,omitempty" validate:"omitempty,max=100"`
	SortOrder string `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"` // "asc" or "desc"
}

type HistoricalEventsSearchResponse struct {
	HistoricalEvents []entity.HistoricalEventBrief `json:"historical_events"`
	TotalCount       int64                         `json:"total_count"`
}

type HistoricalEventEsRepository interface {
	IndexHistoricalEvent(ctx context.Context, event entity.HistoricalEvent) error
	GetHistoricalEventById(ctx context.Context, id string) (*entity.HistoricalEvent, error)
	SearchHistoricalEvents(ctx context.Context, query HistoricalEventsQuery) (*HistoricalEventsSearchResponse, error)
	UpdateHistoricalEvent(ctx context.Context, id entity.HistoricalEventId, event *entity.HistoricalEvent) error
	DeleteHistoricalEvent(ctx context.Context, id entity.HistoricalEventId) error
}
