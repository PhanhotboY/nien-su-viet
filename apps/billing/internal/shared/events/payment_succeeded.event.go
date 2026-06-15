package event

import (
	"encoding/json"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
)

type Item struct {
	ItemID    string `json:"itemid"`
	ItemName  string `json:"itemname"`
	ItemPrice int64  `json:"itemprice"`
}

type EmbedData struct {
	PurchaseID  string `json:"purchaseid"`
	PlanID      string `json:"planid"`
	RedirectURL string `json:"redirecturl"`
}

type PaymentSucceededEventData struct {
	AppID          int       `json:"app_id"`
	AppTransID     string    `json:"app_trans_id"`
	AppTime        int64     `json:"app_time"`
	Amount         int64     `json:"amount"`
	AppUser        string    `json:"app_user"`
	EmbedData      EmbedData `json:"embed_data"`
	Item           Item      `json:"item"`
	UseFeeAmount   int64     `json:"use_fee_amount"`
	DiscountAmount int64     `json:"discount_amount"`
}

type PaymentSucceededEventDataInput struct {
	AppID          int    `json:"app_id"`
	AppTransID     string `json:"app_trans_id"`
	AppTime        int64  `json:"app_time"`
	Amount         int64  `json:"amount"`
	AppUser        string `json:"app_user"`
	EmbedData      string `json:"embed_data"`
	Item           string `json:"item"`
	UseFeeAmount   int64  `json:"use_fee_amount"`
	DiscountAmount int64  `json:"discount_amount"`
}

type PaymentSucceededEvent interface {
	MessageEvent[*PaymentSucceededEventData]
}

type paymentSucceededEvent struct {
	*types.Message
}

func NewPaymentSucceededEvent(msg types.IMessage) PaymentSucceededEvent {
	return &paymentSucceededEvent{
		Message: msg.(*types.Message),
	}
}

func (e *paymentSucceededEvent) SetData(data string) error {
	if err := dtoUtil.ValidateStruct(data, &PaymentSucceededEventDataInput{}); err != nil {
		return err
	}
	e.Data = json.RawMessage(data)
	return nil
}

func (e *paymentSucceededEvent) ParseData() (*PaymentSucceededEventData, error) {
	data := new(PaymentSucceededEventData)
	if err := jsonUtils.UnmarshalJson(string(e.Data), data); err != nil {
		return nil, err
	}

	return data, nil
}
