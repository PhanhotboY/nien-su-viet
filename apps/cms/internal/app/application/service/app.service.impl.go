package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/repository"
)

type appService struct {
	appRepo repository.AppRepository
}

func NewAppService(repo repository.AppRepository) AppService {
	return &appService{appRepo: repo}
}

func (a *appService) UpdateAppInfo(ctx context.Context, app *dto.AppUpdateReq) (string, error) {
	foundApp, err := a.appRepo.GetAppInfo(ctx)
	if err != nil {
		return "", err
	}
	if foundApp == nil {
		return "", fmt.Errorf("App info not found")
	}
	id, err := a.appRepo.UpdateApp(ctx, foundApp.AppId, app.MapToEntity())
	if err != nil {
		return id, err
	}
	return id, nil
}

func (a *appService) GetAppInfo(ctx context.Context) (*entity.App, error) {
	app, err := a.appRepo.GetAppInfo(ctx)
	if app != nil {
		return app, nil
	}

	if err != nil && strings.Contains(err.Error(), "not found") {
		// Create default app info
		_, err = a.appRepo.CreateApp(ctx, &entity.App{
			Title: "Nien Su Viet",
		})
		if err != nil {
			return nil, err
		}
		app, err = a.appRepo.GetAppInfo(ctx)
		if app != nil {
			return app, nil
		}
	}

	return nil, err
}
