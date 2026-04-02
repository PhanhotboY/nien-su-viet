package options

import "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"

type RabbitMQExchangeOptions struct {
	Name       string
	Type       types.ExchangeType
	AutoDelete bool
	Durable    bool
	Args       map[string]any
}
