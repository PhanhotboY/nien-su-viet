package utils

import (
	"math"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Pagination struct {
	Limit      uint32 `json:"limit" default:"10"`
	Page       uint32 `json:"page" default:"1"`
	TotalPages uint32 `json:"total_pages" default:"1"`
	Total      uint32 `json:"total" default:"0"`
}

func NewPagination(page, limit, totalItems uint32) *Pagination {
	totalPages := uint32(1)
	if limit > 0 {
		totalPages = uint32(math.Ceil(float64(totalItems) / float64(limit)))
	}
	return &Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		Total:      totalItems,
	}
}

type PaginatedResponse[T any] struct {
	Data       []T         `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

func NewPaginatedResponse[T any](data []T, page, limit, totalItems uint32) *PaginatedResponse[T] {
	return &PaginatedResponse[T]{
		Data:       data,
		Pagination: NewPagination(page, limit, totalItems),
	}
}

type ListQuery struct {
	Page          uint32 `json:"page" validate:"gte=1"`
	Limit         uint32 `json:"limit" validate:"gte=1,lte=100"`
	Search        string
	SortBy        string
	SortOrder     string
	CreatedAtFrom *timestamppb.Timestamp
	CreatedAtTo   *timestamppb.Timestamp
	UpdatedAtFrom *timestamppb.Timestamp
	UpdatedAtTo   *timestamppb.Timestamp
}
