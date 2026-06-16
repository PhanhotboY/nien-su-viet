package postgres

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	dbhelper "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/helper"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

// https://aiven.io/blog/aiven-for-postgresql-for-your-go-application

const (
	maxConn         = 50
	maxConnIdleTime = 1 * time.Minute
	maxConnLifetime = 3 * time.Minute
	minConns        = 10
)

type DBParams struct {
	fx.In

	Models []dbcontracts.DbModelParam `group:"db_models"`
}

type txContextDb struct {
	logger logger.Logger
	db     *gorm.DB
}

// NewDb func for connection to PostgreSQL database.
func NewDb(c config.Config, logger logger.Logger, params DBParams) (dbcontracts.TxContextDb, error) {
	cfg := c.GetPostgresqlOptions()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port)

	pg, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gormLogger.New(logger, gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.LogLevel(cfg.LogLevel),
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	pg.Use(tracing.NewPlugin(
		tracing.WithoutQueryVariables(),
	))
	sqlDB, err := pg.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(minConns)
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(maxConnLifetime)
	sqlDB.SetConnMaxIdleTime(maxConnIdleTime)

	// Safe auto migration - will only modify schema if needed
	sort.Slice(params.Models, func(i, j int) bool {
		return params.Models[i].Order < params.Models[j].Order
	})
	entities := make([]any, len(params.Models))
	for i, modelParam := range params.Models {
		entities[i] = modelParam.Model
	}
	if err := pg.AutoMigrate(entities...); err != nil {
		// Check if error is about existing relations
		if strings.Contains(err.Error(), "already exists") {
			fmt.Println("Database tables already exist, skipping migration")
		} else {
			return nil, fmt.Errorf("failed to auto migrate: %w", err)
		}
	} else {
		fmt.Println("Database migrations completed successfully!")
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Database connection established successfully!")
	return txContextDb{db: pg, logger: logger}, nil
}

func (t txContextDb) WithTx(ctx context.Context) (dbcontracts.TxContextDb, error) {
	tx, err := dbhelper.GetTxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return txContextDb{db: tx, logger: t.logger}, nil
}

func (t txContextDb) WithTxIfExists(ctx context.Context) dbcontracts.TxContextDb {
	tx := dbhelper.GetTxFromContextIfExists(ctx)
	if tx == nil {
		return t
	}

	return txContextDb{db: tx, logger: t.logger}
}

func (t txContextDb) RunInTx(ctx context.Context, action dbcontracts.ActionFunc) error {
	// https://gorm.io/docs/transactions.html#Transaction
	tx := t.DB().WithContext(ctx).Begin()

	t.logger.Info("beginning database transaction")

	gormContext := dbhelper.SetTxToContext(ctx, tx)
	ctx = gormContext

	defer func() {
		if r := recover(); r != nil {
			tx.WithContext(ctx).Rollback()

			if err, _ := r.(error); err != nil {
				t.logger.Errorf(
					"panic tn the transaction, rolling back transaction with panic err: %+v",
					err,
				)
			} else {
				t.logger.Errorf("panic tn the transaction, rolling back transaction with panic message: %+v", r)
			}
		}
	}()

	err := action(ctx, t)
	if err != nil {
		t.logger.Error("rolling back transaction")
		tx.WithContext(ctx).Rollback()

		return err
	}

	t.logger.Info("committing transaction")

	if err = tx.WithContext(ctx).Commit().Error; err != nil {
		t.logger.Errorf("transaction commit error: %+v", err)
	}

	return err
}

func (t txContextDb) DB() *gorm.DB {
	return t.db
}
