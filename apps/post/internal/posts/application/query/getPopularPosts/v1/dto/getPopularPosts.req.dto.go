package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
)

type GetPopularPostsQueryReq struct {
	dto.PostListQueryDto
}

func (g GetPopularPostsQueryReq) MapToQuery() repository.PostQuery {
	g.Published = true
	g.SortBy = "likes"
	g.SortOrder = "desc"
	return g.PostListQueryDto.MapToQuery()
}

func (g GetPopularPostsQueryReq) MapToPagination() repository.PostPagination {
	return g.PostListQueryDto.MapToPagination()
}
