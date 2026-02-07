package service

import (
	"context"
	"fmt"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/repository"
)

type headerNavItemService struct {
	repo repository.HeaderNavItemRepository
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

// CreateHeaderNavItem implements HeaderNavItemService.
func (h *headerNavItemService) CreateHeaderNavItem(ctx context.Context, item *dto.HeaderNavItemCreateReq) (string, error) {
	id, err := h.repo.CreateNavItem(ctx, item.MapToEntity())
	if err != nil {
		return id, fmt.Errorf("failed to create header nav item: %w", err)
	}
	return id, nil
}

// UpdateHeaderNavItem implements HeaderNavItemService.
func (h *headerNavItemService) UpdateHeaderNavItem(ctx context.Context, id string, item *dto.HeaderNavItemUpdateReq) (string, error) {
	id, err := h.repo.UpdateNavItem(ctx, id, item.MapToEntity())
	if err != nil {
		return id, fmt.Errorf("failed to update header nav item: %w", err)
	}
	return id, nil
}

// DeleteHeaderNavItem implements HeaderNavItemService.
func (h *headerNavItemService) DeleteHeaderNavItem(ctx context.Context, id string) (string, error) {
	id, err := h.repo.DeleteNavItem(ctx, id)
	if err != nil {
		return id, fmt.Errorf("failed to delete header nav item: %w", err)
	}
	return id, nil
}

func NewHeaderNavItemService(repo repository.HeaderNavItemRepository) HeaderNavItemService {
	return &headerNavItemService{repo}
}
