package persistence

import (
	"context"
	"fmt"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/repository"
	"gorm.io/gorm"
)

type appRepository struct {
	db *gorm.DB
}

func NewAppRepository(db *gorm.DB) repository.AppRepository {
	return &appRepository{db}
}

// UpdateApp implements repository.AppRepository.
func (ar *appRepository) UpdateApp(ctx context.Context, appId string, app *entity.App) error {
	result := ar.db.WithContext(ctx).Model(&entity.App{}).Where("id = ?", appId).Updates(app)
	if result.Error != nil {
		return fmt.Errorf("failed to update app: %w", result.Error)
	}
	return nil
}

func (ar *appRepository) GetAppInfo(ctx context.Context) (*entity.App, error) {
	var app entity.App
	result := ar.db.WithContext(ctx).Model(&entity.App{}).First(&app)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch app info: %w", result.Error)
	}
	return &app, nil
}

func (ar *appRepository) CreateApp(ctx context.Context, app *entity.App) error {
	result := ar.db.WithContext(ctx).Model(&entity.App{}).Create(app)
	if result.Error != nil {
		return fmt.Errorf("failed to create app: %w", result.Error)
	}
	return nil
}
