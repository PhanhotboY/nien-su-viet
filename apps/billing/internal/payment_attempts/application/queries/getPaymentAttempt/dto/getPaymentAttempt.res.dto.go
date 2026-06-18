package adto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	pahelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

type GetPaymentAttemptResDto interface {
	cdto.ApplicationResponse[PaymentAttemptData, *billing_service.GetPaymentAttemptResponse]
}

type PaymentAttemptData struct {
	ID                    string                      `json:"id"`
	PurchaseID            string                      `json:"purchase_id"`
	Provider              string                      `json:"provider"`
	Status                entity.PaymentAttemptStatus `json:"status"`
	Amount                int64                       `json:"amount"`
	Currency              string                      `json:"currency"`
	ProviderTransactionID string                      `json:"provider_transaction_id"`
	CheckoutURL           string                      `json:"checkout_url"`
	ProviderMetadata      string                      `json:"provider_metadata"`
	ExpiresAt             *time.Time                  `json:"expires_at,omitempty"`
	CreatedAt             time.Time                   `json:"created_at"`
	UpdatedAt             time.Time                   `json:"updated_at"`
}

type getPaymentAttemptResDto struct {
	Data PaymentAttemptData `json:"data"`
}

func NewGetPaymentAttemptResDto(entity entity.PaymentAttempt) GetPaymentAttemptResDto {
	paymentAttempt := PaymentAttemptData{
		ID:                    entity.ID.String(),
		PurchaseID:            entity.PurchaseID.String(),
		Provider:              entity.Provider,
		Status:                entity.Status,
		Amount:                entity.Amount,
		Currency:              entity.Currency,
		ProviderTransactionID: entity.ProviderTransactionID,
		CheckoutURL:           entity.CheckoutURL,
		ProviderMetadata:      entity.ProviderMetadata.String(),
		ExpiresAt:             entity.ExpiresAt,
		CreatedAt:             entity.CreatedAt,
		UpdatedAt:             entity.UpdatedAt,
	}

	return &getPaymentAttemptResDto{
		Data: paymentAttempt,
	}
}

func (dto *getPaymentAttemptResDto) ToGrpcResponse() *billing_service.GetPaymentAttemptResponse {
	return &billing_service.GetPaymentAttemptResponse{
		Data: &billing_service.PaymentAttempt{
			Id:                    dto.Data.ID,
			PurchaseId:            dto.Data.PurchaseID,
			Provider:              dto.Data.Provider,
			Status:                pahelper.ToGrpcStatus(dto.Data.Status),
			Amount:                sdto.NewMoneyDto(dto.Data.Amount, dto.Data.Currency).ToGrpcMoney(),
			ProviderTransactionId: dto.Data.ProviderTransactionID,
			CheckoutUrl:           dto.Data.CheckoutURL,
			ProviderMetadata:      grpcUtils.JsonToStruct(dto.Data.ProviderMetadata),
			ExpiresAt:             grpcUtils.TimeToTimestamp(dto.Data.ExpiresAt),
			CreatedAt:             grpcUtils.TimeToTimestamp(&dto.Data.CreatedAt),
			UpdatedAt:             grpcUtils.TimeToTimestamp(&dto.Data.UpdatedAt),
		},
	}
}

func (dto *getPaymentAttemptResDto) GetData() PaymentAttemptData {
	return dto.Data
}
