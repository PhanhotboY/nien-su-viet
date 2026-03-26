package dto

import "math"

type PaginationDto struct {
	Limit      uint32 `json:"limit"`
	Page       uint32 `json:"page"`
	TotalPages uint32 `json:"total_pages"`
	Total      uint32 `json:"total"`
}

func NewPaginationDto(page, limit, totalItems uint32) *PaginationDto {
	totalPages := uint32(1)
	if limit > 0 {
		totalPages = uint32(math.Ceil(float64(totalItems) / float64(limit)))
	}
	return &PaginationDto{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		Total:      totalItems,
	}
}
