package payment_attempts

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/infrastructure/persistence"
	tgrpc "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/infrastructure/transport/grpc"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createPaymentAttemptCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt"
	updatePaymentAttemptStatusCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus"
	getPaymentAttemptQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt"
	getPaymentAttemptByProviderQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttemptByProvider"
)

var Module = fx.Module(
	"paymentAttemptsModule",

	// Provide models for DB migration
	// Ensure payment attempts are migrated after purchases
	fx.Provide(dbcontracts.NewDbModelParam(20, &entity.PaymentAttempt{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewPaymentAttemptDbRepo,
		crepo.NewPaymentAttemptCacheRepo,

		// Application Query
		getPaymentAttemptQuery.NewGetPaymentAttemptHandler,
		getPaymentAttemptByProviderQuery.NewGetPaymentAttemptByProviderHandler,
		// Application Command
		createPaymentAttemptCmd.NewCreatePaymentAttemptHandler,
		updatePaymentAttemptStatusCmd.NewUpdatePaymentAttemptStatusHandler,
	),

	// Inbound Infrastructure
	tgrpc.Module,
)
