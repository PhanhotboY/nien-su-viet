package tgrpc

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	createPlanCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan"
	getPlanByIdQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById"
	listPlansQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/listPlans"
)

type plansGrpcServiceServer struct {
	logger logger.Logger

	createPlanCmd.CreatePlanHandler
	listPlansQuery.ListPlansHandler
	getPlanByIdQuery.GetPlanByIdHandler
}

func NewPlansGrpcServiceServer(
	logger logger.Logger,
	createPlanHandler createPlanCmd.CreatePlanHandler,
	listPlansHandler listPlansQuery.ListPlansHandler,
	getPlanByIdHandler getPlanByIdQuery.GetPlanByIdHandler,
) billing_service.PlanServiceServer {
	return &plansGrpcServiceServer{
		logger:             logger,
		CreatePlanHandler:  createPlanHandler,
		ListPlansHandler:   listPlansHandler,
		GetPlanByIdHandler: getPlanByIdHandler,
	}
}

func (s *plansGrpcServiceServer) CreatePlan(ctx context.Context, req *billing_service.CreatePlanRequest) (*billing_service.CreatePlanResponse, error) {
	typedReq, err := createPlanCmd.NewCreatePlanCommand(req)
	if err != nil {
		return nil, err
	}
	res, err := s.CreatePlanHandler.Handle(ctx, typedReq)
	if err != nil {
		return nil, err
	}

	return res.ToGrpcResponse(), nil
}

func (s *plansGrpcServiceServer) ListPlans(ctx context.Context, req *billing_service.ListPlansRequest) (*billing_service.ListPlansResponse, error) {
	query, err := listPlansQuery.NewListPlansQuery(req)
	if err != nil {
		return nil, err
	}
	res, err := s.ListPlansHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return res.ToGrpcResponse(), nil
}

func (s *plansGrpcServiceServer) UpdatePlan(context.Context, *billing_service.UpdatePlanRequest) (*billing_service.UpdatePlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdatePlan not implemented")
}

func (s *plansGrpcServiceServer) GetPlan(ctx context.Context, req *billing_service.GetPlanRequest) (*billing_service.GetPlanResponse, error) {
	query, err := getPlanByIdQuery.NewGetPlanByIdQuery(req)
	if err != nil {
		return nil, err
	}
	res, err := s.GetPlanByIdHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return res.ToGrpcResponse(), nil
}
