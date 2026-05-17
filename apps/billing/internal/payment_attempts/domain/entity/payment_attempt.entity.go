package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"

	paymentTransactionEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
)

type PaymentAttemptStatus string

const (
	PAYMENT_ATTEMPT_STATUS_CREATED   PaymentAttemptStatus = "created"
	PAYMENT_ATTEMPT_STATUS_PENDING   PaymentAttemptStatus = "pending"
	PAYMENT_ATTEMPT_STATUS_SUCCEEDED PaymentAttemptStatus = "suceeded"
	PAYMENT_ATTEMPT_STATUS_FAILED    PaymentAttemptStatus = "failed"
	PAYMENT_ATTEMPT_STATUS_EXPIRED   PaymentAttemptStatus = "expired"
	PAYMENT_ATTEMPT_STATUS_CANCELED  PaymentAttemptStatus = "canceled"
)

type PaymentAttempt struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	PurchaseID uuid.UUID `gorm:"type:uuid;not null;index:idx_payment_attempts_purchase_id"`

	// e.g. "stripe", "paypal", "momo", etc.
	Provider string               `gorm:"type:varchar(32);not null;index:idx_payment_attempts_provider"`
	Status   PaymentAttemptStatus `gorm:"type:varchar(16);not null;index:idx_payment_attempts_status"`

	Amount   int64  `gorm:"type:bigint;not null"`
	Currency string `gorm:"type:varchar(16);not null;"`

	ProviderTransactionID string `gorm:"type:varchar(128);not null;index:idx_payment_attempts_provider_transaction_id"`

	CheckoutURL string `gorm:"type:text"`

	ProviderMetadata datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`

	ExpiresAt *time.Time `gorm:"type:timestamp;"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;autoUpdateTime;"`

	Transactions []paymentTransactionEntity.PaymentTransaction `gorm:"foreignKey:PaymentAttemptID"`
}

func (PaymentAttempt) TableName() string { return "payment_attempts" }
