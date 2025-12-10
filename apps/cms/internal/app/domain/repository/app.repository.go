package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppRepository interface {
	GetAppInfo(ctx *gin.Context) (*entity.App, error)
	CreateApp(ctx *gin.Context, app *entity.App) error
	UpdateApp(ctx *gin.Context, app *entity.App) error
}
