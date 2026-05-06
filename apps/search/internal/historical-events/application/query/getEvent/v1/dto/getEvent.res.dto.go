package dto

import (
	heventDto "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
)

type GetEventRes struct {
	Data heventDto.HistoricalEventDto `json:"data"`
}

func (p *GetEventRes) FromEntity(entity *entity.HistoricalEvent) {
	p.Data.FromEntity(entity)
}
