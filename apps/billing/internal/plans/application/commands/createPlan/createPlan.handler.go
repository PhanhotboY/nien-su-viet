package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type CreatePlanHandler interface {
	grpcTypes.GrpcHandler[*CreatePlanCmd, adto.CreatePlanResDto]
}

type createPlanHandler struct {
	logger        logger.Logger
	planDbRepo    drepo.PlanDbRepo
	planCacheRepo drepo.PlanCacheRepo
}

func NewCreatePlanHandler(l logger.Logger, planDbRepo drepo.PlanDbRepo, planCacheRepo drepo.PlanCacheRepo) CreatePlanHandler {
	return &createPlanHandler{l, planDbRepo, planCacheRepo}
}

func (h *createPlanHandler) Handle(ctx context.Context, cmd *CreatePlanCmd) (adto.CreatePlanResDto, error) {
	planId, err := h.planDbRepo.CreatePlan(ctx, cmd.MapToEntity())
	if err != nil {
		h.logger.Error("failed to create plan", "error", err)
		return nil, grpcerrors.NewInternalServerGrpcError("failed to create plan", "NewCreatePlanHandler.Handle")
	}

	if err = h.planCacheRepo.DeleteAllPlans(ctx); err != nil {
		h.logger.Error("failed to delete all plans in cache", "error", err)
	}

	return adto.NewCreatePlanResDto(planId, true, "Plan created successfully"), nil
}
