package payment_attempts

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/infrastructure/persistence"
	tgrpc "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/infrastructure/transport/grpc"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createPaymentAttemptCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt"
	getPaymentAttemptQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt"
)

var Module = fx.Module(
	"paymentAttemptsModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() dbcontracts.DbModelParam {
			return dbcontracts.DbModelParam{
				Order: 20, // Ensure payment attempts are migrated after subscriptions and purchases
				Model: &entity.PaymentAttempt{},
			}
		},
		fx.ResultTags(`group:"db_models"`),
	)),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewPaymentAttemptDbRepo,
		crepo.NewPaymentAttemptCacheRepo,

		// Application Query
		getPaymentAttemptQuery.NewGetPaymentAttemptHandler,
		// Application Command
		createPaymentAttemptCmd.NewCreatePaymentAttemptHandler,
	),

	// Inbound Infrastructure
	tgrpc.Module,
)
