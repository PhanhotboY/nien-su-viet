package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
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

func (br *postRepository) GetPosts(ctx context.Context, query repository.PostQuery) ([]entity.PostBrief, error) {
	var posts []entity.PostBrief

	db := br.db.WithContext(ctx).Model(&entity.Post{})

	// Apply filters
	db = br.applyPostFilters(db, query)

	// Apply sorting
	db = br.applySorting(db, query)

	// Apply pagination
	db = db.Limit(int(query.Limit)).Offset(int(query.Offset))

	result := db.Find(&posts)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", result.Error)
	}
	return posts, nil
}

// applyPostFilters applies all query filters to the GORM query
func (br *postRepository) applyPostFilters(db *gorm.DB, query repository.PostQuery) *gorm.DB {
	// Filter by published status
	if query.Published != nil {
		db = db.Where("published = ?", query.Published)
	}

	// Filter by author ID
	if query.AuthorID != "" {
		db = db.Where("author_id = ?", query.AuthorID)
	}

	// Filter by search term (search in title, slug, and content)
	if query.Search != "" {
		searchTerm := "%" + query.Search + "%"
		db = db.Where("title ILIKE ? OR slug ILIKE ? OR content ILIKE ?", searchTerm, searchTerm, searchTerm)
	}

	// Filter by created_at range
	if query.CreatedAtFrom != nil {
		db = db.Where("created_at >= ?", query.CreatedAtFrom)
	}
	if query.CreatedAtTo != nil {
		db = db.Where("created_at <= ?", query.CreatedAtTo)
	}

	// Filter by updated_at range
	if query.UpdatedAtFrom != nil {
		db = db.Where("updated_at >= ?", query.UpdatedAtFrom)
	}
	if query.UpdatedAtTo != nil {
		db = db.Where("updated_at <= ?", query.UpdatedAtTo)
	}

	return db
}

// applySorting applies sorting to the GORM query
func (br *postRepository) applySorting(db *gorm.DB, query repository.PostQuery) *gorm.DB {
	// Default sort by created_at descending
	if query.SortBy == "" {
		query.SortBy = "created_at"
	}
	if query.SortOrder == "" {
		query.SortOrder = "DESC"
	}

	// Whitelist allowed sort fields to prevent SQL injection
	allowedFields := map[string]bool{
		"id":           true,
		"title":        true,
		"slug":         true,
		"views":        true,
		"likes":        true,
		"created_at":   true,
		"updated_at":   true,
		"published_at": true,
	}

	if !allowedFields[query.SortBy] {
		query.SortBy = "created_at"
	}

	// Validate sort order
	if query.SortOrder != "ASC" && query.SortOrder != "DESC" {
		query.SortOrder = "DESC"
	}

	return db.Order(query.SortBy + " " + query.SortOrder)
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

func (br *postRepository) CountPosts(ctx context.Context, query repository.PostQuery) (uint32, error) {
	var count int64

	db := br.db.WithContext(ctx).Model(&entity.Post{})

	// Apply filters
	db = br.applyPostFilters(db, query)

	result := db.Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to count posts: %w", result.Error)
	}
	return uint32(count), nil
}
