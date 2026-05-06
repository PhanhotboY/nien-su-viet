package queries

import (
	"context"

	postDto "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPopularPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
)

type GetPopularPostsHandler struct {
	log       logger.Logger
	postRepo  repository.PostEsRepository
	cacheRepo repository.PostCacheRepository
}

type IGetPopularPostsHandler interface {
	grpcTypes.GrpcHandler[*GetPopularPostsQuery, *dto.GetPopularPostsRes]
}

func NewGetPopularPostsHandler(
	log logger.Logger,
	postRepo repository.PostEsRepository,
	cacheRepo repository.PostCacheRepository,
) GetPopularPostsHandler {
	return GetPopularPostsHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (c GetPopularPostsHandler) Handle(
	ctx context.Context,
	query *GetPopularPostsQuery,
) (*dto.GetPopularPostsRes, error) {
	var cached *utils.PaginatedResponse[entity.PostBrief]
	var page, limit uint32 = 1, 10
	if query.Page != nil {
		page = *query.Page
	}
	if query.Limit != nil {
		limit = *query.Limit
	}

	queryEntity := query.MapToQuery()
	queryKey, err := jsonUtils.MarshalToJsonString(queryEntity)
	c.log.Warnf("getting cached popular posts with query: %s", queryKey)
	cached, err = c.cacheRepo.GetPosts(ctx, queryKey)
	if err != nil {
		c.log.Warnf("failed to get popular posts from cache: %v, fallback to db", err)
	}

	if cached == nil {
		res, err := c.postRepo.SearchPosts(ctx, queryEntity)
		if err != nil {
			c.log.Errorf("failed to get popular posts: %v", err)
			return nil, grpcerrors.ParseError(err)
		}

		cached = &utils.PaginatedResponse[entity.PostBrief]{
			Data:       res.Posts,
			Pagination: utils.NewPagination(limit, page, uint32(res.TotalCount)),
		}

		if err := c.cacheRepo.PutPosts(ctx, queryKey, cached); err != nil {
			c.log.Warnf("failed to cache popular posts: %v", err)
		}
	}

	postBriefs := make([]postDto.PostBriefDto, 0)
	for _, post := range cached.Data {
		postBrief := postDto.PostBriefDto{}
		postBrief.FromEntity(&post)
		postBriefs = append(postBriefs, postBrief)
	}

	return dto.NewGetPopularPostsRes(
		postBriefs,
		*cached.Pagination,
	), nil
}
