package dto

import (
	"time"

	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostListQueryDto struct {
	Page          *uint32                `json:"page,omitempty" validate:"omitempty,gte=1"`
	Limit         *uint32                `json:"limit,omitempty" validate:"omitempty,gte=1,lte=1000"`
	Published     *bool                  `json:"published,omitempty"`
	Search        *string                `json:"search,omitempty" validate:"omitempty,max=255"`
	SortBy        *string                `json:"sort_by,omitempty" validate:"omitempty,max=100"`
	SortOrder     *string                `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"`
	AuthorID      *string                `json:"author_id,omitempty" validate:"omitempty,uuid4"`
	CreatedAtFrom *timestamppb.Timestamp `json:"created_at_from,omitempty"`
	CreatedAtTo   *timestamppb.Timestamp `json:"created_at_to,omitempty" validate:"omitempty,gtfield=CreatedAtFrom"`
	UpdatedAtFrom *timestamppb.Timestamp `json:"updated_at_from,omitempty"`
	UpdatedAtTo   *timestamppb.Timestamp `json:"updated_at_to,omitempty" validate:"omitempty,gtfield=UpdatedAtFrom"`
}

func NewPostListQueryDtoWithDefaultValue() PostListQueryDto {
	var defaultPage, defaultLimit uint32 = 1, 10
	var defaultSortBy, defaultSortOrder string = "created_at", "desc"
	return PostListQueryDto{
		Page:      &defaultPage,
		Limit:     &defaultLimit,
		SortBy:    &defaultSortBy,
		SortOrder: &defaultSortOrder,
	}
}

func (p PostListQueryDto) MapToQuery() repository.PostQuery {
	esQuery := esdsl.NewBoolQuery()

	if p.Search != nil && *p.Search != "" {
		searchQuery := esdsl.NewBoolQuery().Should(
			esdsl.NewMatchQuery("title", *p.Search),
			esdsl.NewMatchQuery("content", *p.Search),
		)
		esQuery = esQuery.Must(searchQuery)
	}

	if p.AuthorID != nil && *p.AuthorID != "" {
		esQuery = esQuery.Filter(esdsl.NewTermQuery("author_id", esdsl.NewFieldValue().String(*p.AuthorID)))
	}

	if p.Published != nil {
		esQuery = esQuery.Filter(esdsl.NewTermQuery("published", esdsl.NewFieldValue().Bool(*p.Published)))
	}

	if p.CreatedAtFrom != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("created_at").Gte(p.CreatedAtFrom.AsTime().Format(time.RFC3339)))
	}
	if p.CreatedAtTo != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("created_at").Lte(p.CreatedAtTo.AsTime().Format(time.RFC3339)))
	}

	if p.UpdatedAtFrom != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("updated_at").Gte(p.UpdatedAtFrom.AsTime().Format(time.RFC3339)))
	}
	if p.UpdatedAtTo != nil {
		esQuery = esQuery.Filter(esdsl.NewDateRangeQuery("updated_at").Lte(p.UpdatedAtTo.AsTime().Format(time.RFC3339)))
	}

	if p.SortBy != nil && *p.SortBy != "" {
		order := "asc"
		if p.SortOrder != nil && *p.SortOrder == "desc" {
			order = "desc"
		}
		_ = order
	}

	return repository.PostQuery{
		QueryVariant: esQuery,
		Page:         *p.Page,
		Limit:        *p.Limit,
		SortBy:       *p.SortBy,
		SortOrder:    *p.SortOrder,
	}
}
