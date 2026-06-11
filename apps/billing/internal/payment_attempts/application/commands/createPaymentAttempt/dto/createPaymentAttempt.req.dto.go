package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
)

type CreatePaymentAttemptReqDto struct {
	PurchaseID            string        `json:"purchase_id" validate:"required,uuid4"`
	Provider              string        `json:"provider" validate:"required"`
	Amount                sdto.MoneyDto `json:"amount" validate:"required"`
	ProviderTransactionID string        `json:"provider_transaction_id" validate:"required"`
	CheckoutURL           string        `json:"checkout_url"`
	ProviderMetadata      string        `json:"provider_metadata"`
	ExpiresAt             *time.Time    `json:"expires_at" validate:"omitempty"`
}

func (dto *CreatePaymentAttemptReqDto) MapToEntity() *entity.PaymentAttempt {
	expiresAt := time.Now().Add(15 * time.Minute) // Default expiration time is 15 minutes from now
	if dto.ExpiresAt != nil {
		expiresAt = *dto.ExpiresAt
	}
	return &entity.PaymentAttempt{
		PurchaseID:            uuid.MustParse(dto.PurchaseID),
		Provider:              dto.Provider,
		Amount:                dto.Amount.Amount,
		Currency:              dto.Amount.Currency,
		Status:                entity.PAYMENT_ATTEMPT_STATUS_CREATED,
		ProviderTransactionID: dto.ProviderTransactionID,
		CheckoutURL:           dto.CheckoutURL,
		ProviderMetadata:      []byte(dto.ProviderMetadata),
		ExpiresAt:             &expiresAt,
	}
}
