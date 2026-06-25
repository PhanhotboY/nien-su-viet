package adto

type GetPaymentAttemptByProviderReqDto struct {
	Provider              string `json:"provider" validate:"required"`
	ProviderTransactionId string `json:"provider_transaction_id" validate:"required"`
}
