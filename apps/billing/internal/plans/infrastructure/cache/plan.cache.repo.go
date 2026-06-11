package crepo

import (
	"context"
	"reflect"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
)

const (
	PLAN_CACHE_PREFIX = "billing_service.Plan"
)

type planCacheRepo struct {
	logger logger.Logger
	redis  redis.RedisClientWithExpire
}

func NewPlanCacheRepo(
	logger logger.Logger,
	redis redis.RedisClientWithExpire,
) drepo.PlanCacheRepo {
	return &planCacheRepo{
		logger: logger,
		redis:  redis,
	}
}

func (r *planCacheRepo) GetPlan(ctx context.Context, key string) (*entity.Plan, error) {
	r.logger.Debug("Getting plan from cache", "key", key)
	var res = new(entity.Plan)
	if err := r.redis.HGet(ctx, PLAN_CACHE_PREFIX, key, res); err != nil {
		return nil, err
	}
	if reflect.DeepEqual(*res, entity.Plan{}) {
		return nil, nil
	}
	return res, nil
}

func (r *planCacheRepo) PutPlan(ctx context.Context, key string, plan *entity.Plan) error {
	r.logger.Debug("Putting plan in cache", "key", key)
	return r.redis.HSet(ctx, PLAN_CACHE_PREFIX, key, plan)
}

func (r *planCacheRepo) DeletePlan(ctx context.Context, key string) error {
	r.logger.Debug("Deleting plan from cache", "key", key)
	if _, err := r.redis.HDel(ctx, PLAN_CACHE_PREFIX, key); err != nil {
		return err
	}
	return nil
}

func (r *planCacheRepo) DeleteAllPlans(ctx context.Context) error {
	r.logger.Debug("Deleting all plans from cache")
	if _, err := r.redis.MDel(ctx, PLAN_CACHE_PREFIX); err != nil {
		return err
	}
	return nil
}
