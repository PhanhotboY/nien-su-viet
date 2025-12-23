package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// FooterNavItem defines the footer navigation item entity
type FooterNavItem struct {
	Id         string    `gorm:"column:id;primaryKey" json:"id"`                     // Unique identifier
	Order      int       `gorm:"column:order;not null" json:"order"`                 // Display order
	LinkNewTab *bool     `gorm:"column:link_new_tab" json:"link_new_tab,omitempty"`  // Whether to open link in new tab
	LinkUrl    string    `gorm:"column:link_url;not null" json:"link_url"`           // URL of the link
	LinkLabel  string    `gorm:"column:link_label;not null" json:"link_label"`       // Display label for the link
	LinkType   LinkType  `gorm:"column:link_type;not null" json:"link_type"`         // Type of link (internal or external)
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // Creation timestamp
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // Last update timestamp
}

// LinkType defines the type of link
type LinkType string

const (
	LinkTypeInternal LinkType = "internal" // Internal link
	LinkTypeExternal LinkType = "external" // External link
)

func (l *LinkType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*l = LinkType(string(v))
		return nil
	case string:
		*l = LinkType(v)
		return nil
	default:
		return fmt.Errorf("unsupported Scan type for LinkType: %T", value)
	}
}

func (l LinkType) Value() (driver.Value, error) {
	return string(l), nil
}

// GORM override table name
func (FooterNavItem) TableName() string {
	return "footer_nav_items"
}
