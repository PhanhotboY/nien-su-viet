package service

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppService interface {
	GetAppInfo(ctx context.Context) (*entity.App, error)
	UpdateAppInfo(ctx context.Context, app *dto.AppUpdateReq) error
}
