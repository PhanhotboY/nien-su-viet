package tgrpc

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	createPurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase"
	getPurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase"
)

type purchasesGrpcServiceServer struct {
	logger logger.Logger

	createPurchaseCmd.CreatePurchaseHandler
	getPurchaseCmd.GetPurchaseHandler
}

func NewPurchasesGrpcServiceServer(
	logger logger.Logger,
	createPurchaseHandler createPurchaseCmd.CreatePurchaseHandler,
	getPurchaseHandler getPurchaseCmd.GetPurchaseHandler,
) billing_service.PurchaseServiceServer {
	return &purchasesGrpcServiceServer{
		logger:                logger,
		CreatePurchaseHandler: createPurchaseHandler,
		GetPurchaseHandler:    getPurchaseHandler,
	}
}

func (s *purchasesGrpcServiceServer) CreatePurchase(ctx context.Context, req *billing_service.CreatePurchaseRequest) (*billing_service.CreatePurchaseResponse, error) {
	typedReq, err := createPurchaseCmd.NewCreatePurchaseCommand(req)
	if err != nil {
		return nil, err
	}
	res, err := s.CreatePurchaseHandler.Handle(ctx, typedReq)
	if err != nil {
		return nil, err
	}

	return res.ToGrpcResponse(), nil
}

func (s *purchasesGrpcServiceServer) GetPurchase(ctx context.Context, req *billing_service.GetPurchaseRequest) (*billing_service.GetPurchaseResponse, error) {
	typedReq, err := getPurchaseCmd.NewGetPurchaseQuery(req)
	if err != nil {
		return nil, err
	}
	res, err := s.GetPurchaseHandler.Handle(ctx, typedReq)
	if err != nil {
		return nil, err
	}

	return res.ToGrpcResponse(), nil
}

func (s *purchasesGrpcServiceServer) ListPurchasesBySubscription(context.Context, *billing_service.ListPurchasesBySubscriptionRequest) (*billing_service.ListPurchasesBySubscriptionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListPurchasesBySubscription not implemented")
}

func (s *purchasesGrpcServiceServer) ListPurchasesByUser(context.Context, *billing_service.ListPurchasesByUserRequest) (*billing_service.ListPurchasesByUserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListPurchasesByUser not implemented")
}
