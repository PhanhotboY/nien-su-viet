package http

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterAppRoutes(rg *gin.RouterGroup, handler *AppHandler) {
	app := rg.Group("/app")
	// app.Use(middleware.AuthGuardMiddlewareWithHMAC())
	app.GET("/", response.Wrap(handler.GetAppInfo))
}
