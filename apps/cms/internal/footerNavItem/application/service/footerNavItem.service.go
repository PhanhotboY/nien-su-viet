package service

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type FooterNavItemService interface {
	GetFooterNavItems(ctx context.Context) ([]*dto.FooterNavItemData, error)
	CreateFooterNavItem(ctx context.Context, item *dto.FooterNavItemCreateReq) (*response.OperationResult, error)
	UpdateFooterNavItem(ctx context.Context, id string, item *dto.FooterNavItemUpdateReq) (*response.OperationResult, error)
	DeleteFooterNavItem(ctx context.Context, id string) (*response.OperationResult, error)
}
