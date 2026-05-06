package dto

import (
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HistoricalEventListQueryDto struct {
	Page      *uint32 `json:"page,omitempty" validate:"omitempty,gte=1"`
	Limit     *uint32 `json:"limit,omitempty" validate:"omitempty,gte=1,lte=1000"`
	Search    *string `json:"search,omitempty" validate:"omitempty,max=255"`
	SortBy    *string `json:"sort_by,omitempty" validate:"omitempty,max=100"`
	SortOrder *string `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"` // "asc" or "desc"
	AuthorId  *string `json:"author_id,omitempty" validate:"omitempty,uuid4"`
	// Filter by category
	CategoryIds []string `json:"category_ids,omitempty" validate:"omitempty,dive,uuid4"`
	// Filter by from date range
	FromYear  *int32 `json:"from_year,omitempty" validate:"omitempty,gte=1,lte=9999"`
	FromMonth *int32 `json:"from_month,omitempty" validate:"omitempty,gte=1,lte=12"`
	FromDay   *int32 `json:"from_day,omitempty" validate:"omitempty,gte=1,lte=31"`
	// Filter by to date range
	ToYear  *int32 `json:"to_year,omitempty" validate:"omitempty,gte=1,lte=9999"`
	ToMonth *int32 `json:"to_month,omitempty" validate:"omitempty,gte=1,lte=12"`
	ToDay   *int32 `json:"to_day,omitempty" validate:"omitempty,gte=1,lte=31"`
	// Search specific year
	SearchYear *int32 `json:"search_year,omitempty" validate:"omitempty,gte=1,lte=9999"`
	// Filter by created date range
	CreatedAtFrom *timestamppb.Timestamp `json:"created_at_from,omitempty"`
	CreatedAtTo   *timestamppb.Timestamp `json:"created_at_to,omitempty" validate:"omitempty,gtfield=CreatedAtFrom"`
	// Filter by updated date range
	UpdatedAtFrom *timestamppb.Timestamp `json:"updated_at_from,omitempty"`
	UpdatedAtTo   *timestamppb.Timestamp `json:"updated_at_to,omitempty" validate:"omitempty,gtfield=UpdatedAtFrom"`
}

func NewHistoricalEventListQueryDtoWithDefaultValue() HistoricalEventListQueryDto {
	var defaultPage, defaultLimit uint32 = 1, 10
	var defaultSortBy, defaultSortOrder string = "createdAt", "desc"
	return HistoricalEventListQueryDto{
		Page:      &defaultPage,
		Limit:     &defaultLimit,
		SortBy:    &defaultSortBy,
		SortOrder: &defaultSortOrder,
	}
}

func (h HistoricalEventListQueryDto) MapToQuery() repository.HistoricalEventsQuery {
	esQuery := esdsl.NewBoolQuery()

	if h.Search != nil && *h.Search != "" {
		searchQuery := esdsl.NewBoolQuery().Should(
			esdsl.NewMatchQuery("name", *h.Search),
			esdsl.NewMatchQuery("description", *h.Search),
			esdsl.NewMatchQuery("content", *h.Search),
		)
		esQuery = esQuery.Must(searchQuery)
	}

	if h.AuthorId != nil && *h.AuthorId != "" {
		esQuery = esQuery.Filter(esdsl.NewTermQuery("authorId", esdsl.NewFieldValue().String(*h.AuthorId)))
	}

	if len(h.CategoryIds) > 0 {
		categoryQuery := esdsl.NewBoolQuery()
		for _, id := range h.CategoryIds {
			categoryQuery = categoryQuery.Should(esdsl.NewTermQuery("categoryIds", esdsl.NewFieldValue().String(id)))
		}
		esQuery = esQuery.Filter(categoryQuery)
	}

	if h.SearchYear != nil {
		esQuery = esQuery.Filter(esdsl.NewTermQuery("year", esdsl.NewFieldValue().Int64(int64(*h.SearchYear))))
	}

	if h.FromYear != nil || h.FromMonth != nil || h.FromDay != nil {
		from := fmt.Sprintf("%04d-%02d-%02d", valueOrDefaultInt32(h.FromYear, 1), valueOrDefaultInt32(h.FromMonth, 1), valueOrDefaultInt32(h.FromDay, 1))
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("date").Gte(from))
	}

	if h.ToYear != nil || h.ToMonth != nil || h.ToDay != nil {
		to := fmt.Sprintf("%04d-%02d-%02d", valueOrDefaultInt32(h.ToYear, 9999), valueOrDefaultInt32(h.ToMonth, 12), valueOrDefaultInt32(h.ToDay, 31))
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("date").Lte(to))
	}

	if h.CreatedAtFrom != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("createdAt").Gte(h.CreatedAtFrom.AsTime().Format(time.RFC3339)))
	}
	if h.CreatedAtTo != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("createdAt").Lte(h.CreatedAtTo.AsTime().Format(time.RFC3339)))
	}
	if h.UpdatedAtFrom != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("updatedAt").Gte(h.UpdatedAtFrom.AsTime().Format(time.RFC3339)))
	}
	if h.UpdatedAtTo != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("updatedAt").Lte(h.UpdatedAtTo.AsTime().Format(time.RFC3339)))
	}

	return repository.HistoricalEventsQuery{
		QueryVariant: esQuery,
		Page:         *h.Page,
		Limit:        *h.Limit,
		SortBy:       *h.SortBy,
		SortOrder:    *h.SortOrder,
	}
}

func valueOrDefaultInt32(v *int32, def int32) int32 {
	if v == nil {
		return def
	}
	return *v
}
