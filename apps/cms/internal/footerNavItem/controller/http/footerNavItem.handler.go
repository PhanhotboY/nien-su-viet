package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/dto"
)

type FooterNavItemHandler struct {
	service service.FooterNavItemService
}

func NewFooterNavItemHandler(sv service.FooterNavItemService) *FooterNavItemHandler {
	return &FooterNavItemHandler{service: sv}
}

// GetFooterNavItemsInput represents the input for getting footer nav items
type GetFooterNavItemsInput struct{}

// GetFooterNavItemsOutput represents the output for getting footer nav items
type GetFooterNavItemsOutput struct {
	Body dto.FooterNavItemListResponse
}

// CreateFooterNavItemInput represents the input for creating a footer nav item
type CreateFooterNavItemInput struct {
	Body dto.FooterNavItemCreateReq
}

// CreateFooterNavItemOutput represents the output for creating a footer nav item
type CreateFooterNavItemOutput struct {
	Body dto.OperationResponse
}

// UpdateFooterNavItemInput represents the input for updating a footer nav item
type UpdateFooterNavItemInput struct {
	ID   string `path:"id" maxLength:"100" example:"item_123" doc:"Footer nav item ID"`
	Body dto.FooterNavItemUpdateReq
}

// UpdateFooterNavItemOutput represents the output for updating a footer nav item
type UpdateFooterNavItemOutput struct {
	Body dto.OperationResponse
}

// DeleteFooterNavItemInput represents the input for deleting a footer nav item
type DeleteFooterNavItemInput struct {
	ID string `path:"id" maxLength:"100" example:"item_123" doc:"Footer nav item ID"`
}

// DeleteFooterNavItemOutput represents the output for deleting a footer nav item
type DeleteFooterNavItemOutput struct {
	Body dto.OperationResponse
}

// GetFooterNavItems retrieves all footer navigation items
func (h *FooterNavItemHandler) GetFooterNavItems(ctx context.Context, input *GetFooterNavItemsInput) (*GetFooterNavItemsOutput, error) {
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

	return &GetFooterNavItemsOutput{
		Body: dto.FooterNavItemListResponse{
			Code:    200,
			Message: "success",
			Data:    itemData,
		},
	}, nil
}

// CreateFooterNavItem creates a new footer navigation item
func (h *FooterNavItemHandler) CreateFooterNavItem(ctx context.Context, input *CreateFooterNavItemInput) (*CreateFooterNavItemOutput, error) {
	result, err := h.service.CreateFooterNavItem(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return &CreateFooterNavItemOutput{
		Body: dto.OperationResponse{
			Code:    200,
			Message: "success",
			Data: dto.OperationResult{
				Success: result.Success,
			},
		},
	}, nil
}

// UpdateFooterNavItem updates an existing footer navigation item
func (h *FooterNavItemHandler) UpdateFooterNavItem(ctx context.Context, input *UpdateFooterNavItemInput) (*UpdateFooterNavItemOutput, error) {
	result, err := h.service.UpdateFooterNavItem(ctx, input.ID, &input.Body)
	if err != nil {
		return nil, err
	}

	return &UpdateFooterNavItemOutput{
		Body: dto.OperationResponse{
			Code:    200,
			Message: "success",
			Data: dto.OperationResult{
				Success: result.Success,
			},
		},
	}, nil
}

// DeleteFooterNavItem deletes a footer navigation item
func (h *FooterNavItemHandler) DeleteFooterNavItem(ctx context.Context, input *DeleteFooterNavItemInput) (*DeleteFooterNavItemOutput, error) {
	result, err := h.service.DeleteFooterNavItem(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteFooterNavItemOutput{
		Body: dto.OperationResponse{
			Code:    200,
			Message: "success",
			Data: dto.OperationResult{
				Success: result.Success,
			},
		},
	}, nil
}
