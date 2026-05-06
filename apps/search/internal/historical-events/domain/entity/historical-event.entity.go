package entity

import (
	"time"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type HistoricalEventId = string

type EventDateType string

const (
	EventDateTypeExact       EventDateType = "EXACT"
	EventDateTypeApproximate EventDateType = "APPROXIMATE"
)

// HistoricalEvent defines the historical event entity
type HistoricalEvent struct {
	Id           HistoricalEventId `json:"id"`                   // Primary key
	Name         string            `json:"name"`                 // Event name
	Thumbnail    *string           `json:"thumbnail,omitempty"`  // Thumbnail URL
	FromDateType EventDateType     `json:"fromDateType"`         // From date type
	FromDay      *int              `json:"fromDay,omitempty"`    // From day
	FromMonth    *int              `json:"fromMonth,omitempty"`  // From month
	FromYear     int               `json:"fromYear"`             // From year
	ToDateType   *EventDateType    `json:"toDateType,omitempty"` // To date type
	ToDay        *int              `json:"toDay,omitempty"`      // To day
	ToMonth      *int              `json:"toMonth,omitempty"`    // To month
	ToYear       *int              `json:"toYear,omitempty"`     // To year
	AuthorId     string            `json:"authorId"`             // Author info
	Categories   []any             `json:"categories"`           // Event categories
	Excerpt      string            `json:"excerpt"`              // Event excerpt
	Content      string            `json:"content"`              // Full content
	CreatedAt    time.Time         `json:"createdAt"`            // Creation timestamp
	UpdatedAt    time.Time         `json:"updatedAt"`            // Last update timestamp
}

// HistoricalEventBrief defines a brief version of historical event entity
type HistoricalEventBrief struct {
	Id           HistoricalEventId `json:"id"`                   // Primary key
	Name         string            `json:"name"`                 // Event name
	Thumbnail    *string           `json:"thumbnail,omitempty"`  // Thumbnail URL
	FromDateType EventDateType     `json:"fromDateType"`         // From date type
	FromDay      *int              `json:"fromDay,omitempty"`    // From day
	FromMonth    *int              `json:"fromMonth,omitempty"`  // From month
	FromYear     int               `json:"fromYear"`             // From year
	ToDateType   *EventDateType    `json:"toDateType,omitempty"` // To date type
	ToDay        *int              `json:"toDay,omitempty"`      // To day
	ToMonth      *int              `json:"toMonth,omitempty"`    // To month
	ToYear       *int              `json:"toYear,omitempty"`     // To year
	AuthorId     any               `json:"authorId"`             // Author info
}

// HistoricalEventPreview includes excerpt and categories
type HistoricalEventPreview struct {
	Id           HistoricalEventId `json:"id"`                   // Primary key
	Name         string            `json:"name"`                 // Event name
	Thumbnail    *string           `json:"thumbnail,omitempty"`  // Thumbnail URL
	FromDateType EventDateType     `json:"fromDateType"`         // From date type
	FromDay      *int              `json:"fromDay,omitempty"`    // From day
	FromMonth    *int              `json:"fromMonth,omitempty"`  // From month
	FromYear     int               `json:"fromYear"`             // From year
	ToDateType   *EventDateType    `json:"toDateType,omitempty"` // To date type
	ToDay        *int              `json:"toDay,omitempty"`      // To day
	ToMonth      *int              `json:"toMonth,omitempty"`    // To month
	ToYear       *int              `json:"toYear,omitempty"`     // To year
	Author       any               `json:"author"`               // Author info
	Categories   []any             `json:"categories"`           // Event categories
	Excerpt      string            `json:"excerpt"`              // Event excerpt
}

func (HistoricalEvent) IndexName() string {
	return "historical_events"
}

func (HistoricalEvent) ToTypeMapping() map[string]types.Property {
	return map[string]types.Property{
		"id":           types.KeywordProperty{},
		"name":         types.TextProperty{},
		"thumbnail":    types.KeywordProperty{},
		"fromDateType": types.KeywordProperty{},
		"fromDay":      types.IntegerNumberProperty{},
		"fromMonth":    types.IntegerNumberProperty{},
		"fromYear":     types.IntegerNumberProperty{},
		"toDateType":   types.KeywordProperty{},
		"toDay":        types.IntegerNumberProperty{},
		"toMonth":      types.IntegerNumberProperty{},
		"toYear":       types.IntegerNumberProperty{},
		"author":       types.ObjectProperty{},
		"categories":   types.ObjectProperty{},
		"excerpt":      types.TextProperty{},
		"content":      types.TextProperty{},
		"createdAt":    types.DateProperty{},
		"updatedAt":    types.DateProperty{},
	}
}

func (HistoricalEventBrief) IndexName() string {
	return "historical_events"
}

func (HistoricalEventPreview) IndexName() string {
	return "historical_events"
}
