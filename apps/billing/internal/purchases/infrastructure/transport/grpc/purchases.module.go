package tgrpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

var Module = fx.Module(
	"purchasesInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPurchasesGrpcServiceServer,
	),

	fx.Invoke(
		func(purchasesGrpcServer grpcServer.GrpcServer, purchasesGrpcServiceServer billing_service.PurchaseServiceServer) error {
			purchasesGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				billing_service.RegisterPurchaseServiceServer(server, purchasesGrpcServiceServer)
			})
			return nil
		},
	),
)
