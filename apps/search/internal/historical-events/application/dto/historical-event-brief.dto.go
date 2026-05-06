package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	pb "github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/grpc/genproto"
)

// HistoricalEventBrief defines a brief version of historical event entity
type HistoricalEventBriefDto struct {
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
	AuthorId     any               `json:"authorId"`             // Author info
}

func (h *HistoricalEventBriefDto) FromEntity(entity *entity.HistoricalEventBrief) {
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
	h.AuthorId = entity.AuthorId
}

func ConvertHistoricalEventGrpcDateType(dateType entity.EventDateType) pb.EventDateType {
	switch dateType {
	case entity.EventDateTypeExact:
		return pb.EventDateType_EXACT
	case entity.EventDateTypeApproximate:
		fallthrough
	default:
		return pb.EventDateType_APPROXIMATE
	}
}
