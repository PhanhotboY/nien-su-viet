package dto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostListQueryDto struct {
	Page          uint32                 `json:"page" validate:"gte=1"`
	Limit         uint32                 `json:"limit" validate:"gte=1,lte=100"`
	Published     *bool                  `json:"published,omitempty"`
	Search        *string                `json:"search,omitempty" validate:"omitempty,max=255"`
	SortBy        *string                `json:"sort_by,omitempty"`
	SortOrder     *string                `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"`
	AuthorID      *string                `json:"author_id,omitempty" validate:"omitempty,uuid4"`
	CreatedAtFrom *timestamppb.Timestamp `json:"created_at_from,omitempty"`
	CreatedAtTo   *timestamppb.Timestamp `json:"created_at_to,omitempty" validate:"omitempty,gtfield=CreatedAtFrom"`
	UpdatedAtFrom *timestamppb.Timestamp `json:"updated_at_from,omitempty"`
	UpdatedAtTo   *timestamppb.Timestamp `json:"updated_at_to,omitempty" validate:"omitempty,gtfield=UpdatedAtFrom"`
}

func (g PostListQueryDto) MapToQuery() repository.PostQuery {
	var createdAtFrom *time.Time
	if g.CreatedAtFrom != nil {
		t := g.CreatedAtFrom.AsTime()
		createdAtFrom = &t
	}
	var createdAtTo *time.Time
	if g.CreatedAtTo != nil {
		t := g.CreatedAtTo.AsTime()
		createdAtTo = &t
	}
	var updatedAtFrom *time.Time
	if g.UpdatedAtFrom != nil {
		t := g.UpdatedAtFrom.AsTime()
		updatedAtFrom = &t
	}
	var updatedAtTo *time.Time
	if g.UpdatedAtTo != nil {
		t := g.UpdatedAtTo.AsTime()
		updatedAtTo = &t
	}
	var search string
	if g.Search != nil {
		search = *g.Search
	}
	var sortBy string
	if g.SortBy != nil {
		sortBy = *g.SortBy
	}
	var sortOrder string
	if g.SortOrder != nil {
		sortOrder = *g.SortOrder
	}
	var authorID string
	if g.AuthorID != nil {
		authorID = *g.AuthorID
	}
	var limit uint32 = 10
	if g.Limit > 0 {
		limit = g.Limit
	}
	var offset uint32 = 10
	if g.Page > 0 {
		offset = limit * (g.Page - 1)
	}
	return repository.PostQuery{
		Published:     g.Published,
		Search:        search,
		SortBy:        sortBy,
		SortOrder:     sortOrder,
		AuthorID:      authorID,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		PostPagination: repository.PostPagination{
			Limit:  g.Limit,
			Offset: offset,
		},
	}
}
