package entity

import (
	"database/sql/driver"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global/util"
)

// Media defines the media entity
type Media struct {
	Id           string     `gorm:"column:id;primaryKey" json:"id"`                      // Unique identifier
	Alt          *string    `gorm:"column:alt" json:"alt,omitempty"`                     // Alternative text for the media
	Caption      *string    `gorm:"column:caption" json:"caption,omitempty"`             // Caption or description
	URL          string     `gorm:"column:url;not null" json:"url"`                      // Primary URL of the media
	ThumbnailURL string     `gorm:"column:thumbnail_url" json:"thumbnail_url,omitempty"` // Thumbnail URL
	FileName     string     `gorm:"column:file_name;not null" json:"file_name"`          // Original file name
	FileSize     int64      `gorm:"column:file_size" json:"file_size"`                   // File size in bytes
	MimeType     string     `gorm:"column:mime_type;not null" json:"mime_type"`          // MIME type of the media
	Width        int        `gorm:"column:width" json:"width,omitempty"`                 // Media width in pixels
	Height       int        `gorm:"column:height" json:"height,omitempty"`               // Media height in pixels
	FolderId     *string    `gorm:"column:folder_id;index" json:"folder_id,omitempty"`   // Folder ID for organization
	FocalX       *float64   `gorm:"column:focal_x" json:"focal_x,omitempty"`             // Focal point X coordinate
	FocalY       *float64   `gorm:"column:focal_y" json:"focal_y,omitempty"`             // Focal point Y coordinate
	Sizes        MediaSizes `gorm:"column:sizes;type:jsonb" json:"sizes,omitempty"`      // Media sizes metadata
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // Creation timestamp
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // Last update timestamp
}

// MediaSizes represents a list of media size variants (stored as JSONB)
type MediaSizes []MediaSize

func (sizes MediaSizes) Value() (driver.Value, error) {
	if len(sizes) == 0 {
		return nil, nil
	}
	return util.JsonbValue(sizes)
}

func (sizes *MediaSizes) Scan(value interface{}) error {
	return util.JsonbScan(sizes, value)
}

// MediaSize defines metadata for a media size variant
type MediaSize struct {
	Variant  string `json:"variant"`   // Variant key (e.g., thumbnail, small, medium)
	Url      string `json:"url"`       // URL of the media variant
	Width    int    `json:"width"`     // Width of the variant in pixels
	Height   int    `json:"height"`    // Height of the variant in pixels
	MimeType string `json:"mime_type"` // MIME type of the variant
	FileSize int64  `json:"file_size"` // File size of the variant in bytes
}

// GORM override table name
func (Media) TableName() string {
	return "media"
}
