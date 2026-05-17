package payment_attempts

import (
	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
)

var Module = fx.Module(
	"paymentAttemptsModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() any {
			return &entity.PaymentAttempt{}
		},
		fx.ResultTags(`group:"db_models"`),
	)),

	fx.Provide(

	// Outbound Infrastructure

	// Application Query
	// Application Command
	),

	// Inbound Infrastructure
)
