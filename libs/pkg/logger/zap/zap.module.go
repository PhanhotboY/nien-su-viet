package zap

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/config"

	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module("zapModule",

	// - order is not important in provide
	// - provide can have parameter and will resolve if registered
	// - execute its func only if it requested
	fx.Provide(
		config.ProvideLogConfig,
		NewZapLogger,
		fx.Annotate(
			NewZapLogger,
			fx.As(new(logger.Logger))),
	),
)

var ModuleFunc = func(l logger.Logger) fx.Option {
	return fx.Module(
		"zapModule",

		fx.Provide(config.ProvideLogConfig),
		fx.Supply(fx.Annotate(l, fx.As(new(logger.Logger)))),
		fx.Supply(fx.Annotate(l, fx.As(new(ZapLogger)))),
	)
}
