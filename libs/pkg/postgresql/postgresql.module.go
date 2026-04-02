package postgres

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module("postgresPgxModule",
	fx.Provide(NewDb),
	fx.Invoke(registerHooks),
)

func registerHooks(lc fx.Lifecycle, pgClient *gorm.DB, logger logger.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			db, err := pgClient.DB()
			if err != nil {
				logger.Error("Failed to get database instance", "error", err)
				return err
			}
			db.Close()
			logger.Info("Postgres connection closed gracefully")

			return nil
		},
	})
}
