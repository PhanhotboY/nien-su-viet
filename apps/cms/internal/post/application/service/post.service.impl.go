package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/repository"
)

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{postRepo: repo}
}

func (bs *postService) FindPosts(ctx context.Context, query *dto.FindPostsQueryReq) ([]*entity.PostBrief, error) {
	return bs.postRepo.GetPosts(ctx, query.MapToQuery(), query.MapToPagination())
}

func (bs *postService) GetPublishedPosts(ctx context.Context, query *dto.GetPublicPostsQueryReq) ([]*entity.PostBrief, error) {
	return bs.postRepo.GetPosts(ctx, query.MapToQuery(true), query.MapToPagination())
}

func (bs *postService) FindPostByIdOrSlug(ctx context.Context, postId string) (*entity.Post, error) {
	if err := uuid.Validate(postId); err != nil {
		return bs.postRepo.GetPostBySlug(ctx, postId)
	}
	return bs.postRepo.GetPostByID(ctx, postId)
}

func (bs *postService) CreatePost(ctx context.Context, post *dto.PostCreateReq) (string, error) {
	postEntity := post.MapToEntity()
	return bs.postRepo.CreatePost(ctx, postEntity)
}

func (bs *postService) UpdatePost(ctx context.Context, postId string, post *dto.PostUpdateReq) (string, error) {
	postEntity := post.MapToEntity()
	return bs.postRepo.UpdatePost(ctx, postId, postEntity)
}

func (bs *postService) DeletePost(ctx context.Context, postId string) (string, error) {
	return bs.postRepo.DeletePost(ctx, postId)
}

func (bs *postService) PublishPost(ctx context.Context, postId string) (string, error) {
	post, err := bs.postRepo.GetPostByID(ctx, postId)
	if err != nil {
		return postId, err
	}

	post.Published = true
	return bs.postRepo.UpdatePost(ctx, postId, post)
}

func (bs *postService) UnpublishPost(ctx context.Context, postId string) (string, error) {
	post, err := bs.postRepo.GetPostByID(ctx, postId)
	if err != nil {
		return postId, err
	}

	post.Published = false
	return bs.postRepo.UpdatePost(ctx, postId, post)
}
