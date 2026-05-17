package outbox_events

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
)

var Module = fx.Module(
	"outboxEventsModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() any {
			return &entity.OutboxEvent{}
		},
		fx.ResultTags(`group:"db_models"`),
	)),

	fx.Provide(

	// Outbound Infrastructure

	// Application Query
	// Application Command
	),

	// Inbound Infrastructure
)
