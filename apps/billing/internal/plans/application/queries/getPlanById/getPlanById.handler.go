package aqueries

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPlanByIdHandler interface {
	grpcTypes.GrpcHandler[*GetPlanByIdQuery, adto.GetPlanByIdResDto]
}

type getPlanByIdHandler struct {
	logger        logger.Logger
	planDbRepo    drepo.PlanDbRepo
	planCacheRepo drepo.PlanCacheRepo
}

func NewGetPlanByIdHandler(
	logger logger.Logger,
	planDbRepo drepo.PlanDbRepo,
	planCacheRepo drepo.PlanCacheRepo,
) GetPlanByIdHandler {
	return &getPlanByIdHandler{
		logger:        logger,
		planDbRepo:    planDbRepo,
		planCacheRepo: planCacheRepo,
	}
}

func (h *getPlanByIdHandler) Handle(ctx context.Context, query *GetPlanByIdQuery) (adto.GetPlanByIdResDto, error) {
	h.logger.Info("Handling GetPlanByIdQuery", "query", query)

	// Try to get from cache
	plan, err := h.planCacheRepo.GetPlan(ctx, query.Id)
	if err == nil && plan != nil {
		h.logger.Info("Plan retrieved from cache", "id", query.Id)
		return adto.NewGetPlanByIdResDto(plan), nil
	}

	// Fallback to DB
	plan, err = h.planDbRepo.GetPlanById(ctx, query.Id)
	if err != nil {
		h.logger.Error("Failed to get plan from database", "id", query.Id, "error", err)
		return nil, err
	}
	if plan == nil {
		h.logger.Warn("Plan not found", "id", query.Id)
		return nil, grpcerrors.NewNotFoundErrorGrpcError("Plan not found", "GetPlanByIdHandler")
	}

	// Cache the result
	if err := h.planCacheRepo.PutPlan(ctx, query.Id, plan); err != nil {
		h.logger.Warn("Failed to cache plan", "id", query.Id, "error", err)
	}

	h.logger.Info("GetPlanByIdQuery handled successfully!", "id", query.Id)
	return adto.NewGetPlanByIdResDto(plan), nil
}
