package adto

import sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"

type ListPlansReqDto struct {
	OnlyActive bool                  `json:"only_active" validate:"required"`
	Query      sdto.ListQueryRequest `json:"query"`
}

func (d *ListPlansReqDto) ToMap() map[string]any {
	queryMap := d.Query.ToMap()
	queryMap["is_active"] = d.OnlyActive

	return queryMap
}
