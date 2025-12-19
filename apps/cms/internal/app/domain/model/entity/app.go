package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// App defines the app entity
type App struct {
	AppId       int64        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                // Primary key
	Title       string       `gorm:"column:title;not null" json:"title"`                          // Website title
	Description *string      `gorm:"column:description" json:"description,omitempty"`             // Website description
	Logo        *string      `gorm:"column:logo" json:"logo,omitempty"`                           // URL of the website logo
	Social      *SocialLinks `gorm:"column:social;type:jsonb" json:"social,omitempty"`            // Social media links
	TaxCode     *string      `gorm:"column:tax_code" json:"tax_code,omitempty"`                   // Tax code of the website owner
	Address     *Address     `gorm:"column:address;type:jsonb" json:"address,omitempty"`          // Physical address of the website owner
	Msisdn      *string      `gorm:"column:msisdn" json:"msisdn,omitempty"`                       // Contact phone number
	Email       *string      `gorm:"column:email" json:"email,omitempty"`                         // Contact email address
	Map         *string      `gorm:"column:map" json:"map,omitempty"`                             // Map location URL
	HeadScripts *string      `gorm:"column:head_scripts;type:text" json:"head_scripts,omitempty"` // Scripts to be included in the head section
	BodyScripts *string      `gorm:"column:body_scripts;type:text" json:"body_scripts,omitempty"` // Scripts to be included before the closing body tag
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime" json:"created_at"`          // Creation timestamp
	UpdatedAt   time.Time    `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`          // Last update timestamp
}

// jsonbValue is a generic helper for JSONB Value method
func jsonbValue(v interface{}) (driver.Value, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(v)
}

// jsonbScan is a generic helper for JSONB Scan method
func jsonbScan[T any](dest *T, value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, dest)
}

type SocialLinks struct {
	Facebook *string `json:"facebook,omitempty"` // Facebook page URL
	Youtube  *string `json:"youtube,omitempty"`  // YouTube channel URL
	Tiktok   *string `json:"tiktok,omitempty"`   // TikTok profile URL
	Zalo     *string `json:"zalo,omitempty"`     // Zalo profile URL
}

func (links *SocialLinks) Value() (driver.Value, error) {
	if links == nil {
		return nil, nil
	}
	return jsonbValue(*links)
}

func (links *SocialLinks) Scan(value interface{}) error {
	return jsonbScan(links, value)
}

type Address struct {
	Province *string `json:"province,omitempty"` // Province name
	District *string `json:"district,omitempty"` // District name
	Street   *string `json:"street,omitempty"`   // Street address
}

func (addr *Address) Value() (driver.Value, error) {
	if addr == nil {
		return nil, nil
	}
	return jsonbValue(*addr)
}

func (addr *Address) Scan(value interface{}) error {
	return jsonbScan(addr, value)
}

// GORM override table name
func (App) TableName() string {
	return "app_info"
}
