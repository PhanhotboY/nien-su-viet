package entity

import (
	"time"

	"github.com/google/uuid"

	paymentAttemptEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	planEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	subscriptionEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
)

type PurchaseStatus string

const (
	PURCHASE_STATUS_PENDING   PurchaseStatus = "pending"
	PURCHASE_STATUS_COMPLETED PurchaseStatus = "completed"
	PURCHASE_STATUS_FAILED    PurchaseStatus = "failed"
	PURCHASE_STATUS_CANCELED  PurchaseStatus = "canceled"
)

type Purchase struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID string `gorm:"type:varchar(124);not null;index:idx_purchases_user_id"`

	SubscriptionID *uuid.UUID                       `gorm:"type:uuid;index:idx_purchases_subscription_id,where:subscription_id IS NOT NULL"`
	Subscription   *subscriptionEntity.Subscription `gorm:"foreignKey:SubscriptionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	PlanID uuid.UUID       `gorm:"type:uuid;not null;index:idx_purchases_plan_id"`
	Plan   planEntity.Plan `gorm:"foreignKey:PlanID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Amount   int64  `gorm:"type:bigint;not null"`
	Currency string `gorm:"type:varchar(16);not null;index:idx_purchases_currency"`

	Status PurchaseStatus `gorm:"type:varchar(32);not null;index:idx_purchases_status"`

	// Idempotent purchase creation (critical)
	IdempotencyKey string `gorm:"type:varchar(128);not null;uniqueIndex:uk_purchases_idempotency_key"`

	CreatedAt   time.Time  `gorm:"type:timestamp;not null;autoCreateTime;index:idx_purchases_created_at"`
	CompletedAt *time.Time `gorm:"type:timestamp;index:idx_purchases_completed_at"`

	PaymentAttempts []paymentAttemptEntity.PaymentAttempt `gorm:"foreignKey:PurchaseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Purchase) TableName() string { return "purchases" }
