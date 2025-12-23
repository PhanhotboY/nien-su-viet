package initialize

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/controller/http"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/infrastructure/persistence"
	"gorm.io/gorm"
)

func InitFooterNavItem(db *gorm.DB) *http.FooterNavItemHandler {
	repo := persistence.NewFooterNavItemRepository(db)
	service := service.NewFooterNavItemService(repo)
	return http.NewFooterNavItemHandler(service)
}
