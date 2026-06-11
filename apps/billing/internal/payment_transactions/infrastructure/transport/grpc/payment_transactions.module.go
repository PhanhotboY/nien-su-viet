package tgrpc

import (
	"go.uber.org/fx"
	googleGrpc "google.golang.org/grpc"

	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

var Module = fx.Module(
	"paymentTransactionsInfrastructureTransportGrpcModule",

	fx.Provide(
		NewPaymentTransactionsGrpcServiceServer,
	),

	fx.Invoke(
		func(paymentTransactionsGrpcServer grpcServer.GrpcServer, paymentTransactionsGrpcServiceServer billing_service.PaymentTransactionServiceServer) error {
			paymentTransactionsGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				billing_service.RegisterPaymentTransactionServiceServer(server, paymentTransactionsGrpcServiceServer)
			})
			return nil
		},
	),
)
