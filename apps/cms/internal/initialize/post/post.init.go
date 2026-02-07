package initialize

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/controller/http"
	persistence "github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/infrastructure/persistence"
	"gorm.io/gorm"
)

func InitPost(db *gorm.DB) *http.PostHandler {
	postRepo := persistence.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	return http.NewPostHandler(postService)
}
