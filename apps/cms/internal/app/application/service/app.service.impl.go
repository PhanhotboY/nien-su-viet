package service

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/repository"
)

type appService struct {
	appRepo repository.AppRepository
}

// UpdateAppInfo implements AppService.
func (a *appService) UpdateAppInfo(ctx *gin.Context, app *dto.AppUpdateDto) (*struct{ Success bool }, error) {
	err := a.appRepo.UpdateApp(ctx, app.MapToEntity())
	if err != nil {
		return nil, err
	}
	return &struct{ Success bool }{Success: true}, nil
}

// GetAppInfo implements AppService.
func (a *appService) GetAppInfo(ctx *gin.Context) (*entity.App, error) {
	app, err := a.appRepo.GetAppInfo(ctx)
	if app != nil {
		return app, nil
	}

	if err != nil && err.Error() == "record not found" {
		// Create default app info
		err = a.appRepo.CreateApp(ctx, &entity.App{
			Title: "Nien Su Viet",
		})
		if err != nil {
			return nil, err
		}
		// Retrieve the newly created app info
		return a.appRepo.GetAppInfo(ctx)
	}
	return nil, err
}

func NewAppService(appRepo repository.AppRepository) AppService {
	return &appService{appRepo}
}
