package initialize

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
	customMiddleware "github.com/phanhotboy/nien-su-viet/apps/cms/internal/middleware"

	appController "github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/http"
	appInit "github.com/phanhotboy/nien-su-viet/apps/cms/internal/initialize/app"

	mediaInit "github.com/phanhotboy/nien-su-viet/apps/cms/internal/initialize/media"
	mediaController "github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/controller/http"

	headerNavItemController "github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/http"
	headerNavItemInit "github.com/phanhotboy/nien-su-viet/apps/cms/internal/initialize/headerNavItem"

	footerNavItemController "github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/http"
	footerNavItemInit "github.com/phanhotboy/nien-su-viet/apps/cms/internal/initialize/footerNavItem"
)

// PingOutput is the response for ping endpoints
type PingOutput struct {
	Body struct {
		Code    int    `json:"code" example:"200" doc:"HTTP status code"`
		Message string `json:"message" example:"success" doc:"Response message"`
		Data    string `json:"data" example:"pong" doc:"Response data"`
	}
}

// Ping100Handler handles ping requests
func Ping100Handler(ctx context.Context, input *struct{}) (*PingOutput, error) {
	resp := &PingOutput{}
	resp.Body.Code = 200
	resp.Body.Message = "success"
	resp.Body.Data = "pong"
	return resp, nil
}

func InitRouter(db *gorm.DB, isLogger string) http.Handler {
	// Initialize the Chi router
	r := chi.NewRouter()

	// Middlewares
	if isLogger == "debug" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(customMiddleware.CORS)
	r.Use(customMiddleware.ValidatorMiddleware())

	// Create Huma API configuration
	config := huma.DefaultConfig("Nien Su Viet CMS Service API", "1.0.0")
	config.Info.Description = "This is the API documentation for the Nien Su Viet CMS Service."
	config.Info.Contact = &huma.Contact{
		Name:  "API Support",
		URL:   "https://niensuviet.phannd.me",
		Email: "support@phannd.me",
	}
	config.Servers = []*huma.Server{
		{URL: "http://localhost:" + strconv.Itoa(global.Config.Server.Port), Description: "Local development server"},
	}

	// Create Huma API with Chi adapter
	api := humachi.New(r, config)

	r.Get("/docs", docHandler)

	// Create API v1 group
	grp := huma.NewGroup(api, "/api/v1")

	// Register ping endpoints first
	huma.Register(grp, huma.Operation{
		OperationID:   "get-ping-100",
		Method:        http.MethodGet,
		Path:          "/ping/100",
		Summary:       "Ping endpoint",
		Description:   "Returns pong to check if the service is running",
		Tags:          []string{"health"},
		DefaultStatus: http.StatusOK,
	}, Ping100Handler)

	// Initialize handlers
	appHandler := appInit.InitApp(db)
	mediaHandler := mediaInit.InitMedia(db)
	headerNavItemHandler := headerNavItemInit.InitHeaderNavItem(db)
	footerNavItemHandler := footerNavItemInit.InitFooterNavItem(db)

	// Register module routes
	appController.RegisterAppRoutes(grp, appHandler)
	mediaController.RegisterMediaRoutes(grp, mediaHandler)
	headerNavItemController.RegisterHeaderNavItemHandlers(grp, headerNavItemHandler)
	footerNavItemController.RegisterFooterNavItemRoutes(grp, footerNavItemHandler)

	ExportOpenAPI(api)

	return r
}

func docHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!DOCTYPE html>
				<html lang="en">
					<head>
					  <meta charset="utf-8" />
					  <meta name="viewport" content="width=device-width, initial-scale=1" />
					  <meta name="description" content="SwaggerUI" />
					  <title>SwaggerUI</title>
					  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
					</head>
					<body>
					<div id="swagger-ui"></div>
					<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
					<script>
					  window.onload = () => {
					    window.ui = SwaggerUIBundle({
					      url: '/openapi.json',
					      dom_id: '#swagger-ui',
					    });
					  };
					</script>
					</body>
				</html>`))
}
