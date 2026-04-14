package queries

import (
	"context"

	postDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getAllPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
)

type GetAllPostsHandler struct {
	log       logger.Logger
	postRepo  repository.PostRepository
	cacheRepo repository.PostCacheRepository
}

type IGetAllPostsHandler interface {
	grpcTypes.GrpcHandler[*GetAllPostsQuery, *dto.GetAllPostsRes]
}

func NewGetAllPostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
	cacheRepo repository.PostCacheRepository,
) GetAllPostsHandler {
	return GetAllPostsHandler{
		log:       log,
		postRepo:  postRepo,
		cacheRepo: cacheRepo,
	}
}

func (c GetAllPostsHandler) Handle(
	ctx context.Context,
	query *GetAllPostsQuery,
) (*dto.GetAllPostsRes, error) {
	var cached *utils.PaginatedResponse[entity.PostBrief]

	queryKey, err := jsonUtils.MarshalToJsonString(query.MapToQuery())
	c.log.Warnf("getting cached all posts with query: %s", queryKey)
	cached, err = c.cacheRepo.GetPosts(ctx, queryKey)
	if err != nil {
		c.log.Warnf("failed to get all posts from cache: %v, fallback to db", err)
	}

	if cached == nil {
		posts, err := c.postRepo.GetPosts(ctx, query.MapToQuery())
		if err != nil {
			c.log.Errorf("failed to get all posts: %v", err)
			return nil, grpcerrors.ParseError(err)
		}
		total, err := c.postRepo.CountPosts(ctx, query.MapToQuery())
		if err != nil {
			c.log.Errorf("failed to count all posts: %v", err)
			return nil, grpcerrors.ParseError(err)
		}

		cached = &utils.PaginatedResponse[entity.PostBrief]{
			Data:       posts,
			Pagination: utils.NewPagination(query.Page, query.Limit, total),
		}

		if err := c.cacheRepo.PutPosts(ctx, queryKey, cached); err != nil {
			c.log.Warnf("failed to cache all posts: %v", err)
		}
	}

	postBriefs := make([]postDto.PostBriefDto, 0)
	for _, post := range cached.Data {
		postBrief := postDto.PostBriefDto{}
		postBrief.FromEntity(&post)
		postBriefs = append(postBriefs, postBrief)
	}

	return dto.NewGetAllPostsRes(
		postBriefs,
		*cached.Pagination,
	), nil
}
