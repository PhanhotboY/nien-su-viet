package grpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	postsService "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/post_service"
	posts_service "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/post_service"
)

var Module = fx.Module(
	"postsInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPostsGrpcServiceServer,
	),

	// Register the gRPC server and its routes
	fx.Invoke(
		func(postsGrpcServer grpcServer.GrpcServer, postsServiceServer posts_service.PostsServiceServer) error {
			postsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				postsService.RegisterPostsServiceServer(server, postsServiceServer)
			})
			return nil
		},
	),
)
