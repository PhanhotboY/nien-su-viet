package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterFooterNavItemRoutes(api huma.API, handler *FooterNavItemHandler) {
	// GET /api/v1/footer-nav-items - Get all footer navigation items
	huma.Register(api, huma.Operation{
		OperationID:   "get-footer-nav-items",
		Method:        http.MethodGet,
		Path:          "/api/v1/footer-nav-items",
		Summary:       "Get footer nav items",
		Description:   "Retrieve all footer navigation items ordered by display order",
		Tags:          []string{"footer-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.GetFooterNavItems)

	// POST /api/v1/footer-nav-items - Create a new footer navigation item
	huma.Register(api, huma.Operation{
		OperationID:   "create-footer-nav-item",
		Method:        http.MethodPost,
		Path:          "/api/v1/footer-nav-items",
		Summary:       "Create footer nav item",
		Description:   "Create a new footer navigation item",
		Tags:          []string{"footer-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.CreateFooterNavItem)

	// PUT /api/v1/footer-nav-items/{id} - Update a footer navigation item
	huma.Register(api, huma.Operation{
		OperationID:   "update-footer-nav-item",
		Method:        http.MethodPut,
		Path:          "/api/v1/footer-nav-items/{id}",
		Summary:       "Update footer nav item",
		Description:   "Update an existing footer navigation item",
		Tags:          []string{"footer-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.UpdateFooterNavItem)

	// DELETE /api/v1/footer-nav-items/{id} - Delete a footer navigation item
	huma.Register(api, huma.Operation{
		OperationID:   "delete-footer-nav-item",
		Method:        http.MethodDelete,
		Path:          "/api/v1/footer-nav-items/{id}",
		Summary:       "Delete footer nav item",
		Description:   "Delete a footer navigation item",
		Tags:          []string{"footer-nav-items"},
		DefaultStatus: http.StatusOK,
	}, handler.DeleteFooterNavItem)
}
