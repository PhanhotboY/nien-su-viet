package tgrpc

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	createPaymentAttemptCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt"
	getPaymentAttemptQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt"
)

type paymentAttemptsGrpcServiceServer struct {
	logger logger.Logger

	createPaymentAttemptCmd.CreatePaymentAttemptHandler
	getPaymentAttemptQuery.GetPaymentAttemptHandler
}

func NewPaymentAttemptsGrpcServiceServer(
	logger logger.Logger,
	createPaymentAttemptHandler createPaymentAttemptCmd.CreatePaymentAttemptHandler,
	getPaymentAttemptHandler getPaymentAttemptQuery.GetPaymentAttemptHandler,
) billing_service.PaymentServiceServer {
	return &paymentAttemptsGrpcServiceServer{
		logger:                      logger,
		CreatePaymentAttemptHandler: createPaymentAttemptHandler,
		GetPaymentAttemptHandler:    getPaymentAttemptHandler,
	}
}

func (s *paymentAttemptsGrpcServiceServer) CreatePaymentAttempt(ctx context.Context, req *billing_service.CreatePaymentAttemptRequest) (*billing_service.CreatePaymentAttemptResponse, error) {
	typedReq, err := createPaymentAttemptCmd.NewCreatePaymentAttemptCommand(req)
	if err != nil {
		return nil, err
	}
	res, err := s.CreatePaymentAttemptHandler.Handle(ctx, typedReq)
	if err != nil {
		return nil, err
	}

	return res.ToGrpcResponse(), nil
}

func (s *paymentAttemptsGrpcServiceServer) GetPaymentAttempt(ctx context.Context, req *billing_service.GetPaymentAttemptRequest) (*billing_service.GetPaymentAttemptResponse, error) {
	typedReq, err := getPaymentAttemptQuery.NewGetPaymentAttemptQuery(req)
	if err != nil {
		return nil, err
	}
	res, err := s.GetPaymentAttemptHandler.Handle(ctx, typedReq)
	if err != nil {
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &billing_service.GetPaymentAttemptResponse{}, s.logger)
}

func (s *paymentAttemptsGrpcServiceServer) ListPaymentAttemptsByPurchase(context.Context, *billing_service.ListPaymentAttemptsByPurchaseRequest) (*billing_service.ListPaymentAttemptsByPurchaseResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListPaymentAttemptsByPurchase not implemented")
}

func (s *paymentAttemptsGrpcServiceServer) RefundPurchase(context.Context, *billing_service.RefundPurchaseRequest) (*billing_service.RefundPurchaseResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method RefundPurchase not implemented")
}
