package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
	pthelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

type CreatePaymentTransactionReqDto struct {
	PaymentAttemptID string `json:"payment_attempt_id" validate:"required,uuid4"`

	Type   int32 `json:"type" validate:"required"`
	Status int32 `json:"status" validate:"required"`

	Price sdto.MoneyDto `json:"price"`

	// A unique identifier from the payment provider (e.g., Stripe, PayPal) for this transaction.
	ProviderReference string `json:"provider_reference" validate:"required"`

	Metadata map[string]any `json:"metadata,omitempty"`

	ProcessedAt *time.Time `json:"processed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (r *CreatePaymentTransactionReqDto) MapToEntity() *entity.PaymentTransaction {
	return &entity.PaymentTransaction{
		PaymentAttemptID: uuid.MustParse(r.PaymentAttemptID),
		Type:             pthelper.ToEntityTransactionType(billing_service.PaymentTransactionType(r.Type)),
		Status:           pthelper.ToEntityTransactionStatus(billing_service.PaymentTransactionStatus(r.Status)),
		Amount:           r.Price.Amount,
		Currency:         r.Price.Currency,
		CreatedAt:        time.Now(),
	}
}
