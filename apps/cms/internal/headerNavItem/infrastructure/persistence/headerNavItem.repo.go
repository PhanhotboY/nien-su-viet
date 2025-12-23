package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/repository"
	"gorm.io/gorm"
)

type headerNavItemRepository struct {
	db *gorm.DB
}

// DeleteNavItem implements repository.HeaderNavItemRepository.
func (h *headerNavItemRepository) DeleteNavItem(ctx context.Context, id string) error {
	return h.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.HeaderNavItem{}).Error
}

// CreateNavItems implements repository.HeaderNavItemRepository.
func (h *headerNavItemRepository) CreateNavItem(ctx context.Context, navItem *entity.HeaderNavItem) error {
	if navItem.Id == "" {
		navItem.Id = uuid.New().String()
	}
	return h.db.WithContext(ctx).Model(&entity.HeaderNavItem{}).Create(navItem).Error

}

// FindNavItemById implements repository.HeaderNavItemRepository.
func (h *headerNavItemRepository) FindNavItemById(ctx context.Context, id string) (*entity.HeaderNavItem, error) {
	var navItem entity.HeaderNavItem
	result := h.db.WithContext(ctx).Model(&entity.HeaderNavItem{}).Where(&entity.HeaderNavItem{Id: id}).First(&navItem)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find header nav item by id: %w", result.Error)
	}
	return &navItem, nil
}

// GetNavItems implements repository.HeaderNavItemRepository.
func (h *headerNavItemRepository) GetNavItems(ctx context.Context) ([]*entity.HeaderNavItem, error) {
	// Apply a short timeout to avoid hanging requests
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	var navItems []*entity.HeaderNavItem
	result := h.db.WithContext(ctx).Model(&entity.HeaderNavItem{}).Find(&navItems)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get header nav items: %w", result.Error)
	}
	return navItems, nil
}

// UpdateNavItems implements repository.HeaderNavItemRepository.
func (h *headerNavItemRepository) UpdateNavItem(ctx context.Context, id string, navItem *entity.HeaderNavItem) error {
	return h.db.WithContext(ctx).Model(&entity.HeaderNavItem{}).Where("id = ?", id).Updates(navItem).Error

}

func NewHeaderNavItemRepository(db *gorm.DB) repository.HeaderNavItemRepository {
	return &headerNavItemRepository{db}
}
