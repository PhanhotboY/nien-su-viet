package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
)

type GetPublicPostsQueryReq struct {
	dto.PostListQueryDto
}

func (g *GetPublicPostsQueryReq) MapToQuery() repository.PostQuery {
	g.Published = true
	return g.PostListQueryDto.MapToQuery()
}

func (g *GetPublicPostsQueryReq) MapToPagination() repository.PostPagination {
	return g.PostListQueryDto.MapToPagination()
}
