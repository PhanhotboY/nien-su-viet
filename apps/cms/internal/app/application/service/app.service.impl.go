package service

import (
	"context"
	"strconv"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/repository"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type appService struct {
	appRepo repository.AppRepository
}

func NewAppService(repo repository.AppRepository) AppService {
	return &appService{appRepo: repo}
}

func (a *appService) UpdateAppInfo(ctx context.Context, app *dto.AppUpdateReq) error {
	foundApp, err := a.appRepo.GetAppInfo(ctx)
	if err != nil {
		return err
	}
	if foundApp == nil {
		return response.NewNotFoundError("App not found", nil)
	}
	err = a.appRepo.UpdateApp(ctx, strconv.Itoa(int(foundApp.AppId)), app.MapToEntity())
	if err != nil {
		return err
	}
	return nil
}

func (a *appService) GetAppInfo(ctx context.Context) (*entity.App, error) {
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
		return a.appRepo.GetAppInfo(ctx)
	}

	return nil, err
}
