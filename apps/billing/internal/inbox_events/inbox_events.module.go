package inbox_events

import (
	"go.uber.org/fx"

	createInboxEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/createInboxEvent"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/infrastructure/persistence"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

var Module = fx.Module(
	"inboxEventsModule",

	// Provide models for DB migration
	fx.Provide(dbcontracts.NewDbModelParam(0, &entity.InboxEvent{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewInBoxEventDbRepo,

		// Application Query
		// Application Command
		createInboxEvent.NewCreateInboxEventCmdHandler,
	),

	// Inbound Infrastructure
)
