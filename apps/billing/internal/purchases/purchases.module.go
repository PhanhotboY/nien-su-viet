package purchases

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/persistence"
	tgrpc "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/transport/grpc"
	trmq "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/transport/rabbitmq"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createPurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase"
	updatePurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase"
	updatePurchaseStatusCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchaseStatus"

	getPurchaseQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase"
)

var Module = fx.Module(
	"purchasesModule",

	// Provide models for DB migration
	// Ensure purchases are migrated after subscriptions and plans
	fx.Provide(dbcontracts.NewDbModelParam(10, &entity.Purchase{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewPurchaseDbRepo,
		crepo.NewPurchaseCacheRepo,

		// Application Query
		getPurchaseQuery.NewGetPurchaseHandler,
		// Application Command
		createPurchaseCmd.NewCreatePurchaseHandler,
		updatePurchaseCmd.NewUpdatePurchaseHandler,
		updatePurchaseStatusCmd.NewUpdatePurchaseStatusHandler,
	),

	// Inbound Infrastructure
	tgrpc.Module,
	trmq.Module,
)
