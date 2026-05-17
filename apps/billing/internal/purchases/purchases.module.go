package purchases

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
)

var Module = fx.Module(
	"purchasesModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() any {
			return &entity.Purchase{}
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
