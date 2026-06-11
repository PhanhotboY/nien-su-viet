package adto

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/zalopay"
	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
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
	OrderUrl         string `json:"order_url"`
	Zptranstoken     string `json:"zptranstoken"`
	OrderToken       string `json:"order_token"`
	QrCode           string `json:"qr_code"`
}

func NewCreatePurchaseResDto(id string, res *zalopay.CreateOrderResponse) CreatePurchaseResDto {
	return createPurchaseResDto{
		Data: CreatePurchaseResData{
			Id:               id,
			ReturnCode:       res.ReturnCode,
			ReturnMessage:    res.ReturnMessage,
			SubReturnCode:    res.SubReturnCode,
			SubReturnMessage: res.SubReturnMessage,
			OrderUrl:         res.OrderURL,
			Zptranstoken:     res.Zptranstoken,
			OrderToken:       res.OrderToken,
			QrCode:           res.QRCode,
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
		OrderUrl:         r.Data.OrderUrl,
		Zptranstoken:     r.Data.Zptranstoken,
		OrderToken:       r.Data.OrderToken,
		QrCode:           r.Data.QrCode,
	}
}
