package dto

import (
	appDto "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
)

// GetAllPostsRequest is the request DTO for getting all posts
type GetAllPostsQueryReq struct {
	appDto.PostListQueryDto
}

func NewGetAllPostsQueryReqWithDefaultValue() GetAllPostsQueryReq {
	return GetAllPostsQueryReq{
		PostListQueryDto: appDto.NewPostListQueryDtoWithDefaultValue(),
	}
}

func (g GetAllPostsQueryReq) MapToQuery() repository.PostQuery {
	return g.PostListQueryDto.MapToQuery()
}
