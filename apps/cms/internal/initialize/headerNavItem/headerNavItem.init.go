package initialize

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/controller/http"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/infrastructure/persistence"
	"gorm.io/gorm"
)

func InitHeaderNavItem(db *gorm.DB) *http.HeaderNavItemHandler {
	repo := persistence.NewHeaderNavItemRepository(db)
	service := service.NewHeaderNavItemService(repo)
	return http.NewHeaderNavItemHandler(service)
}
