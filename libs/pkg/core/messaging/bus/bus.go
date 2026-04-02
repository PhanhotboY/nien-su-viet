package bus

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/producer"
)

type Bus interface {
	producer.Producer
	consumer.ConsumerConnector
	consumer.BusControl
}
