package subscriptions

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
)

var Module = fx.Module(
	"subscriptionsModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() any {
			return &entity.Subscription{}
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
