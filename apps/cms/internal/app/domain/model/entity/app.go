package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// App defines the app entity
type App struct {
	AppId       int64        `gorm:"column:app_id;primaryKey;autoIncrement"`
	Title       string       `gorm:"column:title;not null"`
	Description *string      `gorm:"column:description"`
	Logo        *string      `gorm:"column:logo"`
	Social      *SocialLinks `gorm:"column:social;type:jsonb"`
	TaxCode     *string      `gorm:"column:tax_code"`
	Address     *Address     `gorm:"column:address;type:jsonb"`
	Msisdn      *string      `gorm:"column:msisdn"`
	Email       *string      `gorm:"column:email"`
	Map         *string      `gorm:"column:map"`
	HeadScripts *string      `gorm:"column:head_scripts;type:text"`
	BodyScripts *string      `gorm:"column:body_scripts;type:text"`
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;autoUpdateTime"`
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
	Facebook *string `json:"facebook,omitempty"`
	Youtube  *string `json:"youtube,omitempty"`
	Tiktok   *string `json:"tiktok,omitempty"`
	Zalo     *string `json:"zalo,omitempty"`
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
	Province *string `json:"province,omitempty"`
	District *string `json:"district,omitempty"`
	Street   *string `json:"street,omitempty"`
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
