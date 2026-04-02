package producer

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	producer3 "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/otel/tracing/producer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/producer"
	types2 "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/metadata"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/serializer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/producer/configurations"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/types"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	"emperror.dev/errors"
	"github.com/rabbitmq/amqp091-go"
)

type rabbitMQProducer struct {
	logger                  logger.Logger
	rabbitmqOptions         settings.RmqConfig
	connection              types.IConnection
	messageSerializer       serializer.MessageSerializer
	producersConfigurations map[string]*configurations.RabbitMQProducerConfiguration
	isProducedNotifications []func(message types2.IMessage)
}

func NewRabbitMQProducer(
	cfg settings.RmqConfig,
	connection types.IConnection,
	rabbitmqProducersConfiguration map[string]*configurations.RabbitMQProducerConfiguration,
	logger logger.Logger,
	eventSerializer serializer.MessageSerializer,
	isProducedNotifications ...func(message types2.IMessage),
) (producer.Producer, error) {
	p := &rabbitMQProducer{
		logger:                  logger,
		rabbitmqOptions:         cfg,
		connection:              connection,
		messageSerializer:       eventSerializer,
		producersConfigurations: rabbitmqProducersConfiguration,
	}

	p.isProducedNotifications = isProducedNotifications

	return p, nil
}

func (r *rabbitMQProducer) IsProduced(h func(message types2.IMessage)) {
	r.isProducedNotifications = append(r.isProducedNotifications, h)
}

func (r *rabbitMQProducer) PublishMessage(
	ctx context.Context,
	message types2.IMessage,
) error {
	return r.PublishMessageWithTopicName(ctx, message, "")
}

func (r *rabbitMQProducer) getProducerConfigurationByMessage(
	message types2.IMessage,
) *configurations.RabbitMQProducerConfiguration {
	messageType := utils.GetMessageBaseReflectType(message)
	return r.producersConfigurations[messageType.String()]
}

func (r *rabbitMQProducer) PublishMessageWithTopicName(
	ctx context.Context,
	message types2.IMessage,
	topicOrExchangeName string,
) error {
	producerConfiguration := r.getProducerConfigurationByMessage(message)

	if producerConfiguration == nil {
		producerConfiguration = configurations.NewDefaultRabbitMQProducerConfiguration(
			message,
		)
	}

	var exchange string
	var routingKey string

	if topicOrExchangeName != "" {
		exchange = topicOrExchangeName
	} else if producerConfiguration != nil && producerConfiguration.ExchangeOptions.Name != "" {
		exchange = producerConfiguration.ExchangeOptions.Name
	} else {
		exchange = utils.GetTopicOrExchangeName(message)
	}

	if producerConfiguration != nil && producerConfiguration.RoutingKey != "" {
		routingKey = producerConfiguration.RoutingKey
	} else {
		routingKey = message.GetRoutingKey()
	}

	producerOptions := &producer3.ProducerTracingOptions{
		MessagingSystem: "rabbitmq",
		DestinationKind: "exchange",
		Destination:     exchange,
		OtherAttributes: []attribute.KeyValue{
			semconv.MessagingRabbitmqDestinationRoutingKey(routingKey),
		},
	}

	serializedObj, err := r.messageSerializer.Serialize(message)
	if err != nil {
		return err
	}

	ctx, beforeProduceSpan := producer3.StartProducerSpan(
		ctx,
		message,
		&metadata.Metadata{},
		string(serializedObj.Data),
		producerOptions,
	)

	// https://github.com/rabbitmq/rabbitmq-tutorials/blob/master/go/publisher_confirms.go
	if r.connection == nil {
		return producer3.FinishProducerSpan(
			beforeProduceSpan,
			errors.New("connection is nil"),
		)
	}

	if r.connection.IsClosed() {
		return producer3.FinishProducerSpan(
			beforeProduceSpan,
			errors.New("connection is closed, wait for connection alive"),
		)
	}

	// create a unique channel on the connection and in the end close the channel
	channel, err := r.connection.Channel()
	// if err != nil {
	// 	return producer3.FinishProducerSpan(beforeProduceSpan, err)
	// }
	defer channel.Close()

	err = r.ensureExchange(producerConfiguration, channel, exchange)
	// if err != nil {
	// 	return producer3.FinishProducerSpan(beforeProduceSpan, err)
	// }

	if err := channel.Confirm(false); err != nil {
		return producer3.FinishProducerSpan(beforeProduceSpan, err)
	}

	confirms := make(chan amqp091.Confirmation)
	channel.NotifyPublish(confirms)

	props := amqp091.Publishing{
		MessageId:       message.GetMessageId(),
		Timestamp:       time.Now(),
		ContentType:     serializedObj.ContentType,
		Body:            serializedObj.Data,
		DeliveryMode:    producerConfiguration.DeliveryMode,
		Expiration:      producerConfiguration.Expiration,
		AppId:           producerConfiguration.AppId,
		Priority:        producerConfiguration.Priority,
		ReplyTo:         producerConfiguration.ReplyTo,
		ContentEncoding: producerConfiguration.ContentEncoding,
	}

	err = channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		true,
		false,
		props,
	)
	if err != nil {
		return producer3.FinishProducerSpan(beforeProduceSpan, err)
	}

	if confirmed := <-confirms; !confirmed.Ack {
		return producer3.FinishProducerSpan(
			beforeProduceSpan,
			errors.New("ack not confirmed"),
		)
	}

	if len(r.isProducedNotifications) > 0 {
		for _, notification := range r.isProducedNotifications {
			if notification != nil {
				notification(message)
			}
		}
	}

	return producer3.FinishProducerSpan(beforeProduceSpan, err)
}

func (r *rabbitMQProducer) ensureExchange(
	producersConfigurations *configurations.RabbitMQProducerConfiguration,
	channel *amqp091.Channel,
	exchangeName string,
) error {
	err := channel.ExchangeDeclare(
		exchangeName,
		string(producersConfigurations.ExchangeOptions.Type),
		producersConfigurations.ExchangeOptions.Durable,
		producersConfigurations.ExchangeOptions.AutoDelete,
		false,
		false,
		producersConfigurations.ExchangeOptions.Args,
	)
	if err != nil {
		return err
	}

	return nil
}
