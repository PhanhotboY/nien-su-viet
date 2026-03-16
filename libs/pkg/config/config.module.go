package config

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"

	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module(
	"configModule",
	fx.Provide(func() environment.Environment {
		return environment.ConfigAppEnv()
	}),
)

var ModuleFunc = func(e environment.Environment) fx.Option {
	return fx.Module(
		"configModule",
		fx.Provide(func() environment.Environment {
			return environment.ConfigAppEnv(e)
		}),
	)
}
