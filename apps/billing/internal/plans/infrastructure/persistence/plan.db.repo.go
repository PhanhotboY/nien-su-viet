package prepo

import (
	"context"
	"slices"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

type planDbRepo struct {
	logger logger.Logger
	db     dbcontracts.TxContextDb
}

func NewPlanDbRepo(l logger.Logger, db dbcontracts.TxContextDb) drepo.PlanDbRepo {
	return &planDbRepo{l, db}
}

func (r *planDbRepo) CreatePlan(ctx context.Context, plan *entity.Plan) (string, error) {
	r.logger.Info("Creating plan", "plan", plan)

	if err := r.db.WithTxIfExists(ctx).DB().Create(plan).Error; err != nil {
		r.logger.Error("Failed to create plan", "error", err)
		return "", err
	}

	r.logger.Info("Plan created successfully", "planId", plan.ID)
	return plan.ID.String(), nil
}

func (r *planDbRepo) UpdatePlan(ctx context.Context, planId string, updates map[string]any) (string, error) {
	r.logger.Info("Updating plan", "planId", planId, "updates", updates)

	var existingPlan entity.Plan
	if err := r.db.WithTxIfExists(ctx).DB().First(&existingPlan, "id = ?", planId).Error; err != nil {
		r.logger.Error("Failed to find plan for update", "planId", planId, "error", err)
		return "", err
	}

	if err := r.db.WithTxIfExists(ctx).DB().Model(&existingPlan).Updates(updates).Error; err != nil {
		r.logger.Error("Failed to update plan", "planId", planId, "error", err)
		return "", err
	}

	r.logger.Info("Plan updated successfully", "planId", planId)
	return planId, nil
}

func (r *planDbRepo) DeletePlan(ctx context.Context, planId string) (string, error) {
	r.logger.Info("Deleting plan", "planId", planId)

	if err := r.db.WithTxIfExists(ctx).DB().Delete(&entity.Plan{}, "id = ?", planId).Error; err != nil {
		r.logger.Error("Failed to delete plan", "planId", planId, "error", err)
		return "", err
	}

	r.logger.Info("Plan deleted successfully", "planId", planId)
	return planId, nil
}

func (r *planDbRepo) GetPlanById(ctx context.Context, planId string) (*entity.Plan, error) {
	r.logger.Info("Getting plan", "planId", planId)

	plan := new(entity.Plan)
	if err := r.db.WithTxIfExists(ctx).DB().First(plan, "id = ?", planId).Error; err != nil {
		r.logger.Error("Failed to get plan", "planId", planId, "error", err)
		return nil, err
	}

	r.logger.Info("Plan retrieved successfully %+v", plan)
	return plan, nil
}

func (r *planDbRepo) GetPlans(ctx context.Context, filter map[string]any) ([]*entity.Plan, error) {
	r.logger.Info("Getting plans with filter", "filter", filter)

	var plans []*entity.Plan
	query := r.db.WithTxIfExists(ctx).DB().Model(&entity.Plan{})
	for key, value := range filter {
		if !slices.Contains([]string{"limit", "page"}, key) {
			query = query.Where(key+" = ?", value)
		}
	}
	query = query.Find(&plans)
	if limit, ok := filter["limit"].(int); ok {
		query = query.Limit(limit)
	}
	if page, ok := filter["page"].(int); ok && page > 0 {
		offset := (page - 1) * filter["limit"].(int)
		query = query.Offset(offset)
	}

	if err := query.Error; err != nil {
		r.logger.Error("Failed to get plans with filter", "filter", filter, "error", err)
		return nil, err
	}

	r.logger.Info("Plans retrieved successfully", "planCount", len(plans))
	return plans, nil
}
