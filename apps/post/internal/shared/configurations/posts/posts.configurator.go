package posts

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/configurations"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/configurations/posts/infrastructure"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp/contracts"
)

type PostsServiceConfigurator struct {
	contracts.Application
	infrastructureConfigurator *infrastructure.InfrastructureConfigurator
	postsModuleConfigurator    *configurations.PostsModuleConfigurator
}

func NewPostsServiceConfigurator(
	app contracts.Application,
) *PostsServiceConfigurator {
	infraConfigurator := infrastructure.NewInfrastructureConfigurator(app)
	postsModuleConfigurator := configurations.NewPostsModuleConfigurator(app)

	return &PostsServiceConfigurator{
		Application:                app,
		infrastructureConfigurator: infraConfigurator,
		postsModuleConfigurator:    postsModuleConfigurator,
	}
}

func (ic *PostsServiceConfigurator) ConfigurePosts() {
	// Shared
	// Infrastructure
	ic.infrastructureConfigurator.ConfigInfrastructures()

	// Shared
	// Posts service configurations

	// Modules
	// post module
	ic.postsModuleConfigurator.ConfigurePostsModule()
}

func (ic *PostsServiceConfigurator) MapPostsEndpoints() {
	// Shared
	// ic.ResolveFunc(
	// 	func(postsServer echocontracts.EchoHttpServer, cfg *config.Config) error {
	// 		postsServer.SetupDefaultMiddlewares()

	// 		// config posts root endpoint
	// 		postsServer.RouteBuilder().
	// 			RegisterRoutes(func(e *echo.Echo) {
	// 				e.GET("", func(ec echo.Context) error {
	// 					return ec.String(
	// 						http.StatusOK,
	// 						fmt.Sprintf(
	// 							"%s is running...",
	// 							cfg.AppOptions.GetMicroserviceNameUpper(),
	// 						),
	// 					)
	// 				})
	// 			})

	// 		// config posts swagger
	// 		ic.configSwagger(postsServer.RouteBuilder())

	// 		return nil
	// 	},
	// )

	// Modules
	// Posts Module endpoints
	ic.postsModuleConfigurator.MapPostsEndpoints()
}
