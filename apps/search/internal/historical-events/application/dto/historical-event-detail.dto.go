package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	pb "github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/grpc/genproto"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

// HistoricalEvent defines the historical event entity
type HistoricalEventDto struct {
	Id           string            `json:"id"`                   // Primary key
	Name         string            `json:"name"`                 // Event name
	Thumbnail    *string           `json:"thumbnail,omitempty"`  // Thumbnail URL
	FromDateType pb.EventDateType  `json:"fromDateType"`         // From date type
	FromDay      *int              `json:"fromDay,omitempty"`    // From day
	FromMonth    *int              `json:"fromMonth,omitempty"`  // From month
	FromYear     int               `json:"fromYear"`             // From year
	ToDateType   *pb.EventDateType `json:"toDateType,omitempty"` // To date type
	ToDay        *int              `json:"toDay,omitempty"`      // To day
	ToMonth      *int              `json:"toMonth,omitempty"`    // To month
	ToYear       *int              `json:"toYear,omitempty"`     // To year
	Author       any               `json:"author"`               // Author info
	Categories   []any             `json:"categories"`           // Event categories
	Excerpt      string            `json:"excerpt"`              // Event excerpt
	Content      string            `json:"content"`              // Full content
	CreatedAt    string            `json:"createdAt"`            // Creation timestamp
	UpdatedAt    string            `json:"updatedAt"`            // Last update timestamp
}

func (h *HistoricalEventDto) FromEntity(entity *entity.HistoricalEvent) {
	h.Id = entity.Id
	h.Name = entity.Name
	h.Thumbnail = entity.Thumbnail
	h.FromDateType = ConvertHistoricalEventGrpcDateType(entity.FromDateType)
	h.FromDay = entity.FromDay
	h.FromMonth = entity.FromMonth
	h.FromYear = entity.FromYear
	if entity.ToDateType != nil {
		toDateType := ConvertHistoricalEventGrpcDateType(*entity.ToDateType)
		h.ToDateType = &toDateType
	}
	h.ToDay = entity.ToDay
	h.ToMonth = entity.ToMonth
	h.ToYear = entity.ToYear
	// h.AuthorId = entity.AuthorId
	h.Categories = entity.Categories
	h.Excerpt = entity.Excerpt
	h.Content = entity.Content
	h.CreatedAt = grpcUtils.TimeToString(entity.CreatedAt)
	h.UpdatedAt = grpcUtils.TimeToString(entity.UpdatedAt)
}
