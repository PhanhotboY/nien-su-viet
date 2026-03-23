package dto

import "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/repository"

type GetPublicPostsQueryReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (g GetPublicPostsQueryReq) MapToQuery(published bool) repository.PostQuery {
	return repository.PostQuery{
		Published: published,
	}
}

func (g GetPublicPostsQueryReq) MapToPagination() repository.PostPagination {
	return repository.PostPagination{
		Limit:  g.Limit,
		Offset: g.Limit * (g.Page - 1),
	}
}

type FindPostsQueryReq struct {
	Page      int  `json:"page"`
	Limit     int  `json:"limit"`
	Published bool `json:"published" example:"true" doc:"Filter by publication status"`
}

func (f FindPostsQueryReq) MapToQuery() repository.PostQuery {
	return repository.PostQuery{
		Published: f.Published,
	}
}

func (f FindPostsQueryReq) MapToPagination() repository.PostPagination {
	return repository.PostPagination{
		Limit:  f.Limit,
		Offset: f.Limit * (f.Page - 1),
	}
}
