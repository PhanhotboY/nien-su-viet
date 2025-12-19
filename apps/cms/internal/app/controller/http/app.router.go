package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterAppRoutes(grp *huma.Group, handler *AppHandler) {
	app := huma.NewGroup(grp, "/app")
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
	}, response.Wrap(handler.UpdateAppInfo))
}
