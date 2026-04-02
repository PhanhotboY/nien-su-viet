package configurations

import (
	consumerConfigurations "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/consumer/configurations"
	producerConfigurations "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/producer/configurations"
)

type RabbitMQConfiguration struct {
	ProducersConfigurations []*producerConfigurations.RabbitMQProducerConfiguration
	ConsumersConfigurations []*consumerConfigurations.RabbitMQConsumerConfiguration
}
