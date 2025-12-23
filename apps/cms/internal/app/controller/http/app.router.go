package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/middleware"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterAppRoutes(api huma.API, handler *AppHandler) {
	app := huma.NewGroup(api, "/app")
	// GET /api/v1/app - Get app information
	huma.Register(app, huma.Operation{
		OperationID:   "get-app-info",
		Method:        http.MethodGet,
		Path:          "",
		Summary:       "Get app info",
		Description:   "Retrieve application information including title, description, logo, social links, and contact details",
		Tags:          []string{"app"},
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Successful response with app information",
			},
		},
	}, response.Wrap(handler.GetAppInfo))

	// PUT /api/v1/app - Update app information
	huma.Register(app, huma.Operation{
		OperationID:   "update-app-info",
		Method:        http.MethodPut,
		Path:          "",
		Summary:       "Update app info",
		Description:   "Update application information",
		Tags:          []string{"app"},
		DefaultStatus: http.StatusOK,
		Middlewares: huma.Middlewares{
			middleware.Authentication(api),
		},
	}, response.Wrap(handler.UpdateAppInfo))
}
