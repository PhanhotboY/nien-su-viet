package consumer

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"
)

// SetupPostConsumers configures post-related RabbitMQ consumers with custom routing keys and events exchange
func SetupPostConsumers(
	s config.Config,
	b bus.RabbitmqBus,
	logger logger.Logger,
) error {
	return nil
}
