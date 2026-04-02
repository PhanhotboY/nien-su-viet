package consumer

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	serializer "github.com/phanhotboy/nien-su-viet/libs/pkg/core/serializer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	consumerConfigurations "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/consumer/configurations"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/consumer/consumercontracts"
	types2 "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"
)

type consumerFactory struct {
	connection      types2.IConnection
	eventSerializer serializer.MessageSerializer
	logger          logger.Logger
	rabbitmqOptions settings.RmqConfig
}

func NewConsumerFactory(
	s settings.Config,
	connection types2.IConnection,
	eventSerializer serializer.MessageSerializer,
	l logger.Logger,
) consumercontracts.ConsumerFactory {
	return &consumerFactory{
		rabbitmqOptions: s.Rmq,
		logger:          l,
		eventSerializer: eventSerializer,
		connection:      connection,
	}
}

func (c *consumerFactory) CreateConsumer(
	consumerConfiguration *consumerConfigurations.RabbitMQConsumerConfiguration,
	isConsumedNotifications ...func(message types.IMessage),
) (consumer.Consumer, error) {
	return NewRabbitMQConsumer(
		c.rabbitmqOptions,
		c.connection,
		consumerConfiguration,
		c.eventSerializer,
		c.logger,
		isConsumedNotifications...)
}

func (c *consumerFactory) Connection() types2.IConnection {
	return c.connection
}
