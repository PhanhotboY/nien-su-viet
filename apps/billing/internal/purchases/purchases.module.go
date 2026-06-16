package purchases

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/persistence"
	tgrpc "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/transport/grpc"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createPurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase"
	updatePurchaseStatusCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchaseStatus"

	getPurchaseQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase"
)

var Module = fx.Module(
	"purchasesModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() dbcontracts.DbModelParam {
			return dbcontracts.DbModelParam{
				Order: 10, // Ensure purchases are migrated after subscriptions and plans
				Model: &entity.Purchase{},
			}
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
