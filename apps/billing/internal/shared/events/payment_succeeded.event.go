package event

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
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

func (e *EmbedData) String() string {
	jsonData, err := jsonUtils.MarshalToJsonString(e)
	if err != nil {
		return "{}"
	}
	return jsonData
}

type PaymentSucceededEventData struct {
	Provider       string    `json:"provider" validate:"required"`
	AppID          int       `json:"app_id" validate:"required"`
	AppTransID     string    `json:"app_trans_id" validate:"required"`
	AppTime        int64     `json:"app_time" validate:"required"`
	Amount         int64     `json:"amount" validate:"required"`
	AppUser        string    `json:"app_user" validate:"required"`
	EmbedData      EmbedData `json:"embed_data" validate:"required"`
	Item           []Item    `json:"item" validate:"required"`
	UseFeeAmount   int64     `json:"use_fee_amount"`
	DiscountAmount int64     `json:"discount_amount"`
}

type PaymentSucceededEvent interface {
	MessageEvent[*PaymentSucceededEventData, *PaymentSucceededEventData]
}

type paymentSucceededEvent struct {
	*types.Message
}

func NewPaymentSucceededEvent(msg types.IMessage) PaymentSucceededEvent {
	if msg == nil {
		msg = types.NewMessage(uuid.NewString(), []byte("{}"))
	}
	if m, ok := msg.(*paymentSucceededEvent); ok {
		msg = m.Message
	}

	msg.SetPattern(utils.GetMessageName(subscriptionCreatedEvent{}))
	return &paymentSucceededEvent{
		Message: msg.(*types.Message),
	}
}

func (e *paymentSucceededEvent) SetRawData(data string) error {
	if err := dtoUtil.ValidateStruct(data, &PaymentSucceededEventData{}); err != nil {
		return err
	}
	e.Data = json.RawMessage(data)
	return nil
}

func (e *paymentSucceededEvent) SetData(data *PaymentSucceededEventData) error {
	jsonData, err := jsonUtils.MarshalToJsonString(data)
	if err != nil {
		return err
	}
	e.Data = json.RawMessage(jsonData)
	return nil
}

func (e *paymentSucceededEvent) ParseData() (*PaymentSucceededEventData, error) {
	if e.Data == nil {
		return nil, fmt.Errorf("No data")
	}
	data := new(PaymentSucceededEventData)
	if err := jsonUtils.UnmarshalJson(string(e.Data), data); err != nil {
		return nil, err
	}

	return data, nil
}
