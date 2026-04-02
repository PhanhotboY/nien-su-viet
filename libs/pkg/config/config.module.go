package config

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"

	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module(
	"configModule",
	fx.Provide(settings.LoadConfig),
	fx.Supply(settings.GetEnv),
)
