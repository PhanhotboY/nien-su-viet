package queries

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type GetPostHandler struct {
	log      logger.Logger
	postRepo repository.PostRepository
}

type IGetPostHandler interface {
	grpcTypes.GrpcHandler[*GetPostQuery, *dto.GetPostRes]
}

func NewGetPostHandler(
	log logger.Logger,
	postRepo repository.PostRepository,
) GetPostHandler {
	return GetPostHandler{
		log:      log,
		postRepo: postRepo,
	}
}

func (c GetPostHandler) Handle(
	ctx context.Context,
	query *GetPostQuery,
) (*dto.GetPostRes, error) {
	c.log.Info("handling get post query")

	var post *entity.Post
	var err error

	if query.IsValidUUID() {
		c.log.Infof("fetching post by id: %s", query.IDOrSlug)
		post, err = c.postRepo.GetPostByID(ctx, query.IDOrSlug)
	} else {
		c.log.Infof("fetching post by slug: %s", query.IDOrSlug)
		post, err = c.postRepo.GetPostBySlug(ctx, query.IDOrSlug)
	}

	if err != nil {
		c.log.Errorf("failed to get post: %v", err)
		return nil, err
	}

	res := &dto.GetPostRes{}
	res.FromEntity(post)

	return res, nil
}
