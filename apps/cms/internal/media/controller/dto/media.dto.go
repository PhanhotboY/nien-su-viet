package dto

import (
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/domain/model/entity"
)

// MediaData represents media metadata for requests/responses
type MediaCreateReq struct {
	Alt      *string `json:"alt,omitempty" binding:"omitempty,max=255" example:"An example image"`     // Alternative text for the media
	Caption  *string `json:"caption,omitempty" binding:"omitempty,max=500" example:"Photo caption"`    // Caption or description
	URL      string  `json:"url" binding:"required,url" example:"https://cdn.example.com/image.jpg"`   // Primary URL of the media
	FolderId *string `json:"folder_id,omitempty" binding:"omitempty,min=1,max=100" example:"folder_1"` // Folder ID for organization
}

// MediaSizes represents a list of media size variants (stored as JSONB)
type MediaSizes []MediaSize

// MediaSize defines metadata for a media size variant
type MediaSize struct {
	Variant  string `json:"variant" binding:"required,min=1,max=50" example:"thumbnail"`                  // Variant key (e.g., thumbnail, small, medium)
	Url      string `json:"url" binding:"required,url" example:"https://cdn.example.com/image-thumb.jpg"` // URL of the media variant
	Width    int    `json:"width" binding:"required,min=0" example:"320"`                                 // Width of the variant in pixels
	Height   int    `json:"height" binding:"required,min=0" example:"180"`                                // Height of the variant in pixels
	MimeType string `json:"mime_type" binding:"required,min=1,max=255" example:"image/jpeg"`              // MIME type of the variant
	FileSize int64  `json:"file_size" binding:"required,min=0" example:"10240"`                           // File size of the variant in bytes
}

func (dto *MediaCreateReq) MapToEntity() *entity.Media {
	if dto == nil {
		return nil
	}
	media := &entity.Media{
		Alt:      dto.Alt,
		Caption:  dto.Caption,
		URL:      dto.URL,
		FolderId: dto.FolderId,
	}

	return media
}

// MediaData represents media metadata for requests/responses
type MediaRes struct {
	Id           string     `json:"id" binding:"required,min=1,max=100" example:"media_123"`                                     // Unique identifier
	Alt          *string    `json:"alt,omitempty" binding:"omitempty,max=255" example:"An example image"`                        // Alternative text for the media
	Caption      *string    `json:"caption,omitempty" binding:"omitempty,max=500" example:"Photo caption"`                       // Caption or description
	URL          string     `json:"url" binding:"required,url" example:"https://cdn.example.com/image.jpg"`                      // Primary URL of the media
	ThumbnailURL string     `json:"thumbnail_url,omitempty" binding:"omitempty,url" example:"https://cdn.example.com/thumb.jpg"` // Thumbnail URL
	FileName     string     `json:"file_name" binding:"required,min=1,max=255" example:"image.jpg"`                              // Original file name
	FileSize     int64      `json:"file_size" binding:"required,min=0" example:"204800"`                                         // File size in bytes
	MimeType     string     `json:"mime_type" binding:"required,min=1,max=255" example:"image/jpeg"`                             // MIME type of the media
	Width        int        `json:"width,omitempty" binding:"omitempty,min=0" example:"1920"`                                    // Media width in pixels
	Height       int        `json:"height,omitempty" binding:"omitempty,min=0" example:"1080"`                                   // Media height in pixels
	FolderId     *string    `json:"folder_id,omitempty" binding:"omitempty,min=1,max=100" example:"folder_1"`                    // Folder ID for organization
	FocalX       *float64   `json:"focal_x,omitempty" binding:"omitempty" example:"0.5"`                                         // Focal point X coordinate
	FocalY       *float64   `json:"focal_y,omitempty" binding:"omitempty" example:"0.5"`                                         // Focal point Y coordinate
	Sizes        MediaSizes `json:"sizes,omitempty" binding:"omitempty,dive"`                                                    // Media sizes metadata
	CreatedAt    time.Time  `json:"created_at" example:"2024-01-01T12:00:00Z"`                                                   // Creation timestamp
	UpdatedAt    time.Time  `json:"updated_at" example:"2024-01-02T12:00:00Z"`                                                   // Last update timestamp
}

func (dto *MediaRes) FromEntity(entity entity.Media) {
	if dto == nil {
		return
	}
	dto.Id = entity.Id
	dto.Alt = entity.Alt
	dto.Caption = entity.Caption
	dto.URL = entity.URL
	dto.ThumbnailURL = entity.ThumbnailURL
	dto.FileName = entity.FileName
	dto.FileSize = entity.FileSize
	dto.MimeType = entity.MimeType
	dto.Width = entity.Width
	dto.Height = entity.Height
	dto.FolderId = entity.FolderId
	dto.FocalX = entity.FocalX
	dto.FocalY = entity.FocalY
	dto.CreatedAt = entity.CreatedAt
	dto.UpdatedAt = entity.UpdatedAt
	for _, size := range entity.Sizes {
		dto.Sizes = append(dto.Sizes, MediaSize{
			Variant:  size.Variant,
			Url:      size.Url,
			Width:    size.Width,
			Height:   size.Height,
			MimeType: size.MimeType,
			FileSize: size.FileSize,
		})
	}
}

// MediaResponse is the standard response wrapper for media data
type MediaResponse struct {
	Code    int      `json:"code" example:"200" doc:"HTTP status code"`
	Message string   `json:"message" example:"success" doc:"Response message"`
	Data    MediaRes `json:"data" doc:"Media information"`
}

// MediaListResponse is the standard response wrapper for media list
type MediaListResponse struct {
	Code    int        `json:"code" example:"200" doc:"HTTP status code"`
	Message string     `json:"message" example:"success" doc:"Response message"`
	Data    []MediaRes `json:"data" doc:"List of media items"`
}

// OperationResult represents the result of an operation
type OperationResult struct {
	Success bool `json:"success" example:"true" doc:"Indicates if the operation was successful"`
}

// OperationResponse is the standard response wrapper for operation results
type OperationResponse struct {
	Code    int             `json:"code" example:"200" doc:"HTTP status code"`
	Message string          `json:"message" example:"success" doc:"Response message"`
	Data    OperationResult `json:"data" doc:"Operation result"`
}
