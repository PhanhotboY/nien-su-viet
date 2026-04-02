package producer

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/producer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	serializer "github.com/phanhotboy/nien-su-viet/libs/pkg/core/serializer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	producerConfigurations "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/producer/configurations"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/producer/producercontracts"
	types2 "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"
)

type producerFactory struct {
	connection      types2.IConnection
	logger          logger.Logger
	eventSerializer serializer.MessageSerializer
	rabbitmqOptions settings.RmqConfig
}

func NewProducerFactory(
	s settings.Config,
	connection types2.IConnection,
	eventSerializer serializer.MessageSerializer,
	l logger.Logger,
) producercontracts.ProducerFactory {
	return &producerFactory{
		rabbitmqOptions: s.Rmq,
		logger:          l,
		connection:      connection,
		eventSerializer: eventSerializer,
	}
}

func (p *producerFactory) CreateProducer(
	rabbitmqProducersConfiguration map[string]*producerConfigurations.RabbitMQProducerConfiguration,
	isProducedNotifications ...func(message types.IMessage),
) (producer.Producer, error) {
	return NewRabbitMQProducer(
		p.rabbitmqOptions,
		p.connection,
		rabbitmqProducersConfiguration,
		p.logger,
		p.eventSerializer,
		isProducedNotifications...)
}
