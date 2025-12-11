package service

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppService interface {
	GetAppInfo(ctx *gin.Context) (*entity.App, error)
	UpdateAppInfo(ctx *gin.Context, app *dto.AppUpdateDto) (*struct{ Success bool }, error)
}
