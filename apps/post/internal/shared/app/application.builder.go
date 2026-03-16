package app

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp/contracts"
)

type PostsApplicationBuilder struct {
	contracts.ApplicationBuilder
}

func NewPostsApplicationBuilder() *PostsApplicationBuilder {
	return &PostsApplicationBuilder{fxapp.NewApplicationBuilder()}
}

func (a *PostsApplicationBuilder) Build() *PostsApplication {
	return NewPostsApplication(
		a.GetProvides(),
		a.GetDecorates(),
		a.Options(),
		a.Logger(),
		a.Environment(),
	)
}
