package testhelper

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts"
	"github.com/phanhotboy/nien-su-viet/apps/post/test/integration/shared/infrastructure"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	zaptest "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test/zap"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/zap"

	postConfig "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/config"
)

func GetDIServices(t *testing.T, services ...any) {

	cfg := config.LoadConfig(postConfig.ConfigType())
	var logger logger.Logger
	logger = zap.NewZapLogger(cfg)

	fxtest.New(t,
		fx.Supply(fx.Annotate(
			cfg,
			fx.As(new(config.Config)),
		)),
		zap.ModuleFunc(logger),
		infrastructure.TestModuleFunc(t),

		posts.Module,

		fx.NopLogger,
		zaptest.ModuleFunc(t, logger),

		fx.Populate(services...),
	).RequireStart()
}
