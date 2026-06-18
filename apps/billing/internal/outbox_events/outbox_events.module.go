package outbox_events

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/infrastructure/persistence"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createOutboxEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/application/commands/createOutboxEvent"
)

var Module = fx.Module(
	"outboxEventsModule",

	// Provide models for DB migration
	fx.Provide(dbcontracts.NewDbModelParam(0, &entity.OutboxEvent{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewOutboxEventDbRepo,

		// Application Query

		// Application Command
		createOutboxEvent.NewCreateOutboxEventHandler,
	),

	// Inbound Infrastructure
)
