package service

import (
	"context"
	"fmt"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/domain/repository"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type footerNavItemService struct {
	repo repository.FooterNavItemRepository
}

func NewFooterNavItemService(repo repository.FooterNavItemRepository) FooterNavItemService {
	return &footerNavItemService{repo}
}

// GetFooterNavItems returns list of footer nav items
func (s *footerNavItemService) GetFooterNavItems(ctx context.Context) ([]*dto.FooterNavItemData, error) {
	items, err := s.repo.GetNavItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get footer nav items: %w", err)
	}
	result := make([]*dto.FooterNavItemData, 0, len(items))
	for _, item := range items {
		d := &dto.FooterNavItemData{}
		d.FromEntity(item)
		result = append(result, d)
	}
	return result, nil
}

// CreateFooterNavItem creates a footer nav item
func (s *footerNavItemService) CreateFooterNavItem(ctx context.Context, item *dto.FooterNavItemCreateReq) (*response.OperationResult, error) {
	if err := s.repo.CreateNavItem(ctx, item.MapToEntity()); err != nil {
		return nil, fmt.Errorf("failed to create footer nav item: %w", err)
	}
	return &response.OperationResult{Success: true}, nil
}

// UpdateFooterNavItem updates a footer nav item
func (s *footerNavItemService) UpdateFooterNavItem(ctx context.Context, id string, item *dto.FooterNavItemUpdateReq) (*response.OperationResult, error) {
	if err := s.repo.UpdateNavItem(ctx, id, item.MapToEntity()); err != nil {
		return nil, fmt.Errorf("failed to update footer nav item: %w", err)
	}
	return &response.OperationResult{Success: true}, nil
}

// DeleteFooterNavItem deletes a footer nav item
func (s *footerNavItemService) DeleteFooterNavItem(ctx context.Context, id string) (*response.OperationResult, error) {
	if err := s.repo.DeleteNavItem(ctx, id); err != nil {
		return nil, fmt.Errorf("failed to delete footer nav item: %w", err)
	}
	return &response.OperationResult{Success: true}, nil
}
