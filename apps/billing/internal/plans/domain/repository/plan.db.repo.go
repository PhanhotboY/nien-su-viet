package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
)

type PlanDbRepo interface {
	CreatePlan(ctx context.Context, plan *entity.Plan) (string, error)
	UpdatePlan(ctx context.Context, planId string, updates map[string]any) (string, error)
	DeletePlan(ctx context.Context, planId string) (string, error)
	GetPlanById(ctx context.Context, planId string) (*entity.Plan, error)
	GetPlans(ctx context.Context, filter map[string]any) ([]*entity.Plan, error)
}
