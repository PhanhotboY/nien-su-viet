package queries

import (
	"context"

	postDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
)

type GetPublishedPostsHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IGetPublishedPostsHandler interface {
	grpcTypes.GrpcHandler[*GetPublishedPostsQuery, *dto.GetPublishedPostsRes]
}

func NewGetPublishedPostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) GetPublishedPostsHandler {
	return GetPublishedPostsHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (c GetPublishedPostsHandler) Handle(
	ctx context.Context,
	query *GetPublishedPostsQuery,
) (*dto.GetPublishedPostsRes, error) {
	var cached *utils.PaginatedResponse[entity.PostBrief]

	queryEntity := query.MapToQuery()
	queryKey, err := jsonUtils.MarshalToJsonString(queryEntity)
	cached, err = c.cacheRepo.GetPosts(ctx, queryKey)
	if err != nil {
		c.log.Warnf("failed to get published posts from cache: %v, fallback to db", err)
	}

	if cached == nil {
		posts, err := c.postRepo.GetPosts(ctx, queryEntity)
		if err != nil {
			c.log.Errorf("failed to get published posts: %v", err)
			return nil, grpcerrors.ParseError(err)
		}
		total, err := c.postRepo.CountPosts(ctx, queryEntity)
		if err != nil {
			c.log.Errorf("failed to count published posts: %v", err)
			return nil, grpcerrors.ParseError(err)
		}

		cached = &utils.PaginatedResponse[entity.PostBrief]{
			Data:       posts,
			Pagination: utils.NewPagination(queryEntity.Limit, queryEntity.Offset, total),
		}

		if err := c.cacheRepo.PutPosts(ctx, queryKey, cached); err != nil {
			c.log.Warnf("failed to cache published posts: %v", err)
		}
	}

	postBriefs := make([]postDto.PostBriefDto, 0)
	for _, post := range cached.Data {
		postBrief := postDto.PostBriefDto{}
		postBrief.FromEntity(&post)
		postBriefs = append(postBriefs, postBrief)
	}

	return dto.NewGetPublishedPostsRes(
		postBriefs,
		*cached.Pagination,
	), nil
}
