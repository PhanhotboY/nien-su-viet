package grpc

import (
	"github.com/go-playground/validator"
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/metrics"
	postsService "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

var Module = fx.Module(
	"postsInfrastructureTransportGrpcModule",

	// fx.Provide(fx.Annotate(repositories.NewMongoPostReadRepository)),
	fx.Invoke(
		func(postsGrpcServer grpcServer.GrpcServer, postsMetrics *metrics.PostsMetrics, logger logger.Logger, validator *validator.Validate) error {
			postGrpcService := NewPostsGrpcServerHandler(logger, validator, postsMetrics)
			postsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				postsService.RegisterPostsServiceServer(server, postGrpcService)
			})
			return nil
		},
	),
)
