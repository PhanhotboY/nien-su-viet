package posts

import (
	"go.uber.org/fx"

	getAllPostsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getAllPosts/v1/queries"
	getPopularPostsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPopularPosts/v1/queries"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPost/v1/queries"
	getPublishedPostsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPublishedPosts/v1/queries"

	createPostCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/createPost/v1/commands"
	deletePostCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/deletePost/v1/commands"
	updatePostCommand "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/command/updatePost/v1/commands"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/infrastructure/cache"
	rmqConsumer "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/infrastructure/messaging/rmq/consumer"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/infrastructure/persistence"

	// rmqProvider	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/infrastructure/messaging/rmq"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/infrastructure/transport/grpc"
)

var Module = fx.Module(
	"postsModule",

	fx.Provide(
		// Outbound Infrastructure
		cache.NewPostCacheRepository,
		persistence.NewPostEsRepository,

		// Application Query
		getPublishedPostsQuery.NewGetPublishedPostsHandler,
		getPopularPostsQuery.NewGetPopularPostsHandler,
		getPostQuery.NewGetPostHandler,
		getAllPostsQuery.NewGetAllPostsHandler,

		// Application Command
		createPostCommand.NewCreatePostHandler,
		updatePostCommand.NewUpdatePostHandler,
		deletePostCommand.NewDeletePostHandler,
	),

	// Inbound Infrastructure
	grpc.Module,
	rmqConsumer.Module,
)
