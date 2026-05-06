package dto

import "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/entity"

type DeletePostDataDto struct {
	Id entity.PostId `json:"id"` // Primary key
}
