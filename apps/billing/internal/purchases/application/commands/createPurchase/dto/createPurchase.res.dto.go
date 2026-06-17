package adto

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/infrastructure/zalopay"
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"

	getPADto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt/dto"
)

type CreatePurchaseResDto interface {
	cdto.ApplicationResponse[CreatePurchaseResData, *billing_service.CreatePurchaseResponse]
}

type createPurchaseResDto struct {
	Data CreatePurchaseResData `json:"data"`
}

type CreatePurchaseResData struct {
	Id               string `json:"id"`
	ReturnCode       int    `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int    `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`

	PaymentAttempt getPADto.GetPaymentAttemptResDto `json:"payment_attempt"`
}

func NewCreatePurchaseResDto(id string, res *zalopay.CreateOrderResponse, pa getPADto.GetPaymentAttemptResDto) CreatePurchaseResDto {
	return createPurchaseResDto{
		Data: CreatePurchaseResData{
			Id:               id,
			ReturnCode:       res.ReturnCode,
			ReturnMessage:    res.ReturnMessage,
			SubReturnCode:    res.SubReturnCode,
			SubReturnMessage: res.SubReturnMessage,
			PaymentAttempt:   pa,
		},
	}
}

func (r createPurchaseResDto) GetData() CreatePurchaseResData {
	return r.Data
}

func (r createPurchaseResDto) ToGrpcResponse() *billing_service.CreatePurchaseResponse {
	return &billing_service.CreatePurchaseResponse{
		Id:               r.Data.Id,
		ReturnCode:       int32(r.Data.ReturnCode),
		ReturnMessage:    r.Data.ReturnMessage,
		SubReturnCode:    int32(r.Data.SubReturnCode),
		SubReturnMessage: r.Data.SubReturnMessage,

		PaymentAttempt: r.Data.PaymentAttempt.ToGrpcResponse().Data,
	}
}
