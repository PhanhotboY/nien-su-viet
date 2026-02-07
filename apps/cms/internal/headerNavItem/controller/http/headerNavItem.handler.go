package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/request"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type HeaderNavItemHandler struct {
	service service.HeaderNavItemService
}

func NewHeaderNavItemHandler(sv service.HeaderNavItemService) *HeaderNavItemHandler {
	return &HeaderNavItemHandler{service: sv}
}

// PathIDInput represents input with a path parameter ID
type PathIDInput struct {
	ID string `path:"id" maxLength:"100" example:"item_123" doc:"Header nav item ID"`
}

// GetHeaderNavItems retrieves all header navigation items
func (h *HeaderNavItemHandler) GetHeaderNavItems(ctx context.Context, input *struct{}) (*response.APIBodyResponse[[]dto.HeaderNavItemData], error) {
	items, err := h.service.GetHeaderNavItems(ctx)
	if err != nil {
		return nil, err
	}

	// Convert pointers to values
	itemData := make([]dto.HeaderNavItemData, 0, len(items))
	for _, item := range items {
		if item != nil {
			itemData = append(itemData, *item)
		}
	}

	return response.SuccessResponse(200, itemData), nil
}

// CreateHeaderNavItem creates a new header navigation item
func (h *HeaderNavItemHandler) CreateHeaderNavItem(ctx context.Context, input *request.APIBodyRequest[dto.HeaderNavItemCreateReq]) (*response.APIBodyResponse[response.OperationResult], error) {
	id, err := h.service.CreateHeaderNavItem(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(201, id), nil
}

// UpdateHeaderNavItem updates an existing header navigation item
func (h *HeaderNavItemHandler) UpdateHeaderNavItem(ctx context.Context, input *struct {
	PathIDInput
	request.APIBodyRequest[dto.HeaderNavItemUpdateReq]
}) (*response.APIBodyResponse[response.OperationResult], error) {
	id, err := h.service.UpdateHeaderNavItem(ctx, input.ID, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(200, id), nil
}

// DeleteHeaderNavItem deletes a header navigation item
func (h *HeaderNavItemHandler) DeleteHeaderNavItem(ctx context.Context, input *PathIDInput) (*response.APIBodyResponse[response.OperationResult], error) {
	id, err := h.service.DeleteHeaderNavItem(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(200, id), nil
}
