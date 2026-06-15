package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type InboxEventStatus string

const (
	INBOX_EVENT_STATUS_PENDING   InboxEventStatus = "pending"
	INBOX_EVENT_STATUS_PUBLISHED InboxEventStatus = "published"
	INBOX_EVENT_STATUS_FAILED    InboxEventStatus = "failed"
	INBOX_EVENT_STATUS_RETRYING  InboxEventStatus = "retrying"
	INBOX_EVENT_STATUS_DEAD      InboxEventStatus = "dead"
)

type InboxEvent struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	EventType string `gorm:"type:varchar(128);not null;index:idx_inbox_events_event_type"`

	Provider        string `gorm:"type:varchar(64);not null;uniqueIndex:idx_inbox_provider_event"`
	ExternalEventID string `gorm:"type:varchar(128);uniqueIndex:idx_inbox_provider_event"`

	// EventEnvelope stored as JSON in DB (from events.EventEnvelope proto)
	Payload   datatypes.JSON `gorm:"type:jsonb;not null"`
	Signature string         `gorm:"type:varchar(256);"`

	Status     InboxEventStatus `gorm:"type:varchar(16);not null;index:idx_inbox_events_status"`
	RetryCount int32            `gorm:"type:int;not null;default:0"`

	NextRetryAt *time.Time `gorm:"type:timestamp;"`

	PublishedAt *time.Time `gorm:"type:timestamp;"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;autoCreateTime;"`
}

func (InboxEvent) TableName() string { return "inbox_events" }
