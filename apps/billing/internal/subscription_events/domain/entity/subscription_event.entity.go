package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type SubscriptionEvent struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	SubscriptionID uuid.UUID `gorm:"type:uuid;not null;index:idx_subscription_events_subscription_id"`

	EventType string `gorm:"type:varchar(128);not null;index:idx_subscription_events_event_type"`

	Payload datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime;index:idx_subscription_events_created_at"`
}

func (SubscriptionEvent) TableName() string { return "subscription_events" }
