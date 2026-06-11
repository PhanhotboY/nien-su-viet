package adto

import (
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/common"
)

type CreatePaymentTransactionResDto interface {
	cdto.ApplicationResponse[CreatePaymentTransactionResData, *billing_service.CreatePaymentAttemptResponse]
}

type CreatePaymentTransactionResData struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type createPaymentTransactionResDto struct {
	Data CreatePaymentTransactionResData `json:"data"`
}

func NewCreatePaymentTransactionResDto(id string, success bool, msg string) CreatePaymentTransactionResDto {
	return &createPaymentTransactionResDto{
		Data: CreatePaymentTransactionResData{id, success, msg},
	}
}

func (r *createPaymentTransactionResDto) GetData() CreatePaymentTransactionResData {
	return r.Data
}

func (r *createPaymentTransactionResDto) ToGrpcResponse() *billing_service.CreatePaymentAttemptResponse {
	return &billing_service.CreatePaymentAttemptResponse{
		Data: &common.OperationMetadata{
			Id:      r.Data.ID,
			Success: r.Data.Success,
		},
	}
}
