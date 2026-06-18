package adto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	purhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

type GetPurchaseResDto interface {
	cdto.ApplicationResponse[PurchaseDto, *billing_service.GetPurchaseResponse]
}

type PurchaseDto struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	SubscriptionID string `json:"subscription_id"`
	PlanID         string `json:"plan_id"`

	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`

	Status billing_service.PurchaseStatus `json:"status"`

	// Idempotent purchase creation (critical)
	IdempotencyKey string `json:"idempotency_key"`

	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type getPurchaseResDto struct {
	Data PurchaseDto `json:"data"`
}

func NewGetPurchaseResDto(entity entity.Purchase) GetPurchaseResDto {
	subscriptionId := ""
	if entity.SubscriptionID != nil {
		subscriptionId = entity.SubscriptionID.String()
	}

	purchase := PurchaseDto{
		ID:             entity.ID.String(),
		UserID:         entity.UserID,
		SubscriptionID: subscriptionId,
		PlanID:         entity.PlanID.String(),

		Amount:   entity.Amount,
		Currency: entity.Currency,

		Status: purhelper.ToGrpcStatus(entity.Status),

		IdempotencyKey: entity.IdempotencyKey,

		CreatedAt:   entity.CreatedAt,
		CompletedAt: entity.CompletedAt,
	}

	return getPurchaseResDto{
		Data: purchase,
	}
}

func (r getPurchaseResDto) GetData() PurchaseDto {
	return r.Data
}

func (r getPurchaseResDto) ToGrpcResponse() *billing_service.GetPurchaseResponse {
	return &billing_service.GetPurchaseResponse{
		Data: &billing_service.Purchase{
			Id:             r.Data.ID,
			UserId:         r.Data.UserID,
			SubscriptionId: r.Data.SubscriptionID,
			PlanId:         r.Data.PlanID,

			Amount: sdto.NewMoneyDto(r.Data.Amount, r.Data.Currency).ToGrpcMoney(),
			Status: r.Data.Status,

			IdempotencyKey: r.Data.IdempotencyKey,

			CreatedAt:   grpcUtils.TimeToTimestamp(&r.Data.CreatedAt),
			CompletedAt: grpcUtils.TimeToTimestamp(r.Data.CompletedAt),
		},
	}
}
