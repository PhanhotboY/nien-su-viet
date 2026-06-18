package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
	oehelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/helper"
)

type CreateOutboxEventReqDto struct {
	// Which domain aggregate/entity emitted this event
	AggregateType string `json:"aggregate_type"`
	// ID of the aggregate/entity that emitted this event
	AggregateID string `json:"aggregate_id"`

	EventType string `json:"event_type"`

	// EventEnvelope stored as JSON in DB (from events.EventEnvelope proto)
	Payload string `json:"payload"`

	Status     *int32 `json:"status"`
	RetryCount *int32 `json:"retry_count"`

	NextRetryAt *time.Time `json:"next_retry_at"`
	LastError   *string    `json:"last_error"`

	PublishedAt *time.Time `json:"published_at"`
}

func (d CreateOutboxEventReqDto) MapToEntity() *entity.OutboxEvent {
	retryCount := int32(0)
	if d.RetryCount != nil {
		retryCount = *d.RetryCount
	}

	return &entity.OutboxEvent{
		AggregateType: d.AggregateType,
		AggregateID:   uuid.MustParse(d.AggregateID),
		EventType:     d.EventType,
		Payload:       []byte(d.Payload),
		Status:        oehelper.ToEntityStatus(d.Status, entity.OUTBOX_EVENT_STATUS_PENDING),
		RetryCount:    retryCount,
		NextRetryAt:   d.NextRetryAt,
		LastError:     d.LastError,
		PublishedAt:   d.PublishedAt,
	}
}
