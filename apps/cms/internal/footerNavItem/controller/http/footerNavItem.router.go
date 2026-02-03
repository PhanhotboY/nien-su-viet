package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterFooterNavItemRoutes(api huma.API, handler *FooterNavItemHandler) {
	footerNavItems := huma.NewGroup(api, "/footer-nav-items")

	// GET /api/v1/footer-nav-items - Get all footer navigation items
	huma.Register(footerNavItems, huma.Operation{
		OperationID:   "get-footer-nav-items",
		Method:        http.MethodGet,
		Path:          "",
		Summary:       "Get footer nav items",
		Description:   "Retrieve all footer navigation items ordered by display order",
		Tags:          []string{"footer-nav-items"},
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Successful response with footer navigation items",
			},
		},
	}, response.Wrap(handler.GetFooterNavItems))

	// POST /api/v1/footer-nav-items - Create a new footer navigation item
	huma.Register(footerNavItems, huma.Operation{
		OperationID: "create-footer-nav-item",
		Method:      http.MethodPost,
		Path:        "",
		Summary:     "Create footer nav item",
		Description: "Create a new footer navigation item",
		Tags:        []string{"footer-nav-items"},
		// Middlewares: huma.Middlewares{
		// 	middleware.Authentication(api),
		// },
		DefaultStatus: http.StatusCreated,
		Responses: map[string]*huma.Response{
			"201": {
				Description: "Footer navigation item created successfully",
			},
		},
	}, response.Wrap(handler.CreateFooterNavItem))

	// PUT /api/v1/footer-nav-items/{id} - Update a footer navigation item
	huma.Register(footerNavItems, huma.Operation{
		OperationID: "update-footer-nav-item",
		Method:      http.MethodPut,
		Path:        "/{id}",
		Summary:     "Update footer nav item",
		Description: "Update an existing footer navigation item",
		Tags:        []string{"footer-nav-items"},
		// Middlewares: huma.Middlewares{
		// 	middleware.Authentication(api),
		// },
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Footer navigation item updated successfully",
			},
		},
	}, response.Wrap(handler.UpdateFooterNavItem))

	// DELETE /api/v1/footer-nav-items/{id} - Delete a footer navigation item
	huma.Register(footerNavItems, huma.Operation{
		OperationID: "delete-footer-nav-item",
		Method:      http.MethodDelete,
		Path:        "/{id}",
		Summary:     "Delete footer nav item",
		Description: "Delete a footer navigation item",
		Tags:        []string{"footer-nav-items"},
		// Middlewares: huma.Middlewares{
		// 	middleware.Authentication(api),
		// },
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Footer navigation item deleted successfully",
			},
		},
	}, response.Wrap(handler.DeleteFooterNavItem))
}
