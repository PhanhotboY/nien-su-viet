package zalopay

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CreateOrderRequest struct {
	AppID       string `json:"app_id"`
	AppTransID  string `json:"app_trans_id"`
	AppTime     int64  `json:"app_time"`
	Amount      int64  `json:"amount"`
	AppUser     string `json:"app_user"`
	Item        string `json:"item"`
	EmbedData   string `json:"embed_data"`
	Description string `json:"description"`
	BankCode    string `json:"bank_code,omitempty"`
	Mac         string `json:"mac"`
}

type CreateOrderResponse struct {
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_message"`
	OrderURL      string `json:"order_url,omitempty"`
	Zptranstoken  string `json:"zptranstoken,omitempty"`
	AppTransID    string `json:"apptransid,omitempty"`
	SubReturnCode int    `json:"sub_return_code,omitempty"`
}

type QueryOrderRequest struct {
	AppID      string `json:"app_id"`
	AppTransID string `json:"app_trans_id"`
	Mac        string `json:"mac"`
}

type QueryOrderResponse struct {
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_message"`
	Amount        int64  `json:"amount,omitempty"`
	Discount      int64  `json:"discount,omitempty"`
	RefundAmount  int64  `json:"refund_amount,omitempty"`
	Status        int    `json:"status,omitempty"`
	AppTransID    string `json:"apptransid,omitempty"`
	UserFeeAmount int64  `json:"user_fee_amount,omitempty"`
}

type CallbackPayload struct {
	Data string `json:"data"`
	Mac  string `json:"mac"`
}

type CallbackData struct {
	AppID      string `json:"app_id"`
	AppTransID string `json:"app_trans_id"`
	AppTime    int64  `json:"app_time"`
	Amount     int64  `json:"amount"`
	AppUser    string `json:"app_user"`
	EmbedData  string `json:"embed_data"`
	Item       string `json:"item"`
	ZpTransID  string `json:"zp_trans_id"`
	BankCode   string `json:"bank_code"`
	Status     int    `json:"status"`
}

func (r CreateOrderRequest) computeMac(key string) string {
	raw := strings.Join([]string{
		r.AppID,
		r.AppTransID,
		strconv.FormatInt(r.AppTime, 10),
		strconv.FormatInt(r.Amount, 10),
		r.AppUser,
		r.Item,
		r.EmbedData,
	}, "|")
	return HMACSHA256Hex(raw, key)
}

func (r QueryOrderRequest) computeMac(key string) string {
	raw := strings.Join([]string{
		r.AppID,
		r.AppTransID,
	}, "|")
	return HMACSHA256Hex(raw, key)
}

func ParseCallback(body []byte) (*CallbackData, error) {
	var payload CallbackPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("zalopay: decode callback wrapper: %w", err)
	}
	if payload.Data == "" {
		return nil, fmt.Errorf("zalopay: missing callback data")
	}

	var data CallbackData
	if err := json.Unmarshal([]byte(payload.Data), &data); err != nil {
		return nil, fmt.Errorf("zalopay: decode callback data: %w", err)
	}
	return &data, nil
}

func GenerateAppTransID(prefix string, t time.Time) string {
	return fmt.Sprintf("%s%d", prefix, t.Unix())
}
