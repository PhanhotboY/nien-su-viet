package service

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/dto"
)

type HeaderNavItemService interface {
	GetHeaderNavItems(ctx context.Context) ([]*dto.HeaderNavItemData, error)
	CreateHeaderNavItem(ctx context.Context, item *dto.HeaderNavItemCreateReq) error
	UpdateHeaderNavItem(ctx context.Context, id string, item *dto.HeaderNavItemUpdateReq) error
	DeleteHeaderNavItem(ctx context.Context, id string) error
}
