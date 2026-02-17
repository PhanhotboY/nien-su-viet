package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/domain/model/entity"
)

type AppData struct {
	Title       string       `json:"title,omitempty" binding:"omitempty,min=1,max=255" example:"Nien Su Viet" doc:"Website title"`                                                                 // Website title
	Description *string      `json:"description,omitempty" binding:"omitempty,max=1000" example:"Vietnam history timeline website" doc:"Website description"`                                      // Website description
	Logo        *string      `json:"logo,omitempty" binding:"omitempty,url" example:"https://example.com/logo.png" doc:"URL of the website logo"`                                                  // URL of the website logo
	Social      *SocialLinks `json:"social,omitempty" binding:"omitempty" doc:"Social media links"`                                                                                                // Social media links
	TaxCode     *string      `json:"tax_code,omitempty" binding:"omitempty,max=50" example:"0123456789" doc:"Tax code of the website owner"`                                                       // Tax code of the website owner
	Address     *Address     `json:"address,omitempty" binding:"omitempty" doc:"Physical address of the website owner"`                                                                            // Physical address of the website owner
	Msisdn      *string      `json:"msisdn,omitempty" binding:"omitempty,e164" example:"+84987654321" doc:"Contact phone number (E.164 format)"`                                                   // Contact phone number (E.164 format)
	Email       *string      `json:"email,omitempty" binding:"omitempty,email" example:"contact@example.com" doc:"Contact email address"`                                                          // Contact email address
	Map         *string      `json:"map,omitempty" binding:"omitempty,url" example:"https://maps.google.com/?q=location" doc:"Map location URL"`                                                   // Map location URL
	HeadScripts *string      `json:"head_scripts,omitempty" binding:"omitempty,max=5000" example:"<script>console.log('head');</script>" doc:"Scripts to be included in the head section"`         // Scripts to be included in the head section
	BodyScripts *string      `json:"body_scripts,omitempty" binding:"omitempty,max=5000" example:"<script>console.log('body');</script>" doc:"Scripts to be included before the closing body tag"` // Scripts to be included before the closing body tag
}

func (dto *AppData) FromEntity(entity *entity.App) {
	if dto == nil || entity == nil {
		return
	}

	social := &SocialLinks{}
	if entity.Social != nil {
		social.Facebook = entity.Social.Facebook
		social.Youtube = entity.Social.Youtube
		social.Tiktok = entity.Social.Tiktok
		social.Zalo = entity.Social.Zalo
	}
	address := &Address{}
	if entity.Address != nil {
		address.Province = entity.Address.Province
		address.District = entity.Address.District
		address.Street = entity.Address.Street
	}

	dto.Title = entity.Title
	dto.Description = entity.Description
	dto.Logo = entity.Logo
	dto.Social = social
	dto.TaxCode = entity.TaxCode
	dto.Address = address
	dto.Msisdn = entity.Msisdn
	dto.Email = entity.Email
	dto.Map = entity.Map
	dto.HeadScripts = entity.HeadScripts
	dto.BodyScripts = entity.BodyScripts
}

type AppUpdateReq struct {
	Title       string       `json:"title,omitempty" binding:"omitempty,min=1,max=255" example:"Nien Su Viet" doc:"Website title"`                                                                 // Website title
	Description *string      `json:"description,omitempty" binding:"omitempty,max=1000" example:"Vietnam history timeline website" doc:"Website description"`                                      // Website description
	Logo        *string      `json:"logo,omitempty" binding:"omitempty,url" example:"https://example.com/logo.png" doc:"URL of the website logo"`                                                  // URL of the website logo
	Social      *SocialLinks `json:"social,omitempty" binding:"omitempty" doc:"Social media links"`                                                                                                // Social media links
	TaxCode     *string      `json:"tax_code,omitempty" binding:"omitempty,max=50" example:"0123456789" doc:"Tax code of the website owner"`                                                       // Tax code of the website owner
	Address     *Address     `json:"address,omitempty" binding:"omitempty" doc:"Physical address of the website owner"`                                                                            // Physical address of the website owner
	Msisdn      *string      `json:"msisdn,omitempty" binding:"omitempty,e164" example:"+84987654321" doc:"Contact phone number (E.164 format)"`                                                   // Contact phone number (E.164 format)
	Email       *string      `json:"email,omitempty" binding:"omitempty,email" example:"contact@example.com" doc:"Contact email address"`                                                          // Contact email address
	Map         *string      `json:"map,omitempty" binding:"omitempty,url" example:"https://maps.google.com/?q=location" doc:"Map location URL"`                                                   // Map location URL
	HeadScripts *string      `json:"head_scripts,omitempty" binding:"omitempty,max=5000" example:"<script>console.log('head');</script>" doc:"Scripts to be included in the head section"`         // Scripts to be included in the head section
	BodyScripts *string      `json:"body_scripts,omitempty" binding:"omitempty,max=5000" example:"<script>console.log('body');</script>" doc:"Scripts to be included before the closing body tag"` // Scripts to be included before the closing body tag
}

type SocialLinks struct {
	Facebook *string `json:"facebook,omitempty" binding:"omitempty,url" example:"https://facebook.com/example" doc:"Facebook page URL"`
	Youtube  *string `json:"youtube,omitempty" binding:"omitempty,url" example:"https://youtube.com/@example" doc:"YouTube channel URL"`
	Tiktok   *string `json:"tiktok,omitempty" binding:"omitempty,url" example:"https://tiktok.com/@example" doc:"TikTok profile URL"`
	Zalo     *string `json:"zalo,omitempty" binding:"omitempty,url" example:"https://zalo.me/example" doc:"Zalo profile URL"`
}

type Address struct {
	Province *string `json:"province,omitempty" binding:"omitempty,max=100" example:"Ho Chi Minh City" doc:"Province name"`
	District *string `json:"district,omitempty" binding:"omitempty,max=100" example:"District 1" doc:"District name"`
	Street   *string `json:"street,omitempty" binding:"omitempty,max=255" example:"123 Nguyen Hue Blvd" doc:"Street address"`
}

func (dto *AppUpdateReq) MapToEntity() *entity.App {
	if dto == nil {
		return nil
	}

	var socialLinks *entity.SocialLinks
	if dto.Social != nil {
		socialLinks = &entity.SocialLinks{
			Facebook: dto.Social.Facebook,
			Youtube:  dto.Social.Youtube,
			Tiktok:   dto.Social.Tiktok,
			Zalo:     dto.Social.Zalo,
		}
	}
	var address *entity.Address
	if dto.Address != nil {
		address = &entity.Address{
			Province: dto.Address.Province,
			District: dto.Address.District,
			Street:   dto.Address.Street,
		}
	}

	return &entity.App{
		Title:       dto.Title,
		Description: dto.Description,
		Logo:        dto.Logo,
		Social:      socialLinks,
		TaxCode:     dto.TaxCode,
		Address:     address,
		Msisdn:      dto.Msisdn,
		Email:       dto.Email,
		Map:         dto.Map,
		HeadScripts: dto.HeadScripts,
		BodyScripts: dto.BodyScripts,
	}
}
