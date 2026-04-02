package dto

import (
	appDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"
)

// GetAllPostsRequest is the request DTO for getting all posts
type GetAllPostsQueryReq struct {
	appDto.PostListQueryDto
}

func (g GetAllPostsQueryReq) MapToQuery() repository.PostQuery {
	return g.PostListQueryDto.MapToQuery()
}
