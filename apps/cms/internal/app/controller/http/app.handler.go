package http

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppHandler struct {
	service service.AppService
}

func NewAppHandler(service service.AppService) *AppHandler {
	return &AppHandler{service}
}

// @Summary Get app info
// @Description Get website information (create if not exist)
// @Tags app
// @Accept json
// @Produce json
// @Success 200 {object} dto.AppInfoResponse
// @Router /app [get]
func (h *AppHandler) GetAppInfo(ctx *gin.Context) (*entity.App, error) {
	return h.service.GetAppInfo(ctx)
}

// @Summary Update app info
// @Description Update website information
// @Tags app
// @Accept json
// @Produce json
// @Param app body dto.AppUpdateDto true "App Info"
// @Success 200 {object} response.APIOperationResponse
// @Failure 400 {object} response.APIError
// @Router /app [put]
func (h *AppHandler) UpdateAppInfo(ctx *gin.Context) (*struct{ Success bool }, error) {
	var app dto.AppUpdateDto
	if err := ctx.ShouldBindJSON(&app); err != nil {
		return nil, err
	}
	return h.service.UpdateAppInfo(ctx, &app)
}
