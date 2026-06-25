package subscriptions

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
	crepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/infrastructure/cache"
	prepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/infrastructure/persistence"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createSubscription "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription"
	updateSubscriptionStatus "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/updateSubscriptionStatus"
)

var Module = fx.Module(
	"subscriptionsModule",

	// Provide models for DB migration
	fx.Provide(dbcontracts.NewDbModelParam(0, &entity.Subscription{})),

	fx.Provide(
		// Outbound Infrastructure
		prepo.NewSubscriptionDbRepo,
		crepo.NewSubscriptionCacheRepo,

		// Application Query

		// Application Command
		createSubscription.NewCreateSubscriptionHandler,
		updateSubscriptionStatus.NewUpdateSubscriptionStatusHandler,
	),

	// Inbound Infrastructure
)
