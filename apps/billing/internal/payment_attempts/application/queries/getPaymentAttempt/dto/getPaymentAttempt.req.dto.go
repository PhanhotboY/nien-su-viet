package adto

type GetPaymentAttemptReqDto struct {
	PaymentAttemptId string `json:"id" validate:"required,uuid"`
}
