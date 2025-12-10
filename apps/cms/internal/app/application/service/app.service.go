package service

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppService interface {
	GetAppInfo(ctx *gin.Context) (*entity.App, error)
}
