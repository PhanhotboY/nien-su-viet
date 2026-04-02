package contracts

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.uber.org/fx"
)

type Application interface {
	Container
	RegisterHook(function interface{})
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
	Logger() logger.Logger
	Settings() settings.Config
}
