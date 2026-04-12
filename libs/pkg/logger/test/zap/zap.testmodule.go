package zaptest

import (
	"testing"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	"go.uber.org/fx"
)

var ModuleFunc = func(t *testing.T, l logger.Logger) fx.Option {
	testZapLogger := NewZapTestLogger(t, l)

	return fx.Module(
		"zapModule",

		fx.Supply(fx.Annotate(testZapLogger, fx.As(new(testlogger.TestLogger)))),
		fx.Supply(fx.Annotate(testZapLogger, fx.As(new(ZapTestLogger)))),
	)
}
