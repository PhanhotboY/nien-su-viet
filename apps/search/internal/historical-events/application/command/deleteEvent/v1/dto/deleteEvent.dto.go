package dto

import "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"

type DeleteEventDataDto struct {
	Id entity.HistoricalEventId `json:"id"` // Primary key
}
