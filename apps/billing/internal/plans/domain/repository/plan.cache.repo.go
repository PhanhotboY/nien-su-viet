package drepo

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
)

type PlanCacheRepo interface {
	PutPlan(ctx context.Context, key string, Plan *entity.Plan) error
	GetPlan(ctx context.Context, key string) (*entity.Plan, error)
	DeletePlan(ctx context.Context, key string) error
	DeleteAllPlans(ctx context.Context) error
}
