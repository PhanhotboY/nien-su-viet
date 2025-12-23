package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/footerNavItem/domain/model/entity"
)

type FooterNavItemData struct {
	Id         string   `json:"id" example:"d54934ae-0756-4da0-a8e2-e2f2dccb2904"`
	Order      int      `json:"order" example:"1"`
	LinkNewTab *bool    `json:"link_new_tab,omitempty" example:"false"`
	LinkUrl    string   `json:"link_url" example:"/contact"`
	LinkLabel  string   `json:"link_label" example:"Contact"`
	LinkType   LinkType `json:"link_type" example:"internal"`
}

func (dto *FooterNavItemData) FromEntity(e *entity.FooterNavItem) {
	if dto == nil || e == nil {
		return
	}
	dto.Id = e.Id
	dto.Order = e.Order
	dto.LinkNewTab = e.LinkNewTab
	dto.LinkUrl = e.LinkUrl
	dto.LinkLabel = e.LinkLabel
	dto.LinkType = LinkType(e.LinkType)
}

// FooterNavItemResponse is the standard response wrapper for footer nav item data
type FooterNavItemResponse struct {
	Code    int               `json:"code" example:"200" doc:"HTTP status code"`
	Message string            `json:"message" example:"success" doc:"Response message"`
	Data    FooterNavItemData `json:"data" doc:"Footer navigation item information"`
}

// FooterNavItemListResponse is the standard response wrapper for footer nav item list
type FooterNavItemListResponse struct {
	Code    int                 `json:"code" example:"200" doc:"HTTP status code"`
	Message string              `json:"message" example:"success" doc:"Response message"`
	Data    []FooterNavItemData `json:"data" doc:"List of footer navigation items"`
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

type FooterNavItemCreateReq struct {
	Order      int      `json:"order" binding:"required,min=0,max=1000" example:"1"`
	LinkNewTab *bool    `json:"link_new_tab,omitempty" binding:"omitempty" example:"false"`
	LinkUrl    string   `json:"link_url" binding:"required,min=1,max=2048" example:"/contact"`
	LinkLabel  string   `json:"link_label" binding:"required,min=1,max=255" example:"Contact"`
	LinkType   LinkType `json:"link_type" binding:"required,oneof=internal external" example:"internal"`
}

type LinkType string

const (
	LinkTypeInternal LinkType = "internal"
	LinkTypeExternal LinkType = "external"
)

func (dto *FooterNavItemCreateReq) MapToEntity() *entity.FooterNavItem {
	if dto == nil {
		return nil
	}
	return &entity.FooterNavItem{
		Order:      dto.Order,
		LinkNewTab: dto.LinkNewTab,
		LinkUrl:    dto.LinkUrl,
		LinkLabel:  dto.LinkLabel,
		LinkType:   entity.LinkType(dto.LinkType),
	}
}

type FooterNavItemUpdateReq struct {
	Order      int      `json:"order" binding:"required,min=0,max=1000" example:"1"`
	LinkNewTab *bool    `json:"link_new_tab,omitempty" binding:"omitempty" example:"false"`
	LinkUrl    string   `json:"link_url" binding:"required,min=1,max=2048" example:"/contact"`
	LinkLabel  string   `json:"link_label" binding:"required,min=1,max=255" example:"Contact"`
	LinkType   LinkType `json:"link_type" binding:"required,oneof=internal external" example:"internal"`
}

func (dto *FooterNavItemUpdateReq) MapToEntity() *entity.FooterNavItem {
	if dto == nil {
		return nil
	}
	return &entity.FooterNavItem{
		Order:      dto.Order,
		LinkNewTab: dto.LinkNewTab,
		LinkUrl:    dto.LinkUrl,
		LinkLabel:  dto.LinkLabel,
		LinkType:   entity.LinkType(dto.LinkType),
	}
}
