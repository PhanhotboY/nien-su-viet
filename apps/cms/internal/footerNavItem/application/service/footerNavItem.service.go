package service

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/dto"
)

type FooterNavItemService interface {
	GetFooterNavItems(ctx context.Context) ([]*dto.FooterNavItemData, error)
	CreateFooterNavItem(ctx context.Context, item *dto.FooterNavItemCreateReq) (string, error)
	UpdateFooterNavItem(ctx context.Context, id string, item *dto.FooterNavItemUpdateReq) (string, error)
	DeleteFooterNavItem(ctx context.Context, id string) (string, error)
}
