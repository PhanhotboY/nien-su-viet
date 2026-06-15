package app

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/config"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/infrastructure"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/webhook"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	// configure dependencies
	appBuilder := fxapp.NewApplicationBuilder(config.ConfigType())

	// provide infrastructure dependencies, e.g., database, cache, etc.
	appBuilder.ProvideModule(infrastructure.Module)

	appBuilder.ProvideModule(subscriptions.Module)
	appBuilder.ProvideModule(purchases.Module)
	appBuilder.ProvideModule(payment_attempts.Module)
	appBuilder.ProvideModule(payment_transactions.Module)
	appBuilder.ProvideModule(outbox_events.Module)
	appBuilder.ProvideModule(inbox_events.Module)
	appBuilder.ProvideModule(processed_events.Module)
	appBuilder.ProvideModule(plans.Module)
	appBuilder.ProvideModule(webhook.Module)

	app := appBuilder.Build()

	app.Logger().Info("Starting posts_service application")
	app.Run()
}
