package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/phanhotboy/nien-su-viet/apps/cms/docs"
	appController "github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/http"
	initialize "github.com/phanhotboy/nien-su-viet/apps/cms/internal/initialize/app"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/middleware"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type StringAPIResponse struct {
	response.APIResponse[string]
}

// === Health check endpoints
// @Summary ping 100 response without response wrapper
// @Schemes
// @Description do ping manual response
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} StringAPIResponse
// @Router /ping/100 [get]
func Ping100Handler(ctx *gin.Context) {
	response.SuccessResponse(ctx, "pong")
}

// @Summary ping 200 response with response wrapper
// @Schemes
// @Description do ping wrapped response
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} StringAPIResponse
// @Router /ping/200 [get]
func Ping200Handler(ctx *gin.Context) (string, error) {
	return "pong", nil
}

func InitRouter(db *gorm.DB, isLogger string) *gin.Engine {
	// Initialize the router
	// This function will set up the routes and middleware for the application
	// It will return a gin.Engine instance that can be used to run the server

	var r *gin.Engine
	// Set the mode based on the environment
	if isLogger == "debug" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	r.Use(middleware.CORS) // cross
	r.Use(middleware.ValidatorMiddleware())
	// r.Use() // logging

	// r.Use() // limiter global
	// r.Use(middlewares.Validator())      // middleware

	// r.Use(middleware.NewRateLimiter().GlobalRateLimiter()) // 100 req/s

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	// === Đăng ký routes theo module
	v1 := r.Group("/api/v1")

	v1.GET("/ping/100", Ping100Handler)
	v1.GET("/ping/200", response.Wrap(Ping200Handler))

	// Register the app routes
	// === DI các handler
	appHandler := initialize.InitApp(db)
	appController.RegisterAppRoutes(v1, appHandler)

	return r
}
