package tgrpc

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	listPTByAttempt "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/queries/listPaymentTransactionsByAttempt"
)

type paymentTransactionsGrpcServiceServer struct {
	logger logger.Logger

	listPTByAttemptHandler listPTByAttempt.ListPaymentTransactionsByAttemptHandler
}

func NewPaymentTransactionsGrpcServiceServer(
	logger logger.Logger,
	listPTByAttemptHandler listPTByAttempt.ListPaymentTransactionsByAttemptHandler,
) billing_service.PaymentTransactionServiceServer {
	return &paymentTransactionsGrpcServiceServer{
		logger:                 logger,
		listPTByAttemptHandler: listPTByAttemptHandler,
	}
}

func (s *paymentTransactionsGrpcServiceServer) ListByPurchase(
	ctx context.Context,
	req *billing_service.ListPaymentTransactionsByPurchaseRequest,
) (*billing_service.ListPaymentTransactionsByPurchaseResponse, error) {
	query, err := listPTByAttempt.NewListPaymentTransactionsByAttemptQuery(req)
	if err != nil {
		s.logger.Errorf("Failed to create ListPaymentTransactionsByAttemptQuery: %v", err)
		return nil, err
	}

	res, err := s.listPTByAttemptHandler.Handle(ctx, query)
	if err != nil {
		s.logger.Errorf("Failed to handle ListPaymentTransactionsByAttemptQuery: %v", err)
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &billing_service.ListPaymentTransactionsByPurchaseResponse{}, s.logger)
}

func (s *paymentTransactionsGrpcServiceServer) ListByAttempt(
	ctx context.Context,
	req *billing_service.ListPaymentTransactionsByAttemptRequest,
) (*billing_service.ListPaymentTransactionsByAttemptResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListByAttempt not implemented")
}
