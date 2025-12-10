package http

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/application/service"
	appEntity "github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type AppHandler struct {
	service service.AppService
}

func NewAppHandler(service service.AppService) *AppHandler {
	return &AppHandler{service}
}

type AppInfoResponse struct {
	response.APIResponse[appEntity.App]
}

// @Summary Get app info
// @Description Get website information (create if not exist)
// @Tags app
// @Accept json
// @Produce json
// @Success 200 {object} AppInfoResponse
// @Router /app [get]
func (h *AppHandler) GetAppInfo(ctx *gin.Context) (*appEntity.App, error) {
	return h.service.GetAppInfo(ctx)
}
