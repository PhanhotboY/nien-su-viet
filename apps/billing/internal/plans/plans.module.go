package plans

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/infrastructure/persistence"
	tgrpc "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/infrastructure/transport/grpc"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createPlanCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan"
	getPlanByIdQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById"
	listPlansQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/listPlans"
)

var Module = fx.Module(
	"plansModule",

	// Provide models for DB migration
	fx.Provide(dbcontracts.NewDbModelParam(0, &entity.Plan{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewPlanDbRepo,
		crepo.NewPlanCacheRepo,

		// Application Query
		listPlansQuery.NewListPlansHandler,
		getPlanByIdQuery.NewGetPlanByIdHandler,
		// Application Command
		createPlanCmd.NewCreatePlanHandler,
	),

	// Inbound Infrastructure
	tgrpc.Module,
)
