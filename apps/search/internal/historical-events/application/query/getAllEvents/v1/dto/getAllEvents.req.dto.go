package dto

import (
	appDto "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
)

type GetAllEventsQueryReq struct {
	appDto.HistoricalEventListQueryDto
}

func (g GetAllEventsQueryReq) MapToQuery() repository.HistoricalEventsQuery {
	return g.HistoricalEventListQueryDto.MapToQuery()
}

func NewGetAllEventsQueryReqWithDefaultValue() *GetAllEventsQueryReq {
	return &GetAllEventsQueryReq{
		HistoricalEventListQueryDto: appDto.NewHistoricalEventListQueryDtoWithDefaultValue(),
	}
}
