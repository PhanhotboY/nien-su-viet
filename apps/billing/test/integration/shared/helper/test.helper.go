package testhelper

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	"github.com/phanhotboy/nien-su-viet/apps/billing/test/integration/shared/infrastructure"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	zaptest "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test/zap"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/zap"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases"
	billingConfig "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/config"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/webhook"
)

func GetDIServices(t *testing.T, services ...any) {

	cfg := config.LoadConfig(billingConfig.ConfigType())
	var logger logger.Logger
	logger = zap.NewZapLogger(cfg)

	fxtest.New(t,
		fx.Supply(fx.Annotate(
			cfg,
			fx.As(new(config.Config)),
		)),
		zap.ModuleFunc(logger),
		infrastructure.TestModuleFunc(t),

		plans.Module,
		purchases.Module,
		payment_attempts.Module,
		payment_transactions.Module,
		subscriptions.Module,
		outbox_events.Module,
		inbox_events.Module,
		processed_events.Module,
		webhook.Module,

		fx.NopLogger,
		zaptest.ModuleFunc(t, logger),

		fx.Populate(services...),
	).RequireStart()
}
