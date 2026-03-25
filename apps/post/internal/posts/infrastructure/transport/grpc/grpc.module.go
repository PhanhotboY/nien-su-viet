package grpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	postsService "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
)

var Module = fx.Module(
	"postsInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPostsGrpcServerHandler,
	),
	// fx.Provide(fx.Annotate(repositories.NewMongoPostReadRepository)),
	fx.Invoke(
		func(postsGrpcServer grpcServer.GrpcServer, postGrpcService *PostsGrpcServerHandler) error {
			postsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				postsService.RegisterPostsServiceServer(server, postGrpcService)
			})
			return nil
		},
	),
)
