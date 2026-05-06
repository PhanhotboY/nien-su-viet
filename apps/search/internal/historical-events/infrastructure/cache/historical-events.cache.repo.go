package cache

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	redis2 "github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
	"github.com/redis/go-redis/v9"
)

const (
	heventCachePrefixKey = "hevent_service"
)

type heventCacheRepository struct {
	log         logger.Logger
	redisClient redis2.RedisClientWithExpire
}

func NewHistoricalEventCacheRepository(
	log logger.Logger,
	redisClient redis2.RedisClientWithExpire,
) repository.HistoricalEventCacheRepository {
	return &heventCacheRepository{
		log:         log,
		redisClient: redisClient,
	}
}

func (r *heventCacheRepository) PutHistoricalEvent(
	ctx context.Context,
	key string,
	HistoricalEvent *entity.HistoricalEvent,
) error {
	if err := r.redisClient.HSet(ctx, r.getRedisHistoricalEventPrefixKey(), key, HistoricalEvent); err != nil {
		r.log.Errorw(
			fmt.Sprintf(
				"error in updating HistoricalEvent with key %s",
				key,
			),
			logger.Fields{
				"Id":        HistoricalEvent.Id,
				"Key":       key,
				"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	r.log.Infow(
		fmt.Sprintf(
			"HistoricalEvent with key '%s', prefix '%s' updated successfully",
			key,
			r.getRedisHistoricalEventPrefixKey(),
		),
		logger.Fields{
			"Id":        HistoricalEvent.Id,
			"Key":       key,
			"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
		},
	)

	return nil
}

func (r *heventCacheRepository) PutHistoricalEvents(
	ctx context.Context,
	key string,
	HistoricalEvents *utils.PaginatedResponse[entity.HistoricalEventBrief],
) error {

	if err := r.redisClient.HSet(ctx, r.getRedisHistoricalEventPrefixKey(), key, HistoricalEvents); err != nil {
		r.log.Errorw(
			fmt.Sprintf(
				"error in updating HistoricalEvents with key %s",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	r.log.Infow(
		fmt.Sprintf(
			"HistoricalEvents with key '%s', prefix '%s' updated successfully",
			key,
			r.getRedisHistoricalEventPrefixKey(),
		),
		logger.Fields{
			"Key":       key,
			"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
		},
	)

	return nil
}

func (r *heventCacheRepository) GetHistoricalEvent(
	ctx context.Context,
	key string,
) (*entity.HistoricalEvent, error) {

	HistoricalEvent := new(entity.HistoricalEvent)
	err := r.redisClient.HGet(ctx, r.getRedisHistoricalEventPrefixKey(), key, HistoricalEvent)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		r.log.Errorw(
			fmt.Sprintf(
				"error in getting HistoricalEvent with Key %s from database",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return nil, err
	}
	if HistoricalEvent.Id == "" {
		return nil, nil
	}

	r.log.Infow(
		fmt.Sprintf(
			"HistoricalEvent with key '%s', prefix '%s' loaded",
			key,
			r.getRedisHistoricalEventPrefixKey(),
		),
		logger.Fields{
			"Id":        HistoricalEvent.Id,
			"Key":       key,
			"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
		},
	)

	return HistoricalEvent, nil
}

func (r *heventCacheRepository) GetHistoricalEvents(
	ctx context.Context,
	key string,
) (*utils.PaginatedResponse[entity.HistoricalEventBrief], error) {
	var hevents *utils.PaginatedResponse[entity.HistoricalEventBrief]
	err := r.redisClient.HGet(ctx, r.getRedisHistoricalEventPrefixKey(), key, &hevents)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		r.log.Errorw(
			fmt.Sprintf(
				"error in getting HistoricalEvents with Key %s from database",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return nil, err
	}

	r.log.Infow(
		fmt.Sprintf(
			"HistoricalEvents with with key '%s', prefix '%s' loaded",
			key,
			r.getRedisHistoricalEventPrefixKey(),
		),
		logger.Fields{
			"Key":       key,
			"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
		},
	)

	return hevents, nil
}

func (r *heventCacheRepository) DeleteHistoricalEvent(
	ctx context.Context,
	key string,
) error {
	if _, err := r.redisClient.HDel(ctx, r.getRedisHistoricalEventPrefixKey(), key); err != nil {
		r.log.Errorw(
			fmt.Sprintf(
				"error in deleting HistoricalEvent with key %s",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	r.log.Infow(
		fmt.Sprintf(
			"HistoricalEvents with key %s, prefix: %s deleted successfully",
			key,
			r.getRedisHistoricalEventPrefixKey(),
		),
		logger.Fields{"Key": key, "PrefixKey": r.getRedisHistoricalEventPrefixKey()},
	)

	return nil
}

func (r *heventCacheRepository) DeleteAllHistoricalEvents(ctx context.Context) error {
	if _, err := r.redisClient.Del(ctx, r.getRedisHistoricalEventPrefixKey()); err != nil {
		r.log.Errorw(
			"error in deleting all HistoricalEvents",
			logger.Fields{
				"PrefixKey": r.getRedisHistoricalEventPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	r.log.Infow(
		"all HistoricalEvents deleted",
		logger.Fields{"PrefixKey": r.getRedisHistoricalEventPrefixKey()},
	)

	return nil
}

func (r *heventCacheRepository) getRedisHistoricalEventPrefixKey() string {
	return heventCachePrefixKey
}
