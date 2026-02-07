package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/domain/repository"
	"gorm.io/gorm"
)

type footerNavItemRepository struct {
	db *gorm.DB
}

func (r *footerNavItemRepository) GetNavItems(ctx context.Context) ([]*entity.FooterNavItem, error) {
	var items []*entity.FooterNavItem
	res := r.db.WithContext(ctx).Model(&entity.FooterNavItem{}).Find(&items)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to get footer nav items: %w", res.Error)
	}
	return items, nil
}

func (r *footerNavItemRepository) FindNavItemById(ctx context.Context, id string) (*entity.FooterNavItem, error) {
	var item entity.FooterNavItem
	res := r.db.WithContext(ctx).Model(&entity.FooterNavItem{}).Where("id = ?", id).First(&item)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find footer nav item: %w", res.Error)
	}
	return &item, nil
}

func (r *footerNavItemRepository) CreateNavItem(ctx context.Context, navItem *entity.FooterNavItem) (string, error) {
	if navItem.Id == "" {
		navItem.Id = uuid.New().String()
	}
	err := r.db.WithContext(ctx).Model(&entity.FooterNavItem{}).Create(navItem).Error
	return navItem.Id, err
}

func (r *footerNavItemRepository) UpdateNavItem(ctx context.Context, id string, navItem *entity.FooterNavItem) (string, error) {
	err := r.db.WithContext(ctx).Model(&entity.FooterNavItem{}).Where("id = ?", id).Updates(navItem).Error
	return id, err
}

func (r *footerNavItemRepository) DeleteNavItem(ctx context.Context, id string) (string, error) {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.FooterNavItem{}).Error
	return id, err
}

func NewFooterNavItemRepository(db *gorm.DB) repository.FooterNavItemRepository {
	return &footerNavItemRepository{db}
}
