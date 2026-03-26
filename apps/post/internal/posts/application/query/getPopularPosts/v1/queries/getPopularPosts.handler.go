package queries

import (
	"context"

	postDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPopularPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPopularPostsHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
	// dbRepository repository.PostRepository
	// cacheRepository repository.PostCacheRepository
}

type IGetPopularPostsHandler interface {
	grpcTypes.GrpcHandler[*GetPopularPostsQuery, *dto.GetPopularPostsRes]
}

func NewGetPopularPostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) GetPopularPostsHandler {
	return GetPopularPostsHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (c GetPopularPostsHandler) Handle(
	ctx context.Context,
	query *GetPopularPostsQuery,
) (*dto.GetPopularPostsRes, error) {
	posts, err := c.postRepo.GetPosts(ctx, query.MapToQuery(), query.MapToPagination())
	if err != nil {
		c.log.Errorf("failed to get popular posts: %v", err)
		return nil, err
	}
	total, err := c.postRepo.CountPosts(ctx, query.MapToQuery())
	if err != nil {
		c.log.Errorf("failed to count popular posts: %v", err)
		return nil, err
	}

	postBriefs := make([]postDto.PostBriefDto, 0)
	for _, post := range posts {
		postBrief := postDto.PostBriefDto{}
		postBrief.FromEntity(post)
		postBriefs = append(postBriefs, postBrief)
	}

	return dto.NewGetPopularPostsRes(postBriefs, *sharedDto.NewPaginationDto(query.Page, query.Limit, total)), nil
}
