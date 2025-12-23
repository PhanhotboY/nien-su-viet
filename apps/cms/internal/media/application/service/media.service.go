package service

import (
	"context"
	"os"

	mediaDto "github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type MediaService interface {
	GetMediaList(ctx context.Context) ([]mediaDto.MediaRes, error)
	CreateFile(filename string) (*os.File, error)
	UpdateMedia(ctx context.Context) (*response.OperationResult, error)
}
