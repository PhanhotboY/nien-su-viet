package sdto

import "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"

type MoneyDto struct {
	Amount   int64  `json:"amount" validate:"required,min=10000"`
	Currency string `json:"currency" validate:"omitempty,len=3"`
}

func NewMoneyDto(amount int64, currency string) *MoneyDto {
	if currency == "" {
		currency = "VND"
	}
	return &MoneyDto{
		Amount:   amount,
		Currency: currency,
	}
}

func (m *MoneyDto) ToGrpcMoney() *billing_service.Money {
	return &billing_service.Money{
		Amount:   m.Amount,
		Currency: m.Currency,
	}
}
