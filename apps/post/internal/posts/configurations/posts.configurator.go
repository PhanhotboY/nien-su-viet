package configurations

import (
	"github.com/go-playground/validator"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc"
	posts_service "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	contracts2 "github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp/contracts"
	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	googleGrpc "google.golang.org/grpc"
)

type PostsModuleConfigurator struct {
	contracts2.Application
}

func NewPostsModuleConfigurator(
	app contracts2.Application,
) *PostsModuleConfigurator {
	return &PostsModuleConfigurator{
		Application: app,
	}
}

func (c *PostsModuleConfigurator) ConfigurePostsModule() {
	c.ResolveFunc(
		func(logger logger.Logger,
		// server echocontracts.EchoHttpServer,
		// postRepository repositories.PostMongoRepository,
		// postAggregateStore store.AggregateStore[*aggregate.Post],
		// tracer tracing.AppTracer,
		) error {
			// config Posts Mappings
			// err := mappings.ConfigurePostsMappings()
			// if err != nil {
			// 	return err
			// }

			// // config Posts Mediators
			// err = mediatr.ConfigPostsMediator(logger, postRepository, postAggregateStore, tracer)
			// if err != nil {
			// 	return err
			// }

			return nil
		},
	)
}

func (c *PostsModuleConfigurator) MapPostsEndpoints() {
	// config Posts Http Endpoints
	// c.ResolveFuncWithParamTag(func(endpoints []route.Endpoint) {
	// 	for _, endpoint := range endpoints {
	// 		endpoint.MapEndpoint()
	// 	}
	// }, `group:"post-routes"`,
	// )

	// config Posts Grpc Endpoints
	c.ResolveFunc(
		func(postsGrpcServer grpcServer.GrpcServer, logger logger.Logger, validator *validator.Validate) error {
			postGrpcService := grpc.NewPostGrpcServiceServer(logger, validator)
			postsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				posts_service.RegisterPostsServiceServer(server, postGrpcService)
			})
			return nil
		},
	)
}
