package adto

type ListPaymentTransactionsByAttemptReqDto struct {
	PaymentAttemptId string `json:"payment_attempt_id" validate:"required,uuid"`
	Query            string `json:"query" validate:"omitempty"`
}
