package entity

import (
	"time"

	"github.com/google/uuid"

	planEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
)

type SubscriptionStatus string

const (
	SUBSCRIPTION_STATUS_PENDING  SubscriptionStatus = "pending"
	SUBSCRIPTION_STATUS_ACTIVE   SubscriptionStatus = "active"
	SUBSCRIPTION_STATUS_PAST_DUE SubscriptionStatus = "past_due"
	SUBSCRIPTION_STATUS_CANCELED SubscriptionStatus = "canceled"
	SUBSCRIPTION_STATUS_EXPIRED  SubscriptionStatus = "expired"
)

type Subscription struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID `gorm:"type:uuid;not null;index:idx_subscriptions_user_id"`

	PlanID uuid.UUID       `gorm:"type:uuid;not null;index:idx_subscriptions_plan_id"`
	Plan   planEntity.Plan `gorm:"foreignKey:PlanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Status SubscriptionStatus `gorm:"type:varchar(16);not null;index:idx_subscriptions_status"`

	// Very common query: find the current subscription that covers now
	CurrentPeriodStart time.Time `gorm:"type:timestamp;not null;index:idx_subscriptions_current_period_start"`
	CurrentPeriodEnd   time.Time `gorm:"type:timestamp;not null;index:idx_subscriptions_current_period_end"`

	CancelAtPeriodEnd bool `gorm:"type:boolean;not null;default:false;index:idx_subscriptions_cancel_at_period_end"`

	CanceledAt *time.Time `gorm:"type:timestamp;index:idx_subscriptions_canceled_at"`
	ExpiredAt  *time.Time `gorm:"type:timestamp;index:idx_subscriptions_expired_at"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime;index:idx_subscriptions_created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;autoUpdateTime;index:idx_subscriptions_updated_at"`
}

func (Subscription) TableName() string { return "subscriptions" }
