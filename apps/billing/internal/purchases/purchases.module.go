package purchases

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/persistence"
	tgrpc "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/transport/grpc"

	createPurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase/commands"
	updatePurchaseStatusCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchaseStatus/commands"

	getPurchaseQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase/queries"
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
		prepo.NewPurchaseDbRepo,
		crepo.NewPurchaseCacheRepo,

		// Application Query
		getPurchaseQuery.NewGetPurchaseHandler,
		// Application Command
		createPurchaseCmd.NewCreatePurchaseHandler,
		updatePurchaseStatusCmd.NewUpdatePurchaseStatusHandler,
	),

	// Inbound Infrastructure
	tgrpc.Module,
)
