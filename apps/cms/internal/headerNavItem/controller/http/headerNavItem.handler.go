package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/dto"
)

type HeaderNavItemHandler struct {
	service service.HeaderNavItemService
}

func NewHeaderNavItemHandler(sv service.HeaderNavItemService) *HeaderNavItemHandler {
	return &HeaderNavItemHandler{service: sv}
}

// GetHeaderNavItemsInput represents the input for getting header nav items
type GetHeaderNavItemsInput struct{}

// GetHeaderNavItemsOutput represents the output for getting header nav items
type GetHeaderNavItemsOutput struct {
	Body dto.HeaderNavItemListResponse
}

// CreateHeaderNavItemInput represents the input for creating a header nav item
type CreateHeaderNavItemInput struct {
	Body dto.HeaderNavItemCreateReq
}

// CreateHeaderNavItemOutput represents the output for creating a header nav item
type CreateHeaderNavItemOutput struct {
	Body dto.OperationResponse
}

// UpdateHeaderNavItemInput represents the input for updating a header nav item
type UpdateHeaderNavItemInput struct {
	ID   string `path:"id" maxLength:"100" example:"item_123" doc:"Header nav item ID"`
	Body dto.HeaderNavItemUpdateReq
}

// UpdateHeaderNavItemOutput represents the output for updating a header nav item
type UpdateHeaderNavItemOutput struct {
	Body dto.OperationResponse
}

// DeleteHeaderNavItemInput represents the input for deleting a header nav item
type DeleteHeaderNavItemInput struct {
	ID string `path:"id" maxLength:"100" example:"item_123" doc:"Header nav item ID"`
}

// DeleteHeaderNavItemOutput represents the output for deleting a header nav item
type DeleteHeaderNavItemOutput struct {
	Body dto.OperationResponse
}

// GetHeaderNavItems retrieves all header navigation items
func (h *HeaderNavItemHandler) GetHeaderNavItems(ctx context.Context, input *GetHeaderNavItemsInput) (*GetHeaderNavItemsOutput, error) {
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

	return &GetHeaderNavItemsOutput{
		Body: dto.HeaderNavItemListResponse{
			Code:    200,
			Message: "success",
			Data:    itemData,
		},
	}, nil
}

// CreateHeaderNavItem creates a new header navigation item
func (h *HeaderNavItemHandler) CreateHeaderNavItem(ctx context.Context, input *CreateHeaderNavItemInput) (*CreateHeaderNavItemOutput, error) {
	result, err := h.service.CreateHeaderNavItem(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return &CreateHeaderNavItemOutput{
		Body: dto.OperationResponse{
			Code:    200,
			Message: "success",
			Data: dto.OperationResult{
				Success: result.Success,
			},
		},
	}, nil
}

// UpdateHeaderNavItem updates an existing header navigation item
func (h *HeaderNavItemHandler) UpdateHeaderNavItem(ctx context.Context, input *UpdateHeaderNavItemInput) (*UpdateHeaderNavItemOutput, error) {
	result, err := h.service.UpdateHeaderNavItem(ctx, input.ID, &input.Body)
	if err != nil {
		return nil, err
	}

	return &UpdateHeaderNavItemOutput{
		Body: dto.OperationResponse{
			Code:    200,
			Message: "success",
			Data: dto.OperationResult{
				Success: result.Success,
			},
		},
	}, nil
}

// DeleteHeaderNavItem deletes a header navigation item
func (h *HeaderNavItemHandler) DeleteHeaderNavItem(ctx context.Context, input *DeleteHeaderNavItemInput) (*DeleteHeaderNavItemOutput, error) {
	result, err := h.service.DeleteHeaderNavItem(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteHeaderNavItemOutput{
		Body: dto.OperationResponse{
			Code:    200,
			Message: "success",
			Data: dto.OperationResult{
				Success: result.Success,
			},
		},
	}, nil
}
