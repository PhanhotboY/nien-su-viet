package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/request"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type FooterNavItemHandler struct {
	service service.FooterNavItemService
}

func NewFooterNavItemHandler(sv service.FooterNavItemService) *FooterNavItemHandler {
	return &FooterNavItemHandler{service: sv}
}

// PathIDInput represents input with a path parameter ID
type PathIDInput struct {
	ID string `path:"id" maxLength:"100" example:"item_123" doc:"Footer nav item ID"`
}

// GetFooterNavItems retrieves all footer navigation items
func (h *FooterNavItemHandler) GetFooterNavItems(ctx context.Context, input *struct{}) (*response.APIBodyResponse[[]dto.FooterNavItemData], error) {
	items, err := h.service.GetFooterNavItems(ctx)
	if err != nil {
		return nil, err
	}

	// Convert pointers to values
	itemData := make([]dto.FooterNavItemData, 0, len(items))
	for _, item := range items {
		if item != nil {
			itemData = append(itemData, *item)
		}
	}

	return response.SuccessResponse(200, itemData), nil
}

// CreateFooterNavItem creates a new footer navigation item
func (h *FooterNavItemHandler) CreateFooterNavItem(ctx context.Context, input *request.APIBodyRequest[dto.FooterNavItemCreateReq]) (*response.APIBodyResponse[response.OperationResult], error) {
	err := h.service.CreateFooterNavItem(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(201), nil
}

// UpdateFooterNavItem updates an existing footer navigation item
func (h *FooterNavItemHandler) UpdateFooterNavItem(ctx context.Context, input *struct {
	PathIDInput
	request.APIBodyRequest[dto.FooterNavItemUpdateReq]
}) (*response.APIBodyResponse[response.OperationResult], error) {
	err := h.service.UpdateFooterNavItem(ctx, input.ID, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(200), nil
}

// DeleteFooterNavItem deletes a footer navigation item
func (h *FooterNavItemHandler) DeleteFooterNavItem(ctx context.Context, input *PathIDInput) (*response.APIBodyResponse[response.OperationResult], error) {
	err := h.service.DeleteFooterNavItem(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(200), nil
}
