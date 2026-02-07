package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/model/entity"
)

type HeaderNavItemRepository interface {
	GetNavItems(ctx context.Context) ([]*entity.HeaderNavItem, error)
	FindNavItemById(ctx context.Context, id string) (*entity.HeaderNavItem, error)
	UpdateNavItem(ctx context.Context, id string, navItem *entity.HeaderNavItem) (string, error)
	CreateNavItem(ctx context.Context, navItem *entity.HeaderNavItem) (string, error)
	DeleteNavItem(ctx context.Context, id string) (string, error)
}
