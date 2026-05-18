package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OutboxEventStatus string

const (
	OUTBOX_EVENT_STATUS_PENDING   OutboxEventStatus = "pending"
	OUTBOX_EVENT_STATUS_PUBLISHED OutboxEventStatus = "published"
	OUTBOX_EVENT_STATUS_FAILED    OutboxEventStatus = "failed"
	OUTBOX_EVENT_STATUS_RETRYING  OutboxEventStatus = "retrying"
	OUTBOX_EVENT_STATUS_DEAD      OutboxEventStatus = "dead"
)

type OutboxEvent struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	// Which domain aggregate/entity emitted this event
	AggregateType string `gorm:"type:varchar(64);not null;index:idx_outbox_events_aggregate_type"`
	// ID of the aggregate/entity that emitted this event
	AggregateID uuid.UUID `gorm:"type:uuid;not null;index:idx_outbox_events_aggregate_id"`

	EventType string `gorm:"type:varchar(128);not null;index:idx_outbox_events_event_type"`

	// EventEnvelope stored as JSON in DB (from events.EventEnvelope proto)
	Payload datatypes.JSON `gorm:"type:jsonb;not null"`

	Status     OutboxEventStatus `gorm:"type:varchar(16);not null;index:idx_outbox_events_status"`
	RetryCount int32             `gorm:"type:int;not null;default:0"`

	NextRetryAt *time.Time `gorm:"type:timestamp;"`

	PublishedAt *time.Time `gorm:"type:timestamp;"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;autoCreateTime;"`
}

func (OutboxEvent) TableName() string { return "outbox_events" }
