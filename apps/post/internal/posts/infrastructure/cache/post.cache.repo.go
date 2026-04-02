package cache

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	redis2 "github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
	"github.com/redis/go-redis/v9"
)

const (
	postCachePrefixKey = "post_service"
)

type postCacheRepository struct {
	log         logger.Logger
	redisClient redis2.RedisClientWithExpire
	// tracer      tracing.AppTracer
}

func NewPostCacheRepository(
	log logger.Logger,
	redisClient redis2.RedisClientWithExpire,
	// tracer tracing.AppTracer,
) repository.PostCacheRepository {
	return &postCacheRepository{
		log:         log,
		redisClient: redisClient,
		// tracer:      tracer,
	}
}

func (r *postCacheRepository) PutPost(
	ctx context.Context,
	key string,
	Post *entity.Post,
) error {
	// ctx, span := r.tracer.Start(ctx, "redisRepository.PutPost")
	// span.SetAttributes(
	// 	attribute2.String("PrefixKey", r.getRedisPostPrefixKey()),
	// )
	// span.SetAttributes(attribute2.String("Key", key))
	// defer span.End()

	if err := r.redisClient.HSet(ctx, r.getRedisPostPrefixKey(), key, Post); err != nil {
		// 	return utils.TraceErrStatusFromSpan(
		// 		span,
		// 		errors.WrapIf(
		// 			err,
		// 			fmt.Sprintf(
		// 				"error in updating Post with key %s",
		// 				key,
		// 			),
		// 		),
		// 	)
		r.log.Errorw(
			fmt.Sprintf(
				"error in updating Post with key %s",
				key,
			),
			logger.Fields{
				"Id":        Post.Id.String(),
				"Key":       key,
				"PrefixKey": r.getRedisPostPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	// span.SetAttributes(attribute.Object("Post", Post))

	r.log.Infow(
		fmt.Sprintf(
			"Post with key '%s', prefix '%s' updated successfully",
			key,
			r.getRedisPostPrefixKey(),
		),
		logger.Fields{
			"Id":        Post.Id.String(),
			"Key":       key,
			"PrefixKey": r.getRedisPostPrefixKey(),
		},
	)

	return nil
}

func (r *postCacheRepository) PutPosts(
	ctx context.Context,
	key string,
	Posts *utils.PaginatedResponse[entity.PostBrief],
) error {
	// ctx, span := r.tracer.Start(ctx, "redisRepository.PutPosts")
	// span.SetAttributes(
	// 	attribute2.String("PrefixKey", r.getRedisPostPrefixKey()),
	// )
	// span.SetAttributes(attribute2.String("Key", key))
	// defer span.End()

	if err := r.redisClient.HSet(ctx, r.getRedisPostPrefixKey(), key, Posts); err != nil {
		r.log.Errorw(
			fmt.Sprintf(
				"error in updating Posts with key %s",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisPostPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	// span.SetAttributes(attribute.Object("Posts", Posts))

	r.log.Infow(
		fmt.Sprintf(
			"Posts with key '%s', prefix '%s' updated successfully",
			key,
			r.getRedisPostPrefixKey(),
		),
		logger.Fields{
			"Key":       key,
			"PrefixKey": r.getRedisPostPrefixKey(),
		},
	)

	return nil
}

func (r *postCacheRepository) GetPost(
	ctx context.Context,
	key string,
) (*entity.Post, error) {
	// ctx, span := r.tracer.Start(ctx, "redisRepository.GetPostById")
	// span.SetAttributes(
	// 	attribute2.String("PrefixKey", r.getRedisPostPrefixKey()),
	// )
	// span.SetAttributes(attribute2.String("Key", key))
	// defer span.End()

	Post := new(entity.Post)
	err := r.redisClient.HGet(ctx, r.getRedisPostPrefixKey(), key, Post)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		// return nil, utils.TraceErrStatusFromSpan(
		// 	span,
		// 	errors.WrapIf(
		// 		err,
		// 		fmt.Sprintf(
		// 			"error in getting Post with Key %s from database",
		// 			key,
		// 		),
		// 	),
		// )
		r.log.Errorw(
			fmt.Sprintf(
				"error in getting Post with Key %s from database",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisPostPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return nil, err
	}
	if Post.Id == uuid.Nil {
		return nil, nil
	}

	// span.SetAttributes(attribute.Object("Post", Post))

	r.log.Infow(
		fmt.Sprintf(
			"Post with key '%s', prefix '%s' laoded",
			key,
			r.getRedisPostPrefixKey(),
		),
		logger.Fields{
			"Id":        Post.Id.String(),
			"Key":       key,
			"PrefixKey": r.getRedisPostPrefixKey(),
		},
	)

	return Post, nil
}

func (r *postCacheRepository) GetPosts(
	ctx context.Context,
	key string,
) (*utils.PaginatedResponse[entity.PostBrief], error) {
	// ctx, span := r.tracer.Start(ctx, "redisRepository.GetPosts")
	// span.SetAttributes(
	// 	attribute2.String("PrefixKey", r.getRedisPostPrefixKey()),
	// )
	// span.SetAttributes(attribute2.String("Key", key))
	// defer span.End()

	var posts *utils.PaginatedResponse[entity.PostBrief]
	err := r.redisClient.HGet(ctx, r.getRedisPostPrefixKey(), key, &posts)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		// return nil, utils.TraceErrStatusFromSpan(
		// 	span,
		// 	errors.WrapIf(
		// 		err,
		// 		fmt.Sprintf(
		// 			"error in getting Posts with Key %s from database",
		// 			key,
		// 		),
		// 	),
		// )
		r.log.Errorw(
			fmt.Sprintf(
				"error in getting Posts with Key %s from database",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisPostPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return nil, err
	}

	// span.SetAttributes(attribute.Object("Posts", posts))

	r.log.Infow(
		fmt.Sprintf(
			"Posts with with key '%s', prefix '%s' laoded",
			key,
			r.getRedisPostPrefixKey(),
		),
		logger.Fields{
			"Key":       key,
			"PrefixKey": r.getRedisPostPrefixKey(),
		},
	)

	return posts, nil
}

func (r *postCacheRepository) DeletePost(
	ctx context.Context,
	key string,
) error {
	// ctx, span := r.tracer.Start(ctx, "redisRepository.DeletePost")
	// span.SetAttributes(
	// 	attribute2.String("PrefixKey", r.getRedisPostPrefixKey()),
	// )
	// span.SetAttributes(attribute2.String("Key", key))
	// defer span.End()

	if _, err := r.redisClient.HDel(ctx, r.getRedisPostPrefixKey(), key); err != nil {
		// return utils.TraceErrStatusFromSpan(
		// 	span,
		// 	errors.WrapIf(
		// 		err,
		// 		fmt.Sprintf(
		// 			"error in deleting Post with key %s",
		// 			key,
		// 		),
		// 	),
		// )
		r.log.Errorw(
			fmt.Sprintf(
				"error in deleting Post with key %s",
				key,
			),
			logger.Fields{
				"Key":       key,
				"PrefixKey": r.getRedisPostPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	r.log.Infow(
		fmt.Sprintf(
			"Posts with key %s, prefix: %s deleted successfully",
			key,
			r.getRedisPostPrefixKey(),
		),
		logger.Fields{"Key": key, "PrefixKey": r.getRedisPostPrefixKey()},
	)

	return nil
}

func (r *postCacheRepository) DeleteAllPosts(ctx context.Context) error {
	// ctx, span := r.tracer.Start(ctx, "redisRepository.DeleteAllPosts")
	// span.SetAttributes(
	// 	attribute2.String("PrefixKey", r.getRedisPostPrefixKey()),
	// )
	// defer span.End()

	if _, err := r.redisClient.Del(ctx, r.getRedisPostPrefixKey()); err != nil {
		// return utils.TraceErrStatusFromSpan(
		// 	span,
		// 	errors.WrapIf(
		// 		err,
		// 		"error in deleting all Posts",
		// 	),
		// )
		r.log.Errorw(
			"error in deleting all Posts",
			logger.Fields{
				"PrefixKey": r.getRedisPostPrefixKey(),
				"Error":     err.Error(),
			},
		)
		return err
	}

	r.log.Infow(
		"all Posts deleted",
		logger.Fields{"PrefixKey": r.getRedisPostPrefixKey()},
	)

	return nil
}

func (r *postCacheRepository) getRedisPostPrefixKey() string {
	return postCachePrefixKey
}
