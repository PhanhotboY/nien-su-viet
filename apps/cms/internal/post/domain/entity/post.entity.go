package entity

import (
	"time"

	"github.com/google/uuid"
)

type PostId = uuid.UUID

// App defines the app entity
type Post struct {
	Id         PostId     `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"` // Primary key
	Title      string     `gorm:"column:title;not null" json:"title"`                                 // Website title
	Slug       string     `gorm:"column:slug;not null;unique" json:"slug"`                            // URL slug
	CategoryId *uuid.UUID `gorm:"column:category_id;type:uuid;nullable" json:"category_id"`           // Foreign key to category
	Thumbnail  *string    `gorm:"column:thumbnail" json:"thumbnail,omitempty"`                        // Thumbnail URL
	AuthorId   string     `gorm:"column:author_id;not null" json:"author_id"`                         // Foreign key to author

	Content string  `gorm:"column:content;type:text;not null" json:"content"`  // Post content
	Summary *string `gorm:"column:summary;type:text" json:"summary,omitempty"` // Post summary

	Views     int  `gorm:"column:views;default:0" json:"views"`             // Number of views
	Likes     int  `gorm:"column:likes;default:0" json:"likes"`             // Number of likes
	Published bool `gorm:"column:published;default:false" json:"published"` // Publication status

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // Creation timestamp
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // Last update timestamp
}

// App defines the app entity
type PostBrief struct {
	Id         PostId     `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"` // Primary key
	Title      string     `gorm:"column:title;not null" json:"title"`                                 // Website title
	Slug       string     `gorm:"column:slug;not null;unique" json:"slug"`                            // URL slug
	CategoryId *uuid.UUID `gorm:"column:category_id;type:uuid;nullable" json:"category_id"`           // Foreign key to category
	Thumbnail  *string    `gorm:"column:thumbnail" json:"thumbnail,omitempty"`                        // Thumbnail URL
	AuthorId   string     `gorm:"column:author_id;not null" json:"author_id"`                         // Foreign key to author

	Summary *string `gorm:"column:summary;type:text" json:"summary,omitempty"` // Post summary

	Views     int  `gorm:"column:views;default:0" json:"views"`             // Number of views
	Likes     int  `gorm:"column:likes;default:0" json:"likes"`             // Number of likes
	Published bool `gorm:"column:published;default:false" json:"published"` // Publication status

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // Creation timestamp
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // Last update timestamp
}

// GORM override table name
func (Post) TableName() string {
	return "posts"
}
