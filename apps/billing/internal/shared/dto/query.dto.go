package sdto

type ListQueryRequest struct {
	Page      int32  `json:"page,omitempty" validate:"omitempty,gte=1"`
	Limit     int32  `json:"limit,omitempty" validate:"omitempty,gte=1,lte=1000"`
	Search    string `json:"search,omitempty" validate:"omitempty,max=255"`
	SortBy    string `json:"sort_by,omitempty"`                                        // e.g., "created_at", "name", etc.
	SortOrder string `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"` // "asc" or "desc"
}

func (q ListQueryRequest) ToMap() map[string]any {
	queryMap := make(map[string]any)

	if q.Page > 0 {
		queryMap["page"] = q.Page
	}
	if q.Limit > 0 {
		queryMap["limit"] = q.Limit
	}
	if q.Search != "" {
		queryMap["search"] = q.Search
	}
	if q.SortBy != "" {
		queryMap["sort_by"] = q.SortBy
	}
	if q.SortOrder != "" {
		queryMap["sort_order"] = q.SortOrder
	}

	return queryMap
}
