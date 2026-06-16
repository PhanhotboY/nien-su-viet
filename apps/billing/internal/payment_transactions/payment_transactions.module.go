package payment_transactions

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/infrastructure/persistence"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createPT "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction"
	listPTByAttempt "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/queries/listPaymentTransactionsByAttempt"
)

var Module = fx.Module(
	"paymentTransactionsModule",

	// Provide models for DB migration
	// Ensure payment transactions are migrated after payment attempts
	fx.Provide(dbcontracts.NewDbModelParam(30, &entity.PaymentTransaction{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewPaymentTransactionDbRepo,
		crepo.NewPaymentTransactionCacheRepo,

		// Application Query
		listPTByAttempt.NewListPaymentTransactionsByAttemptHandler,
		// Application Command
		createPT.NewCreatePaymentTransactionHandler,
	),

	// Inbound Infrastructure
)
