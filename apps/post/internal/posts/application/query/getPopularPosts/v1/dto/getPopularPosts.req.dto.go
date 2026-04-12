package dto

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
)

type GetPopularPostsQueryReq struct {
	dto.PostListQueryDto
}

func (g GetPopularPostsQueryReq) MapToQuery() repository.PostQuery {
	trueValue := true
	likesValue := "likes"
	descValue := "desc"
	g.Published = &trueValue
	g.SortBy = &likesValue
	g.SortOrder = &descValue
	return g.PostListQueryDto.MapToQuery()
}
