package tgrpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

var Module = fx.Module(
	"plansInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPlansGrpcServiceServer,
	),

	fx.Invoke(
		func(plansGrpcServer grpcServer.GrpcServer, plansGrpcServiceServer billing_service.PlanServiceServer) error {
			plansGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				billing_service.RegisterPlanServiceServer(server, plansGrpcServiceServer)
			})
			return nil
		},
	),
)
