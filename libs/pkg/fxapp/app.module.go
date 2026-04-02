package fxapp

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/external/fxlog"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/zap"

	"go.uber.org/fx"
)

func CreateAppModule(
	app *application,
) *fx.App {
	var opts []fx.Option

	opts = append(opts, fx.Provide(app.provides...))

	opts = append(opts, fx.Decorate(app.decorates...))

	opts = append(opts, fx.Invoke(app.invokes...))

	app.options = append(app.options, opts...)

	AppModule := fx.Module("appModule",
		app.options...,
	)

	var logModule = zap.ModuleFunc(app.logger)
	duration := 30 * time.Second

	// build phase of container will do in this stage, containing provides and invokes but app not started yet and will be started in the future with `fxApp.Register`
	fxApp := fx.New(
		fx.StartTimeout(duration),
		config.Module,
		logModule,
		fxlog.FxLogger,
		fx.ErrorHook(NewFxErrorHandler(app.logger)),
		AppModule,
	)

	return fxApp
}
