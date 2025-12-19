package service

import (
	"context"
	"fmt"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/repository"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type headerNavItemService struct {
	repo repository.HeaderNavItemRepository
}

// DeleteHeaderNavItem implements HeaderNavItemService.
func (h *headerNavItemService) DeleteHeaderNavItem(ctx context.Context, id string) (*response.OperationResult, error) {
	err := h.repo.DeleteNavItem(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete header nav item: %w", err)
	}
	return &response.OperationResult{Success: true}, nil
}

// CreateHeaderNavItem implements HeaderNavItemService.
func (h *headerNavItemService) CreateHeaderNavItem(ctx context.Context, item *dto.HeaderNavItemCreateReq) (*response.OperationResult, error) {
	err := h.repo.CreateNavItem(ctx, item.MapToEntity())
	if err != nil {
		return nil, fmt.Errorf("failed to create header nav item: %w", err)
	}
	return &response.OperationResult{Success: true}, nil
}

// GetHeaderNavItems implements HeaderNavItemService.
func (h *headerNavItemService) GetHeaderNavItems(ctx context.Context) ([]*dto.HeaderNavItemData, error) {
	items, err := h.repo.GetNavItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get header nav items: %w", err)
	}
	result := make([]*dto.HeaderNavItemData, 0, len(items))
	for _, item := range items {
		dtoItem := &dto.HeaderNavItemData{}
		dtoItem.FromEntity(item)
		result = append(result, dtoItem)
	}
	return result, nil
}

// UpdateHeaderNavItem implements HeaderNavItemService.
func (h *headerNavItemService) UpdateHeaderNavItem(ctx context.Context, id string, item *dto.HeaderNavItemUpdateReq) (*response.OperationResult, error) {
	err := h.repo.UpdateNavItem(ctx, id, item.MapToEntity())
	if err != nil {
		return nil, fmt.Errorf("failed to update header nav item: %w", err)
	}
	return &response.OperationResult{Success: true}, nil
}

func NewHeaderNavItemService(repo repository.HeaderNavItemRepository) HeaderNavItemService {
	return &headerNavItemService{repo}
}
