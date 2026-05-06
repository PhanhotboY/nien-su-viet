package dto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
)

type UpdateEventDataDto struct {
	Id           entity.HistoricalEventId `json:"id"`                     // Primary key
	Name         *string                  `json:"name,omitempty"`         // Event name
	Thumbnail    *string                  `json:"thumbnail,omitempty"`    // Thumbnail URL
	FromDateType *entity.EventDateType    `json:"fromDateType,omitempty"` // From date type
	FromDay      *int                     `json:"fromDay,omitempty"`      // From day
	FromMonth    *int                     `json:"fromMonth,omitempty"`    // From month
	FromYear     *int                     `json:"fromYear,omitempty"`     // From year
	ToDateType   *entity.EventDateType    `json:"toDateType,omitempty"`   // To date type
	ToDay        *int                     `json:"toDay,omitempty"`        // To day
	ToMonth      *int                     `json:"toMonth,omitempty"`      // To month
	ToYear       *int                     `json:"toYear,omitempty"`       // To year
	AuthorId     *string                  `json:"authorId,omitempty"`     // Author info
	Categories   []any                    `json:"categories,omitempty"`   // Event categories
	Excerpt      *string                  `json:"excerpt,omitempty"`      // Event excerpt
	Content      *string                  `json:"content,omitempty"`      // Full content
	CreatedAt    *time.Time               `json:"createdAt,omitempty"`    // Creation timestamp
	UpdatedAt    *time.Time               `json:"updatedAt,omitempty"`    // Last update timestamp
}

func (d *UpdateEventDataDto) MapToEntity(existingEvent *entity.HistoricalEvent) {
	if d.Name != nil {
		existingEvent.Name = *d.Name
	}
	if d.Thumbnail != nil {
		existingEvent.Thumbnail = d.Thumbnail
	}
	if d.FromDateType != nil {
		existingEvent.FromDateType = *d.FromDateType
	}
	if d.FromDay != nil {
		existingEvent.FromDay = d.FromDay
	}
	if d.FromMonth != nil {
		existingEvent.FromMonth = d.FromMonth
	}
	if d.FromYear != nil {
		existingEvent.FromYear = *d.FromYear
	}
	if d.ToDateType != nil {
		existingEvent.ToDateType = d.ToDateType
	}
	if d.ToDay != nil {
		existingEvent.ToDay = d.ToDay
	}
	if d.ToMonth != nil {
		existingEvent.ToMonth = d.ToMonth
	}
	if d.ToYear != nil {
		existingEvent.ToYear = d.ToYear
	}
	if d.AuthorId != nil {
		existingEvent.AuthorId = *d.AuthorId
	}
	if d.Categories != nil {
		existingEvent.Categories = d.Categories
	}
	if d.Excerpt != nil {
		existingEvent.Excerpt = *d.Excerpt
	}
	if d.Content != nil {
		existingEvent.Content = *d.Content
	}
	if d.CreatedAt != nil {
		existingEvent.CreatedAt = *d.CreatedAt
	}
	if d.UpdatedAt != nil {
		existingEvent.UpdatedAt = *d.UpdatedAt
	} else {
		existingEvent.UpdatedAt = time.Now()
	}
}
