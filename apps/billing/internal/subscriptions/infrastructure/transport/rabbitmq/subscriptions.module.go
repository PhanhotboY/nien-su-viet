package trmq

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"subscriptionsRmqModule",

	fx.Provide(
		NewUserAddedToOrganizationEventHandler,
	),

	fx.Invoke(NewPurchasesConsumer),
)
