package rmq

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/messaging/rmq/consumer"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"postsInfrastructureMessagingRmqModule",
	consumer.Module,
)
