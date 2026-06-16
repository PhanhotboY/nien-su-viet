package processed_events

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/domain/entity"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

var Module = fx.Module(
	"processedEventsModule",

	// Provide models for DB migration
	fx.Provide(dbcontracts.NewDbModelParam(0, &entity.ProcessedEvent{})),

	fx.Provide(

	// Outbound Infrastructure

	// Application Query
	// Application Command
	),

	// Inbound Infrastructure
)
