package entity

import (
	"time"

	"github.com/google/uuid"
)

type BillingInterval string

const (
	BILLING_INTERVAL_DAY   BillingInterval = "day"
	BILLING_INTERVAL_WEEK  BillingInterval = "week"
	BILLING_INTERVAL_MONTH BillingInterval = "month"
	BILLING_INTERVAL_YEAR  BillingInterval = "year"
)

type Plan struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Code string `gorm:"type:varchar(64);not null;uniqueIndex:uk_subscription_plans_code"`
	Name string `gorm:"type:varchar(255);not null"`

	Price    int64  `gorm:"type:bigint;not null"`
	Currency string `gorm:"type:varchar(16);not null"`

	BillingInterval BillingInterval `gorm:"type:varchar(8);not null;index:idx_subscription_plans_billing_interval"`

	IsActive bool `gorm:"type:boolean;not null;default:true;index:idx_subscription_plans_is_active"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime;"`
}

func (Plan) TableName() string { return "plans" }
