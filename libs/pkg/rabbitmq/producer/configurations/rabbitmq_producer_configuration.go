package configurations

import (
	"reflect"

	types2 "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/producer/options"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"
)

type RabbitMQProducerConfiguration struct {
	ProducerMessageType reflect.Type
	ExchangeOptions     *options.RabbitMQExchangeOptions
	RoutingKey          string
	DeliveryMode        uint8
	Priority            uint8
	AppId               string
	Expiration          string
	ReplyTo             string
	ContentEncoding     string
}

func NewDefaultRabbitMQProducerConfiguration(
	messageType types2.IMessage,
) *RabbitMQProducerConfiguration {
	return &RabbitMQProducerConfiguration{
		ExchangeOptions: &options.RabbitMQExchangeOptions{
			Durable: true,
			Type:    types.ExchangeTopic,
			Name:    utils.GetTopicOrExchangeName(messageType),
		},
		DeliveryMode:        2,
		RoutingKey:          utils.GetRoutingKey(messageType),
		ProducerMessageType: utils.GetMessageBaseReflectType(messageType),
	}
}
