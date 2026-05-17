package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProcessedEventId = uuid.UUID

type ProcessedEvent struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ConsumerName string `gorm:"type:varchar(128);not null;uniqueIndex:idx_consumer_processed_events"`

	// id of outbox or third-party events
	MessageID string `gorm:"type:varchar(128);not null;uniqueIndex:idx_consumer_processed_events"`

	ProcessedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime;index:idx_processed_messages_processed_at"`
}

func (ProcessedEvent) TableName() string { return "processed_messages" }
