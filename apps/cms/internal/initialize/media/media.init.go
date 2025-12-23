package initialize

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/controller/http"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/infrastructure/persistence"
	"gorm.io/gorm"
)

func InitMedia(db *gorm.DB) *http.MediaHandler {
	mediaRepo := persistence.NewMediaRepository(db)
	mediaService := service.NewMediaService(mediaRepo)
	return http.NewMediaHandler(mediaService)
}
