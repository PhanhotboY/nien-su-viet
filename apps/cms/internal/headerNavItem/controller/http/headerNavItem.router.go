package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHeaderNavItemHandlers(api huma.API, handler *HeaderNavItemHandler) {
	// GET /api/v1/header-nav-items - Get all header navigation items
	huma.Register(api, huma.Operation{
		OperationID:   "get-header-nav-items",
		Method:        http.MethodGet,
		Path:          "/api/v1/header-nav-items",
		Summary:       "Get header nav items",
		Description:   "Retrieve all header navigation items ordered by display order",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.GetHeaderNavItems)

	// POST /api/v1/header-nav-items - Create a new header navigation item
	huma.Register(api, huma.Operation{
		OperationID:   "create-header-nav-item",
		Method:        http.MethodPost,
		Path:          "/api/v1/header-nav-items",
		Summary:       "Create header nav item",
		Description:   "Create a new header navigation item",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.CreateHeaderNavItem)

	// PUT /api/v1/header-nav-items/{id} - Update a header navigation item
	huma.Register(api, huma.Operation{
		OperationID:   "update-header-nav-item",
		Method:        http.MethodPut,
		Path:          "/api/v1/header-nav-items/{id}",
		Summary:       "Update header nav item",
		Description:   "Update an existing header navigation item",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.UpdateHeaderNavItem)

	// DELETE /api/v1/header-nav-items/{id} - Delete a header navigation item
	huma.Register(api, huma.Operation{
		OperationID:   "delete-header-nav-item",
		Method:        http.MethodDelete,
		Path:          "/api/v1/header-nav-items/{id}",
		Summary:       "Delete header nav item",
		Description:   "Delete a header navigation item",
		Tags:          []string{"header-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.DeleteHeaderNavItem)
}
