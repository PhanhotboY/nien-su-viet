package hevents

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/infrastructure/cache"
	rmqConsumer "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/infrastructure/messaging/rmq/consumer"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/infrastructure/persistence"

	getAllEventsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getAllEvents/v1/queries"
	getEventQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getEvent/v1/queries"

	createEventCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/createEvent/v1/commands"
	deleteEventCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/deleteEvent/v1/commands"
	updateEventCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/command/updateEvent/v1/commands"

	// rmqProvider	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/infrastructure/messaging/rmq"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/infrastructure/transport/grpc"
)

var Module = fx.Module(
	"historicalEventsModule",

	fx.Provide(
		// Outbound Infrastructure
		cache.NewHistoricalEventCacheRepository,
		persistence.NewHistoricalEventEsRepository,

		// Application Query
		getAllEventsQuery.NewGetAllEventsHandler,
		getEventQuery.NewGetEventHandler,
		createEventCommand.NewCreateEventHandler,
		updateEventCommand.NewUpdateEventHandler,
		deleteEventCommand.NewDeleteEventHandler,
	),

	// Inbound Infrastructure
	grpc.Module,
	rmqConsumer.Module,
)
