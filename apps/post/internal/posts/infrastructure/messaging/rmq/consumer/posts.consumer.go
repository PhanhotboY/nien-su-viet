package consumer

import (
	events "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/events/intergration"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/queries"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"
	consumerConfigurations "github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/consumer/configurations"
)

// SetupPostConsumers configures post-related RabbitMQ consumers with custom routing keys and events exchange
func SetupPostConsumers(
	s settings.Config,
	b bus.RabbitmqBus,
	logger logger.Logger,
) error {

	// Configure consumer for posts.createdrouting key using the events exchange
	err := b.ConnectRabbitMQConsumer(
		events.NewGetPublishedPostsEvent(),
		func(builder consumerConfigurations.RabbitMQConsumerConfigurationBuilder) {
			builder.
				WithExchangeName("events").
				WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							queries.NewGetPublishedPostsHandler(logger),
						)
					},
				)
		})

	return err
}
