package grpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	searchService "github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/grpc/genproto"
	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
)

var Module = fx.Module(
	"historicalEventsInfrastructureTransportGrpcModule",

	fx.Provide(
		NewHistoricalEventsGrpcServerHandler,
	),

	// Register the gRPC server and its routes
	fx.Invoke(
		func(grpcServer grpcServer.GrpcServer, historicalEventsGrpcServiceHandlers searchService.HistoricalEventServiceServer) error {
			grpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				searchService.RegisterHistoricalEventServiceServer(server, historicalEventsGrpcServiceHandlers)
			})
			return nil
		},
	),
)
