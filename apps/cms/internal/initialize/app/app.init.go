package initialize

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/http"
	persistence "github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/infrastructure/persistence"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB) *http.AppHandler {
	appRepo := persistence.NewAppRepository(db)
	appService := service.NewAppService(appRepo)
	return http.NewAppHandler(appService)
}
