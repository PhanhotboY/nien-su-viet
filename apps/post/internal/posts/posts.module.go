package posts

import "go.uber.org/fx"

var Module = fx.Module(
	"postsModule",

	// fx.Provide(fx.Annotate(repositories.NewMongoPostReadRepository)),
)
