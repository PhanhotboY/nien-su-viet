package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterHeaderNavItemHandlers(api huma.API, handler *HeaderNavItemHandler) {
	headerNavItems := huma.NewGroup(api, "/header-nav-items")

	// GET /api/v1/header-nav-items - Get all header navigation items
	huma.Register(headerNavItems, huma.Operation{
		OperationID:   "get-header-nav-items",
		Method:        http.MethodGet,
		Path:          "",
		Summary:       "Get header nav items",
		Description:   "Retrieve all header navigation items ordered by display order",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Successful response with header navigation items",
			},
		},
	}, response.Wrap(handler.GetHeaderNavItems))

	// POST /api/v1/header-nav-items - Create a new header navigation item
	huma.Register(headerNavItems, huma.Operation{
		OperationID: "create-header-nav-item",
		Method:      http.MethodPost,
		Path:        "",
		Summary:     "Create header nav item",
		Description: "Create a new header navigation item",
		Tags:        []string{"header-nav-items"},
		// Middlewares: huma.Middlewares{
		// 	middleware.Authentication(api),
		// },
		DefaultStatus: http.StatusCreated,
		Responses: map[string]*huma.Response{
			"201": {
				Description: "Header navigation item created successfully",
			},
		},
	}, response.Wrap(handler.CreateHeaderNavItem))

	// PUT /api/v1/header-nav-items/{id} - Update a header navigation item
	huma.Register(headerNavItems, huma.Operation{
		OperationID:   "update-header-nav-item",
		Method:        http.MethodPut,
		Path:          "/{id}",
		Summary:       "Update header nav item",
		Description:   "Update an existing header navigation item",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
		// Middlewares: huma.Middlewares{
		// 	middleware.Authentication(api),
		// },
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Header navigation item updated successfully",
			},
		},
	}, response.Wrap(handler.UpdateHeaderNavItem))

	// DELETE /api/v1/header-nav-items/{id} - Delete a header navigation item
	huma.Register(headerNavItems, huma.Operation{
		OperationID:   "delete-header-nav-item",
		Method:        http.MethodDelete,
		Path:          "/{id}",
		Summary:       "Delete header nav item",
		Description:   "Delete a header navigation item",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
		// Middlewares: huma.Middlewares{
		// 	middleware.Authentication(api),
		// },
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Header navigation item deleted successfully",
			},
		},
	}, response.Wrap(handler.DeleteHeaderNavItem))
}
