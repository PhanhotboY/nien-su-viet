package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/domain/repository"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db}
}

// UpdatePost implements repository.PostRepository.
func (br *postRepository) UpdatePost(ctx context.Context, postId string, post *entity.Post) (string, error) {
	result := br.db.WithContext(ctx).Model(&entity.Post{}).Where("id = ?", postId).Updates(post).Update("published", post.Published)
	if result.Error != nil {
		return postId, fmt.Errorf("failed to update post: %w", result.Error)
	}
	return postId, nil
}

func (br *postRepository) GetPosts(ctx context.Context, conds repository.PostQuery, pagination repository.PostPagination) ([]*entity.PostBrief, error) {
	var posts []*entity.PostBrief
	result := br.db.WithContext(ctx).Model(&entity.Post{}).Limit(pagination.Limit).Offset(pagination.Offset).Find(&posts, conds)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch post info: %w", result.Error)
	}
	return posts, nil
}

func (br *postRepository) CreatePost(ctx context.Context, post *entity.Post) (string, error) {
	if err := uuid.Validate(post.Id.String()); err != nil {
		post.Id = uuid.New()
	}
	result := br.db.WithContext(ctx).Model(&entity.Post{}).Create(post)
	if result.Error != nil {
		return post.Id.String(), fmt.Errorf("failed to create post: %w", result.Error)
	}
	return post.Id.String(), nil
}

func (br *postRepository) GetPostByID(ctx context.Context, postId string) (*entity.Post, error) {
	var post entity.Post
	result := br.db.WithContext(ctx).Model(&entity.Post{}).Where("id = ?", postId).First(&post)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch post by id: %w", result.Error)
	}
	return &post, nil
}

func (br *postRepository) GetPostBySlug(ctx context.Context, slug string) (*entity.Post, error) {
	var post entity.Post
	result := br.db.WithContext(ctx).Model(&entity.Post{}).Where("slug = ? AND published = ?", slug, true).First(&post)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch post by slug: %w", result.Error)
	}
	return &post, nil
}

func (br *postRepository) DeletePost(ctx context.Context, postId string) (string, error) {
	var post entity.Post
	result := br.db.WithContext(ctx).Model(&entity.Post{}).Where("id = ?", postId).Delete(&post)
	if result.Error != nil {
		return postId, fmt.Errorf("failed to fetch post by id: %w", result.Error)
	}
	return postId, nil
}
