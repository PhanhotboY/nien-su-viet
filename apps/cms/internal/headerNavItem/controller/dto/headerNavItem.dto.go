package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/headerNavItem/domain/model/entity"
)

// HeaderNavItemData represents a footer/header navigation item in API responses
type HeaderNavItemData struct {
	Id         string   `json:"id" example:"home"`                      // Unique identifier (slug or UUID)
	Order      int      `json:"order" example:"1"`                      // Display order (ascending)
	LinkNewTab *bool    `json:"link_new_tab,omitempty" example:"false"` // Whether to open link in new tab
	LinkUrl    string   `json:"link_url" example:"/"`                   // URL of the link (internal path or external URL)
	LinkLabel  string   `json:"link_label" example:"Home"`              // Display label for the link
	LinkType   LinkType `json:"link_type" example:"internal"`           // Type of link (internal or external)
}

func (dto *HeaderNavItemData) FromEntity(entity *entity.HeaderNavItem) {
	if dto == nil || entity == nil {
		return
	}
	dto.Id = entity.Id
	dto.Order = entity.Order
	dto.LinkNewTab = entity.LinkNewTab
	dto.LinkUrl = entity.LinkUrl
	dto.LinkLabel = entity.LinkLabel
	dto.LinkType = LinkType(entity.LinkType)
}

// HeaderNavItemResponse is the standard response wrapper for header nav item data
type HeaderNavItemResponse struct {
	Code    int               `json:"code" example:"200" doc:"HTTP status code"`
	Message string            `json:"message" example:"success" doc:"Response message"`
	Data    HeaderNavItemData `json:"data" doc:"Header navigation item information"`
}

// HeaderNavItemListResponse is the standard response wrapper for header nav item list
type HeaderNavItemListResponse struct {
	Code    int                 `json:"code" example:"200" doc:"HTTP status code"`
	Message string              `json:"message" example:"success" doc:"Response message"`
	Data    []HeaderNavItemData `json:"data" doc:"List of header navigation items"`
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

// HeaderNavItemCreateReq defines validation rules for creating a navigation item
type HeaderNavItemCreateReq struct {
	Order      int      `json:"order" binding:"required,min=0,max=1000" example:"1"`                     // Display order (0-1000)
	LinkNewTab *bool    `json:"link_new_tab,omitempty" binding:"omitempty" example:"false"`              // Whether to open link in new tab
	LinkUrl    string   `json:"link_url" binding:"required,min=1,max=2048" example:"/"`                  // Link URL or path
	LinkLabel  string   `json:"link_label" binding:"required,min=1,max=255" example:"Home"`              // Display label
	LinkType   LinkType `json:"link_type" binding:"required,oneof=internal external" example:"internal"` // Link type
}

// LinkType defines the type of link
type LinkType string

const (
	LinkTypeInternal LinkType = "internal" // Internal link
	LinkTypeExternal LinkType = "external" // External link
)

func (dto *HeaderNavItemCreateReq) MapToEntity() *entity.HeaderNavItem {
	if dto == nil {
		return nil
	}
	return &entity.HeaderNavItem{
		Order:      dto.Order,
		LinkNewTab: dto.LinkNewTab,
		LinkUrl:    dto.LinkUrl,
		LinkLabel:  dto.LinkLabel,
		LinkType:   entity.LinkType(dto.LinkType),
	}
}

type HeaderNavItemUpdateReq struct {
	Order      int      `json:"order" binding:"required,min=0,max=1000" example:"1"`                     // Display order
	LinkNewTab *bool    `json:"link_new_tab,omitempty" binding:"omitempty" example:"false"`              // Whether to open link in new tab
	LinkUrl    string   `json:"link_url" binding:"required,min=1,max=2048" example:"/"`                  // URL of the link
	LinkLabel  string   `json:"link_label" binding:"required,min=1,max=255" example:"Home"`              // Display label for the link
	LinkType   LinkType `json:"link_type" binding:"required,oneof=internal external" example:"internal"` // Type of link (internal or external)
}

func (dto *HeaderNavItemUpdateReq) MapToEntity() *entity.HeaderNavItem {
	if dto == nil {
		return nil
	}
	return &entity.HeaderNavItem{
		Order:      dto.Order,
		LinkNewTab: dto.LinkNewTab,
		LinkUrl:    dto.LinkUrl,
		LinkLabel:  dto.LinkLabel,
		LinkType:   entity.LinkType(dto.LinkType),
	}
}
