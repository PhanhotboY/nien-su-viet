package dto

type PaginationDto struct {
	Limit      int `json:"limit"`
	Page       int `json:"page"`
	TotalPages int `json:"total_pages"`
	Total      int `json:"total"`
}

func NewPaginationDto(page, limit, totalPages, totalItems int) *PaginationDto {
	return &PaginationDto{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		Total:      totalItems,
	}
}
