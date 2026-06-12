package zalopay

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type CreateOrderRequest struct {
	AppID       int    `json:"app_id"`
	AppTransID  string `json:"app_trans_id"`
	AppTime     int64  `json:"app_time"`
	Amount      int64  `json:"amount"`
	AppUser     string `json:"app_user"`
	Item        string `json:"item"`
	EmbedData   string `json:"embed_data"`
	Description string `json:"description"`
	BankCode    string `json:"bank_code,omitempty"`
	Mac         string `json:"mac"`
	CallbackURL string `json:"callback_url"`
	Title       string `json:"json"`
}

func (r CreateOrderRequest) computeMac(key string) string {
	// hmac_input: app_id +”|”+ app_trans_id +”|”+ app_user +”|”+ amount +"|"+ app_time +”|”+ embed_data +"|"+ item
	raw := strings.Join([]string{
		strconv.Itoa(r.AppID),
		r.AppTransID,
		r.AppUser,
		strconv.FormatInt(r.Amount, 10),
		strconv.FormatInt(r.AppTime, 10),
		r.EmbedData,
		r.Item,
	}, "|")
	return HMACSHA256Hex(raw, key)
}

type CreateOrderResponse struct {
	ReturnCode       int    `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int    `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
	OrderURL         string `json:"order_url,omitempty"`
	Zptranstoken     string `json:"zptranstoken,omitempty"`
	OrderToken       string `json:"order_token,omitempty"`
	QRCode           string `json:"qr_code,omitempty"`
}

type QueryOrderRequest struct {
	AppID      int    `json:"app_id"`
	AppTransID string `json:"app_trans_id"`
	Mac        string `json:"mac"`
}

func (r QueryOrderRequest) computeMac(key string) string {
	raw := strings.Join([]string{
		strconv.Itoa(r.AppID),
		r.AppTransID,
	}, "|")
	return HMACSHA256Hex(raw, key)
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

type Item struct {
	ItemID    string `json:"itemid"`
	ItemName  string `json:"itemname"`
	ItemPrice int64  `json:"itemprice"`
}

func (i Item) String() string {
	itemJSON, _ := json.Marshal(i)
	return fmt.Sprintf("[%s]", string(itemJSON))
}

type EmbedData struct {
	PurchaseID  string `json:"purchaseid"`
	RedirectURL string `json:"redirecturl"`
}

func (e EmbedData) String() string {
	embedJSON, _ := json.Marshal(e)
	return string(embedJSON)
}
