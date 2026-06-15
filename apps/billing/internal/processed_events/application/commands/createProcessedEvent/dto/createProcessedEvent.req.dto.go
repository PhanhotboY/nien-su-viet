package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/entity"
)

type CreateProcessedEventReqDto struct {
	ID string `json:"id" validate:"required,uuid4"`

	ConsumerName string `json:"consumer_name" validate:"required"`

	// id of outbox or third-party events
	MessageID string `json:"message_id" validate:"required"`

	ProcessedAt time.Time `json:"processed_at" validate:"required"`
}

func (d CreateProcessedEventReqDto) MapToEntity() *entity.ProcessedEvent {
	return &entity.ProcessedEvent{
		ID:           uuid.MustParse(d.ID),
		ConsumerName: d.ConsumerName,
		MessageID:    d.MessageID,
		ProcessedAt:  d.ProcessedAt,
	}
}
