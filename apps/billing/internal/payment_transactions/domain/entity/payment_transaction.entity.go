package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type PaymentTransactionType string

const (
	PAYMENT_TRANSACTION_TYPE_PAYMENT    PaymentTransactionType = "payment"
	PAYMENT_TRANSACTION_TYPE_AUTH       PaymentTransactionType = "auth"
	PAYMENT_TRANSACTION_TYPE_CAPTURE    PaymentTransactionType = "capture"
	PAYMENT_TRANSACTION_TYPE_REFUND     PaymentTransactionType = "refund"
	PAYMENT_TRANSACTION_TYPE_VOID       PaymentTransactionType = "void"
	PAYMENT_TRANSACTION_TYPE_CHARGEBACK PaymentTransactionType = "chargeback"
	PAYMENT_TRANSACTION_TYPE_ADJUSTMENT PaymentTransactionType = "adjustment"
	PAYMENT_TRANSACTION_TYPE_FEE        PaymentTransactionType = "fee"
)

type PaymentTransactionStatus string

const (
	PAYMENT_TRANSACTION_STATUS_PENDING   PaymentTransactionStatus = "pending"
	PAYMENT_TRANSACTION_STATUS_SUCCEEDED PaymentTransactionStatus = "succeeded"
	PAYMENT_TRANSACTION_STATUS_FAILED    PaymentTransactionStatus = "failed"
	PAYMENT_TRANSACTION_STATUS_CANCELED  PaymentTransactionStatus = "canceled"
)

type PaymentTransaction struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	PaymentAttemptID uuid.UUID `gorm:"type:uuid;not null;index:idx_payment_transactions_payment_attempt_id"`

	Type   PaymentTransactionType   `gorm:"type:varchar(16);not null;index:idx_payment_transactions_type"`
	Status PaymentTransactionStatus `gorm:"type:varchar(16);not null;index:idx_payment_transactions_status"`

	Amount   int64  `gorm:"type:bigint;not null"`
	Currency string `gorm:"type:varchar(16);not null;"`

	// A unique identifier from the payment provider (e.g., Stripe, PayPal) for this transaction.
	ProviderReference string `gorm:"type:varchar(128);not null;index:idx_payment_transactions_provider_reference"`

	Metadata datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'"`

	ProcessedAt *time.Time `gorm:"type:timestamp;"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;autoCreateTime;"`
}

func (PaymentTransaction) TableName() string { return "payment_transactions" }
