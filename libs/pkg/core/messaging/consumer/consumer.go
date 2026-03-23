package consumer

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
)

type Consumer interface {
	Start(ctx context.Context) error
	Stop() error
	ConnectHandler(handler ConsumerHandler)
	GetName() string
}

type ConsumerHandler interface {
	Handle(ctx context.Context, consumeContext types.MessageConsumeContext) error
}

type ConsumerConnector interface {
	ConnectConsumer(message types.IMessage, consumer Consumer) error
	ConnectConsumerHandler(message types.IMessage, consumerHandler ConsumerHandler) error
}

type ConsumerOptions struct {
	ExitOnError bool
	ConsumerId  string
}

type ConsumerHandlerConfiguration struct {
	Handlers []ConsumerHandler
}

type ConsumerHandlerConfigurationBuilder interface {
	AddHandler(handler ConsumerHandler) ConsumerHandlerConfigurationBuilder
	Build() *ConsumerHandlerConfiguration
}

type ConsumerHandlerConfigurationBuilderFunc func(builder ConsumerHandlerConfigurationBuilder)

type consumerHandlerConfigurationBuilder struct {
	handlers []ConsumerHandler
}

func NewConsumerHandlersConfigurationBuilder() ConsumerHandlerConfigurationBuilder {
	return &consumerHandlerConfigurationBuilder{handlers: []ConsumerHandler{}}
}

func (b *consumerHandlerConfigurationBuilder) AddHandler(handler ConsumerHandler) ConsumerHandlerConfigurationBuilder {
	b.handlers = append(b.handlers, handler)
	return b
}

func (b *consumerHandlerConfigurationBuilder) Build() *ConsumerHandlerConfiguration {
	return &ConsumerHandlerConfiguration{Handlers: b.handlers}
}
