package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
)

type GetAllEventsRes struct {
	Data       []dto.HistoricalEventBriefDto `json:"data"`
	Pagination utils.Pagination              `json:"pagination"`
}

func NewGetAllEventsRes(events []dto.HistoricalEventBriefDto, pagination utils.Pagination) *GetAllEventsRes {
	return &GetAllEventsRes{
		Data:       events,
		Pagination: pagination,
	}
}
