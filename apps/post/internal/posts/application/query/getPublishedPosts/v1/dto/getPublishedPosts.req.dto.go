package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
)

type GetPublicPostsQueryReq struct {
	dto.PostListQueryDto
}

func (g *GetPublicPostsQueryReq) MapToQuery() repository.PostQuery {
	trueValue := true
	g.Published = &trueValue
	return g.PostListQueryDto.MapToQuery()
}
