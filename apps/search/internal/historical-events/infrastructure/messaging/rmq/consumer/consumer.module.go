package consumer

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"historicalEventsRmqConsumerModule",

	fx.Invoke(SetupHistoricalEventsConsumers),
)
