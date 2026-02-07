package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppRepository interface {
	GetAppInfo(ctx context.Context) (*entity.App, error)
	CreateApp(ctx context.Context, app *entity.App) (string, error)
	UpdateApp(ctx context.Context, appId string, app *entity.App) (string, error)
}
