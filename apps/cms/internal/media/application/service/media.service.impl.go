package service

import (
	"context"
	"os"
	"path/filepath"

	mediaDto "github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/domain/repository"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type mediaService struct {
	repo repository.MediaRepository
}

// GetMediaList implements MediaService.
func (m *mediaService) GetMediaList(ctx context.Context) ([]mediaDto.MediaRes, error) {
	mediaEntities, err := m.repo.GetMediaList(ctx)
	if err != nil {
		return nil, err
	}

	mediaList := make([]mediaDto.MediaRes, 0, len(mediaEntities))
	for _, mediaEntity := range mediaEntities {
		media := mediaDto.MediaRes{}
		media.FromEntity(mediaEntity)
		mediaList = append(mediaList, media)
	}
	return mediaList, nil
}

func (m *mediaService) CreateFile(filename string) (*os.File, error) {
	// Create an uploads directory if it doesn’t exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// Build the file path and create it
	dst, err := os.Create(filepath.Join("uploads", filename))

	if err != nil {
		return nil, err
	}

	return dst, nil
}

func (m *mediaService) UpdateMedia(ctx context.Context) (*response.OperationResult, error) {
	return &response.OperationResult{Success: true}, nil
}

func NewMediaService(repo repository.MediaRepository) MediaService {
	return &mediaService{
		repo: repo,
	}
}
