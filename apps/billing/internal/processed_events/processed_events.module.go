package processed_events

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/entity"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/infrastructure/persistence"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createProcessedEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent"
)

var Module = fx.Module(
	"processedEventsModule",

	// Provide models for DB migration
	fx.Provide(dbcontracts.NewDbModelParam(0, &entity.ProcessedEvent{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewProcessedEventDbRepo,

		// Application Query
		// Application Command
		createProcessedEvent.NewCreateProcessedEventHandler,
	),

	// Inbound Infrastructure
)
