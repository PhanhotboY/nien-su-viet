package queries

import (
	"context"

	postDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPublishedPostsHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
	// dbRepository repository.PostRepository
	// cacheRepository repository.PostCacheRepository
}

type IGetPublishedPostsHandler interface {
	grpcTypes.GrpcHandler[*GetPublishedPostsQuery, *dto.GetPublishedPostsRes]
}

func NewGetPublishedPostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) GetPublishedPostsHandler {
	return GetPublishedPostsHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (c GetPublishedPostsHandler) Handle(
	ctx context.Context,
	query *GetPublishedPostsQuery,
) (*dto.GetPublishedPostsRes, error) {
	posts, err := c.postRepo.GetPosts(ctx, query.MapToQuery(), query.MapToPagination())
	if err != nil {
		c.log.Errorf("failed to get published posts: %v", err)
		return nil, err
	}
	c.log.Warnf("[PostService] Get published posts query result %+v", posts)

	postBriefs := make([]postDto.PostBriefDto, 0)
	for _, post := range posts {
		postBrief := postDto.PostBriefDto{}
		postBrief.FromEntity(post)
		postBriefs = append(postBriefs, postBrief)
	}

	return dto.NewGetPublishedPostsRes(
		postBriefs,
		*sharedDto.NewPaginationDto(query.Page, query.Limit, 10, 1),
	), nil
}
