package entity

import (
	"time"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type PostId = string

// Post defines the post entity
type Post struct {
	Id          PostId     `json:"id"`                     // Primary key
	Title       string     `json:"title"`                  // Post title
	Slug        string     `json:"slug"`                   // URL slug
	Content     string     `json:"content"`                // Post content
	Summary     *string    `json:"summary,omitempty"`      // Post summary
	Thumbnail   *string    `json:"thumbnail,omitempty"`    // Thumbnail URL
	AuthorId    string     `json:"author_id"`              // Foreign key to author
	CategoryId  *string    `json:"category_id"`            // Foreign key to category
	Views       int        `json:"views"`                  // Number of views
	Likes       int        `json:"likes"`                  // Number of likes
	Published   bool       `json:"published"`              // Publication status
	PublishedAt *time.Time `json:"published_at,omitempty"` // Publication timestamp
	CreatedAt   time.Time  `json:"created_at"`             // Creation timestamp
	UpdatedAt   time.Time  `json:"updated_at"`             // Last update timestamp
}

// PostBrief defines a brief version of post entity with limited fields
type PostBrief struct {
	Id          PostId     `json:"id"`                     // Primary key
	Title       string     `json:"title"`                  // Post title
	Slug        string     `json:"slug"`                   // URL slug
	Summary     *string    `json:"summary,omitempty"`      // Post summary
	Thumbnail   *string    `json:"thumbnail,omitempty"`    // Thumbnail URL
	AuthorId    string     `json:"author_id"`              // Foreign key to author
	CategoryId  *string    `json:"category_id"`            // Foreign key to category
	Views       int        `json:"views"`                  // Number of views
	Likes       int        `json:"likes"`                  // Number of likes
	Published   bool       `json:"published"`              // Publication status
	PublishedAt *time.Time `json:"published_at,omitempty"` // Publication timestamp
	CreatedAt   time.Time  `json:"created_at"`             // Creation timestamp
	UpdatedAt   time.Time  `json:"updated_at"`             // Last update timestamp
}

// IndexName returns the name of the Elasticsearch index for Post entity
func (Post) IndexName() string {
	return "posts"
}

func (Post) ToTypeMapping() map[string]types.Property {
	return map[string]types.Property{
		"id":           types.KeywordProperty{},
		"title":        types.TextProperty{},
		"slug":         types.KeywordProperty{},
		"summary":      types.TextProperty{},
		"thumbnail":    types.KeywordProperty{},
		"author_id":    types.KeywordProperty{},
		"category_id":  types.KeywordProperty{},
		"views":        types.IntegerNumberProperty{},
		"likes":        types.IntegerNumberProperty{},
		"published":    types.BooleanProperty{},
		"published_at": types.DateProperty{},
		"created_at":   types.DateProperty{},
		"updated_at":   types.DateProperty{},
	}
}

func (PostBrief) IndexName() string {
	return "posts"
}
