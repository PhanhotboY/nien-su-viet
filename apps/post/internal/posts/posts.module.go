package posts

import (
	"go.uber.org/fx"

	createPostCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	deletePostCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/commands"
	deletePostsCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePosts/v1/commands"
	incrementPostLikesCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/commands"
	incrementPostViewsCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/commands"
	publishPostCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/commands"
	unpublishPostCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/unpublishPost/v1/commands"
	updatePostCommand "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"

	getAllPostsQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getAllPosts/v1/queries"
	getPopularPostsQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPopularPosts/v1/queries"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"
	getPublishedPostsQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/queries"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/cache"
	rmqConsumer "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/messaging/rmq/consumer"

	// rmqProvider	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/messaging/rmq"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/persistence"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/transport/grpc"
)

var Module = fx.Module(
	"postsModule",

	// Provide models for DB migration
	fx.Provide(fx.Annotate(
		func() any {
			return &entity.Post{}
		},
		fx.ResultTags(`group:"db_models"`),
	)),

	fx.Provide(
		configPostsMetrics,

		// Outbound Infrastructure
		persistence.NewPostRepository,
		cache.NewPostCacheRepository,

		// Application Query
		getPublishedPostsQuery.NewGetPublishedPostsHandler,
		getPopularPostsQuery.NewGetPopularPostsHandler,
		getPostQuery.NewGetPostHandler,
		getAllPostsQuery.NewGetAllPostsHandler,
		// Application Command
		createPostCommand.NewCreatePostHandler,
		deletePostCommand.NewDeletePostHandler,
		deletePostsCommand.NewDeletePostsHandler,
		updatePostCommand.NewUpdatePostHandler,
		incrementPostLikesCommand.NewIncrementPostLikesHandler,
		incrementPostViewsCommand.NewIncrementPostViewsHandler,
		publishPostCommand.NewPublishPostHandler,
		unpublishPostCommand.NewUnpublishPostHandler,
	),

	// Inbound Infrastructure
	grpc.Module,
	rmqConsumer.Module,
)
