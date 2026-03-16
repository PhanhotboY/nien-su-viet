package app

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/configurations/posts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.uber.org/fx"
)

type PostsApplication struct {
	*posts.PostsServiceConfigurator
}

func NewPostsApplication(
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environment.Environment,
) *PostsApplication {
	app := fxapp.NewApplication(providers, decorates, options, logger, environment)
	return &PostsApplication{
		PostsServiceConfigurator: posts.NewPostsServiceConfigurator(app),
	}
}
