package trmq

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
)

func NewPurchasesConsumer(
	b bus.RabbitmqBus,

	paymentSucceededEventHandler PaymentSucceededEventHandler,
) error {
	b.ConnectConsumerHandler(
		event.NewPaymentSucceededEvent(nil),
		paymentSucceededEventHandler,
	)

	return nil
}
