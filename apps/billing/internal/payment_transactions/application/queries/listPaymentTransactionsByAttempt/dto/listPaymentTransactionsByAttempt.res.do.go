package adto

import (
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	pb "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
	pthelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/helper"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
)

type ListPaymentTransactionsByAttemptResDto interface {
	cdto.ApplicationResponse[[]PaymentTransaction, *pb.ListPaymentTransactionsByAttemptResponse]
}

type PaymentTransaction struct {
	Id                string                          `json:"id,omitempty"`
	PaymentAttemptId  string                          `json:"payment_attempt_id,omitempty"`
	Type              entity.PaymentTransactionType   `json:"type,omitempty"`
	Status            entity.PaymentTransactionStatus `json:"status,omitempty"`
	Amount            *sdto.MoneyDto                  `json:"amount,omitempty"`
	ProviderReference string                          `json:"provider_reference,omitempty"`
	Metadata          string                          `json:"metadata,omitempty"`
	ProcessedAt       *timestamppb.Timestamp          `json:"processed_at,omitempty"`
	CreatedAt         *timestamppb.Timestamp          `json:"created_at,omitempty"`
}

type listPaymentTransactionsByAttemptResDto struct {
	Data []PaymentTransaction
}

func NewListPaymentTransactionsByAttemptResDto(transactions []*entity.PaymentTransaction) ListPaymentTransactionsByAttemptResDto {
	var data = make([]PaymentTransaction, len(transactions))
	for i, t := range transactions {
		data[i] = PaymentTransaction{
			Id:                t.ID.String(),
			PaymentAttemptId:  t.PaymentAttemptID.String(),
			Type:              t.Type,
			Status:            t.Status,
			Amount:            sdto.NewMoneyDto(t.Amount, t.Currency),
			ProviderReference: t.ProviderReference,
			Metadata:          t.Metadata.String(),
			ProcessedAt:       grpcUtils.TimeToTimestamp(t.ProcessedAt),
			CreatedAt:         grpcUtils.TimeToTimestamp(&t.CreatedAt),
		}
	}
	return listPaymentTransactionsByAttemptResDto{Data: data}
}

func (r listPaymentTransactionsByAttemptResDto) GetData() []PaymentTransaction {
	return r.Data
}

func (r listPaymentTransactionsByAttemptResDto) ToGrpcResponse() *pb.ListPaymentTransactionsByAttemptResponse {
	var grpcTransactions = make([]*pb.PaymentTransaction, len(r.Data))
	for i, t := range r.Data {
		grpcTransactions[i] = &pb.PaymentTransaction{
			Id:                t.Id,
			PaymentAttemptId:  t.PaymentAttemptId,
			Type:              pthelper.ToGrpcTransactionType(t.Type),
			Status:            pthelper.ToGrpcTransactionStatus(t.Status),
			Amount:            t.Amount.ToGrpcMoney(),
			ProviderReference: t.ProviderReference,
			Metadata:          grpcUtils.JsonToStruct(string(t.Metadata)),
			ProcessedAt:       t.ProcessedAt,
			CreatedAt:         t.CreatedAt,
		}
	}
	return &pb.ListPaymentTransactionsByAttemptResponse{Data: grpcTransactions}
}
