package service

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type HeaderNavItemService interface {
	GetHeaderNavItems(ctx context.Context) ([]*dto.HeaderNavItemData, error)
	UpdateHeaderNavItem(ctx context.Context, id string, item *dto.HeaderNavItemUpdateReq) (*response.OperationResult, error)
	CreateHeaderNavItem(ctx context.Context, item *dto.HeaderNavItemCreateReq) (*response.OperationResult, error)
	DeleteHeaderNavItem(ctx context.Context, id string) (*response.OperationResult, error)
}
