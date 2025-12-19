package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type AppEntity = entity.App // AppEntity is an alias for entity.App for documentation purposes
// API response wrapper for App information
type AppInfoResponse struct {
	response.APIResponse[AppEntity]
}

type AppUpdateDto struct {
	Title       string       `json:"title,omitempty" binding:"omitempty,min=1,max=255" example:"Nien Su Viet"`                            // Website title
	Description *string      `json:"description,omitempty" binding:"omitempty,max=1000" example:"Vietnam history timeline website"`       // Website description
	Logo        *string      `json:"logo,omitempty" binding:"omitempty,url" example:"https://example.com/logo.png"`                       // URL of the website logo
	Social      *SocialLinks `json:"social,omitempty" binding:"omitempty"`                                                                // Social media links
	TaxCode     *string      `json:"tax_code,omitempty" binding:"omitempty,max=50" example:"0123456789"`                                  // Tax code of the website owner
	Address     *Address     `json:"address,omitempty" binding:"omitempty"`                                                               // Physical address of the website owner
	Msisdn      *string      `json:"msisdn,omitempty" binding:"omitempty,e164" example:"+84987654321"`                                    // Contact phone number (E.164 format)
	Email       *string      `json:"email,omitempty" binding:"omitempty,email" example:"contact@example.com"`                             // Contact email address
	Map         *string      `json:"map,omitempty" binding:"omitempty,url" example:"https://maps.google.com/?q=location"`                 // Map location URL
	HeadScripts *string      `json:"head_scripts,omitempty" binding:"omitempty,max=5000" example:"<script>console.log('head');</script>"` // Scripts to be included in the head section
	BodyScripts *string      `json:"body_scripts,omitempty" binding:"omitempty,max=5000" example:"<script>console.log('body');</script>"` // Scripts to be included before the closing body tag
}

type SocialLinks struct {
	Facebook *string `json:"facebook,omitempty" binding:"omitempty,url" example:"https://facebook.com/example"` // Facebook page URL
	Youtube  *string `json:"youtube,omitempty" binding:"omitempty,url" example:"https://youtube.com/@example"`  // YouTube channel URL
	Tiktok   *string `json:"tiktok,omitempty" binding:"omitempty,url" example:"https://tiktok.com/@example"`    // TikTok profile URL
	Zalo     *string `json:"zalo,omitempty" binding:"omitempty,url" example:"https://zalo.me/example"`          // Zalo profile URL
}

type Address struct {
	Province *string `json:"province,omitempty" binding:"omitempty,max=100" example:"Ho Chi Minh City"`  // Province name
	District *string `json:"district,omitempty" binding:"omitempty,max=100" example:"District 1"`        // District name
	Street   *string `json:"street,omitempty" binding:"omitempty,max=255" example:"123 Nguyen Hue Blvd"` // Street address
}

func (dto *AppUpdateDto) MapToEntity() *entity.App {
	if dto == nil {
		return nil
	}
	return &entity.App{
		Title:       dto.Title,
		Description: dto.Description,
		Logo:        dto.Logo,
		Social: &entity.SocialLinks{
			Facebook: dto.Social.Facebook,
			Youtube:  dto.Social.Youtube,
			Tiktok:   dto.Social.Tiktok,
			Zalo:     dto.Social.Zalo,
		},
		TaxCode: dto.TaxCode,
		Address: &entity.Address{
			Province: dto.Address.Province,
			District: dto.Address.District,
			Street:   dto.Address.Street,
		},
		Msisdn:      dto.Msisdn,
		Email:       dto.Email,
		Map:         dto.Map,
		HeadScripts: dto.HeadScripts,
		BodyScripts: dto.BodyScripts,
	}
}
