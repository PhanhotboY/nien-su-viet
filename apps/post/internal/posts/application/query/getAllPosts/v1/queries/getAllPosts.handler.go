package queries

import (
	"context"

	postDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getAllPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/dto"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetAllPostsHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IGetAllPostsHandler interface {
	grpcTypes.GrpcHandler[*GetAllPostsQuery, *dto.GetAllPostsRes]
}

func NewGetAllPostsHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) GetAllPostsHandler {
	return GetAllPostsHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (c GetAllPostsHandler) Handle(
	ctx context.Context,
	query *GetAllPostsQuery,
) (*dto.GetAllPostsRes, error) {
	c.log.Info("handling get all posts query")

	c.log.Infof("query data: %v", query.Limit)

	posts, err := c.postRepo.GetPosts(ctx, query.MapToQuery(), query.MapToPagination())
	if err != nil {
		c.log.Errorf("failed to get all posts: %v", err)
		return nil, err
	}

	postBriefs := make([]postDto.PostBriefDto, 0)
	for _, post := range posts {
		postBrief := postDto.PostBriefDto{}
		postBrief.FromEntity(post)
		postBriefs = append(postBriefs, postBrief)
	}

	return dto.NewGetAllPostsRes(
		postBriefs,
		*sharedDto.NewPaginationDto(query.Page, query.Limit, 1, 1),
	), nil
}
