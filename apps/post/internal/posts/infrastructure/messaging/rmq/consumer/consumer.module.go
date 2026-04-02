package consumer

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"postsRmqConsumerModule",
	fx.Invoke(SetupPostConsumers),
)
