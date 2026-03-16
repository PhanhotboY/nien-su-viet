package app

import "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/configurations/posts"

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	// configure dependencies
	appBuilder := NewPostsApplicationBuilder()
	appBuilder.ProvideModule(posts.PostsServiceModule)

	app := appBuilder.Build()

	// configure application
	app.ConfigurePosts()

	app.MapPostsEndpoints()

	app.Logger().Info("Starting posts_service application")
	app.Run()
}
