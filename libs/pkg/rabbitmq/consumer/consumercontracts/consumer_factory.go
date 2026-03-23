package consumercontracts

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	types2 "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/consumer/configurations"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"
)

type ConsumerFactory interface {
	CreateConsumer(
		consumerConfiguration *configurations.RabbitMQConsumerConfiguration,
		isConsumedNotifications ...func(message types2.IMessage),
	) (consumer.Consumer, error)

	Connection() types.IConnection
}
