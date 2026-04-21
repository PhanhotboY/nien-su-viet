package postgres

import (
	"fmt"
	"strings"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
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

	Models []any `group:"db_models"`
}

// NewDb func for connection to PostgreSQL database.
func NewDb(s settings.Config, logger logger.Logger, params DBParams) (*gorm.DB, error) {
	cfg := s.Postgresql
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
	if err := pg.AutoMigrate(
		params.Models...,
	); err != nil {
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
	return pg, nil
}
