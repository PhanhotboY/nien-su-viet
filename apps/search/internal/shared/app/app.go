package app

import (
	hevents "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/infrastructure"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	// configure dependencies
	appBuilder := fxapp.NewApplicationBuilder()

	// provide infrastructure dependencies, e.g., database, cache, etc.
	appBuilder.ProvideModule(infrastructure.Module)

	appBuilder.ProvideModule(posts.Module)
	appBuilder.ProvideModule(hevents.Module)

	app := appBuilder.Build()

	app.Logger().Info("Starting posts_service application")
	app.Run()
}
