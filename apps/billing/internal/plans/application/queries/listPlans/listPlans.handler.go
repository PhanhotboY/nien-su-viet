package aqueries

import (
	"context"

	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/listPlans/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/repository"
)

type ListPlansHandler interface {
	grpcTypes.GrpcHandler[*ListPlansQuery, adto.ListPlansResDto]
}

type listPlansHandler struct {
	logger        logger.Logger
	planDbRepo    drepo.PlanDbRepo
	planCacheRepo drepo.PlanCacheRepo
}

func NewListPlansHandler(
	logger logger.Logger,
	planDbRepo drepo.PlanDbRepo,
	planCacheRepo drepo.PlanCacheRepo,
) ListPlansHandler {
	return &listPlansHandler{
		logger:        logger,
		planDbRepo:    planDbRepo,
		planCacheRepo: planCacheRepo,
	}
}

func (h *listPlansHandler) Handle(ctx context.Context, query *ListPlansQuery) (adto.ListPlansResDto, error) {
	h.logger.Info("Handling ListPlansQuery", "query", query)

	plans, err := h.planDbRepo.GetPlans(ctx, query.ToMap())
	if err != nil {
		h.logger.Error("Failed to list plans", "error", err)
		return nil, err
	}

	h.logger.Info("ListPlansQuery handled successfully!")
	return adto.NewListPlansResDto(plans), nil
}
