package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/repository"
	"gorm.io/gorm"
)

type appRepository struct {
	db *gorm.DB
}

// UpdateApp implements repository.AppRepository.
func (ar *appRepository) UpdateApp(ctx *gin.Context, app *entity.App) error {
	result := ar.db.Model(&entity.App{}).Where("app_id = ?", app.AppId).Updates(app)
	if result.Error != nil {
		log.Printf("Error updating app: %v", result.Error)
		return result.Error
	}
	return nil
}

func (ar *appRepository) GetAppInfo(ctx *gin.Context) (*entity.App, error) {
	var app entity.App
	result := ar.db.Model(&entity.App{}).First(&app)
	if result.Error != nil {
		log.Printf("Error fetching app info: %v", result.Error)
		return nil, result.Error
	}
	return &app, nil
}

func (ar *appRepository) CreateApp(ctx *gin.Context, app *entity.App) error {
	result := ar.db.Model(&entity.App{}).Create(app)
	if result.Error != nil {
		log.Printf("Error creating app: %v", result.Error)
		return result.Error
	}
	return nil
}

func NewAppRepository(db *gorm.DB) repository.AppRepository {
	return &appRepository{db}
}
