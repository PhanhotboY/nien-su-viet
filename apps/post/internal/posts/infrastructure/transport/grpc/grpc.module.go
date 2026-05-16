package grpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	postsService "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/post_service"
)

var Module = fx.Module(
	"postsInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPostsGrpcServerHandler,
	),

	// Register the gRPC server and its routes
	fx.Invoke(
		func(postsGrpcServer grpcServer.GrpcServer, postGrpcServiceHandlers *PostsGrpcServerHandler) error {
			postsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				postsService.RegisterPostsServiceServer(server, postGrpcServiceHandlers)
			})
			return nil
		},
	),
)
