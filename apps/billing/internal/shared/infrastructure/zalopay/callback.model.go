package zalopay

import (
	"encoding/json"
	"fmt"

	cdto "github.com/phanhotboy/nien-su-viet/libs/pkg/core/application/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

type CallbackData struct {
	AppID      int    `json:"app_id"`
	AppTransID string `json:"app_trans_id"`
	AppTime    int64  `json:"app_time"`
	Amount     int64  `json:"amount"`
	AppUser    string `json:"app_user"`
	EmbedData  string `json:"embed_data"`
	Item       string `json:"item"`
	ZpTransID  int64  `json:"zp_trans_id"`
	// https://developers.zalopay.vn/v2/general/overview.html#callback_dac-ta-api_cac-kenh-thanh-toan-ho-tro
	Channel        int   `json:"channel"`
	UseFeeAmount   int64 `json:"use_fee_amount"`
	DiscountAmount int64 `json:"discount_amount"`

	ServerTime     int64  `json:"server_time"`
	MerchantUserID string `json:"merchant_user_id"`
}

// https://developers.zalopay.vn/v2/general/overview.html#callback_dac-ta-api_du-lieu-nhan-duoc-tu-callback
type CallbackPayload struct {
	Data string `json:"data"`
	Mac  string `json:"mac"`
	Type int32  `json:"type"` // 1: Order, 2: Agreement
}

func NewCallbackPayload(req *billing_service.CallbackPayload) CallbackPayload {
	return CallbackPayload{
		Data: req.Data,
		Mac:  req.Mac,
		Type: req.Type,
	}
}

func (p CallbackPayload) GetData() CallbackData {
	var data CallbackData
	if err := json.Unmarshal([]byte(p.Data), &data); err != nil {
		fmt.Println("Failed to unmarshal callback Data: ", err)
		return CallbackData{}
	}
	return data
}

type CallbackResponse interface {
	cdto.ApplicationResponse[CallbackResponseData, *billing_service.CallbackResponse]
}

// https://developers.zalopay.vn/v2/general/overview.html#callback_dac-ta-api_thong-tin-appserver-tra-ve-cho-zalopayserver-khi-nhan-callback
type CallbackResponseData struct {
	ReturnCode    int32  `json:"return_code"` // -1: MAC not equal, 0: Internal Error, 1: Success, 2: Duplicated
	ReturnMessage string `json:"return_message"`
}

type callbackResponse struct {
	Data CallbackResponseData `json:"data"`
}

func NewCallbackResponse(returnCode int32, returnMessage string) CallbackResponse {
	return callbackResponse{
		Data: CallbackResponseData{
			ReturnCode:    returnCode,
			ReturnMessage: returnMessage,
		},
	}
}

func (r callbackResponse) GetData() CallbackResponseData {
	return r.Data
}

func (r callbackResponse) ToGrpcResponse() *billing_service.CallbackResponse {
	return &billing_service.CallbackResponse{
		ReturnCode:    int32(r.Data.ReturnCode),
		ReturnMessage: r.Data.ReturnMessage,
	}
}
