package initialize

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
	appEntity "github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	footerNavItemEntity "github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/domain/model/entity"
	headerNavItemEntity "github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/model/entity"
	postEntity "github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
)

func InitPostgreSQL() (*gorm.DB, error) {
	cfg := global.Config.Postgresql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	var err error
	global.PostgresDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := global.PostgresDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	// Safe auto migration - will only modify schema if needed
	if err := global.PostgresDB.AutoMigrate(
		&appEntity.App{},
		&postEntity.Post{},
		&headerNavItemEntity.HeaderNavItem{}, &footerNavItemEntity.FooterNavItem{},
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

	fmt.Println("Database connection established successfully!")
	return global.PostgresDB, nil
}
