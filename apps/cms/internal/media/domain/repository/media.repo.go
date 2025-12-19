package repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/domain/model/entity"
)

type MediaRepository interface {
	GetMediaList(ctx context.Context) ([]entity.Media, error)
	FindMediaById(ctx context.Context, id string) (*entity.Media, error)
	CreateMedia(ctx context.Context, media *entity.Media) error
	UpdateMedia(ctx context.Context, id string, media *entity.Media) error
	DeleteMedia(ctx context.Context, id string) error
}
