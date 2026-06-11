package tgrpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

var Module = fx.Module(
	"paymentAttemptsInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPaymentAttemptsGrpcServiceServer,
	),

	fx.Invoke(
		func(grpcSrv grpcServer.GrpcServer, paymentServiceServer billing_service.PaymentServiceServer) error {
			grpcSrv.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				billing_service.RegisterPaymentServiceServer(server, paymentServiceServer)
			})
			return nil
		},
	),
)
