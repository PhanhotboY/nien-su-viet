package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/domain/repository"
	"gorm.io/gorm"
)

type mediaRepository struct {
	db *gorm.DB
}

// CreateMedia implements repository.MediaRepository.
func (m *mediaRepository) CreateMedia(ctx context.Context, media *entity.Media) error {
	if media.Id == "" {
		media.Id = uuid.New().String()
	}
	return m.db.WithContext(ctx).Create(media).Error
}

// DeleteMedia implements repository.MediaRepository.
func (m *mediaRepository) DeleteMedia(ctx context.Context, id string) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Media{}).Error
}

// FindMediaById implements repository.MediaRepository.
func (m *mediaRepository) FindMediaById(ctx context.Context, id string) (*entity.Media, error) {
	var media entity.Media
	if err := m.db.WithContext(ctx).First(&media, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &media, nil
}

// GetMediaList implements repository.MediaRepository.
func (m *mediaRepository) GetMediaList(ctx context.Context) ([]entity.Media, error) {
	var media []entity.Media
	if err := m.db.WithContext(ctx).Order("created_at DESC").Find(&media).Error; err != nil {
		return nil, err
	}
	return media, nil
}

// UpdateMedia implements repository.MediaRepository.
func (m *mediaRepository) UpdateMedia(ctx context.Context, id string, media *entity.Media) error {
	return m.db.WithContext(ctx).Model(&entity.Media{}).Where("id = ?", id).Updates(media).Error
}

func NewMediaRepository(db *gorm.DB) repository.MediaRepository {
	return &mediaRepository{db}
}
