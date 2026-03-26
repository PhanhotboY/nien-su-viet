package dto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostListQueryDto struct {
	Page          uint32 `json:"page" validate:"gte=1"`
	Limit         uint32 `json:"limit" validate:"gte=1,lte=100"`
	Published     *bool  `json:"published,omitempty"`
	Search        string
	SortBy        string
	SortOrder     string
	AuthorID      string
	CreatedAtFrom *timestamppb.Timestamp
	CreatedAtTo   *timestamppb.Timestamp
	UpdatedAtFrom *timestamppb.Timestamp
	UpdatedAtTo   *timestamppb.Timestamp
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
	return repository.PostQuery{
		Published:     g.Published,
		Search:        g.Search,
		SortBy:        g.SortBy,
		SortOrder:     g.SortOrder,
		AuthorID:      g.AuthorID,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
	}
}

func (g PostListQueryDto) MapToPagination() repository.PostPagination {
	return repository.PostPagination{
		Limit:  g.Limit,
		Offset: g.Limit * (g.Page - 1),
	}
}
