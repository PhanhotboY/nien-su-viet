package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/domain/model/entity"
)

type FooterNavItemRepository interface {
	GetNavItems(ctx context.Context) ([]*entity.FooterNavItem, error)
	FindNavItemById(ctx context.Context, id string) (*entity.FooterNavItem, error)
	CreateNavItem(ctx context.Context, navItem *entity.FooterNavItem) (string, error)
	UpdateNavItem(ctx context.Context, id string, navItem *entity.FooterNavItem) (string, error)
	DeleteNavItem(ctx context.Context, id string) (string, error)
}
