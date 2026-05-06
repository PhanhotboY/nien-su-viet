package consumer

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"

	createPostCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/createPost/v1/commands"
	createPostEvent "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/createPost/v1/events"
	deletePostCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/deletePost/v1/commands"
	deletePostEvent "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/deletePost/v1/events"
	updatePostCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/updatePost/v1/commands"
	updatePostEvent "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/updatePost/v1/events"
)

// SetupPostConsumers configures post-related RabbitMQ consumers with custom routing keys and events exchange
func SetupPostConsumers(
	s settings.Config,
	b bus.RabbitmqBus,
	logger logger.Logger,

	createPostHandler createPostCommand.ICreatePostHandler,
	updatePostHandler updatePostCommand.IUpdatePostHandler,
	deletePostHandler deletePostCommand.IDeletePostHandler,
) error {
	b.ConnectConsumerHandler(createPostEvent.NewPostCreatedEvent(), createPostHandler)
	b.ConnectConsumerHandler(updatePostEvent.NewPostUpdatedEvent(), updatePostHandler)
	b.ConnectConsumerHandler(deletePostEvent.NewPostDeletedEvent(), deletePostHandler)

	return nil
}
