package consumer

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"

	createEventCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/createEvent/v1/commands"
	createEventEvent "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/createEvent/v1/events"
	deleteEventCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/deleteEvent/v1/commands"
	deleteEventEvent "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/deleteEvent/v1/events"
	updateEventCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/updateEvent/v1/commands"
	updateEventEvent "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/updateEvent/v1/events"
)

// SetupHistoricalEventsConsumers configures historical-event-related RabbitMQ consumers with custom routing keys and events exchange
func SetupHistoricalEventsConsumers(
	s settings.Config,
	b bus.RabbitmqBus,
	logger logger.Logger,
	createEventHandler createEventCommand.ICreateEventHandler,
	updateEventHandler updateEventCommand.IUpdateEventHandler,
	deleteEventHandler deleteEventCommand.IDeleteEventHandler,
) error {

	b.ConnectConsumerHandler(
		createEventEvent.NewHistoricalEventCreatedEvent(),
		createEventHandler,
	)

	b.ConnectConsumerHandler(
		deleteEventEvent.NewHistoricalEventDeletedEvent(),
		deleteEventHandler,
	)

	b.ConnectConsumerHandler(
		updateEventEvent.NewHistoricalEventUpdatedEvent(),
		updateEventHandler,
	)

	return nil
}
