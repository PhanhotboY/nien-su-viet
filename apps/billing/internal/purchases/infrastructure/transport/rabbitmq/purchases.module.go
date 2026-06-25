package trmq

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"purchasesRmqModule",

	fx.Provide(
		NewPaymentSucceededEventHandler,
	),

	fx.Invoke(NewPurchasesConsumer),
)
