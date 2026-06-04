package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type PaymentSucceededEvent struct {
	types.Message
}

func NewPaymentSucceededEvent() PaymentSucceededEvent {
	return PaymentSucceededEvent{
		Message: types.Message{},
	}
}
