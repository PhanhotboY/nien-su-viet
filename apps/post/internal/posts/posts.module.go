package posts

import (
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/messaging/rmq"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/metrics"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/transport/grpc"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"postsModule",

	// fx.Provide(metrics.ConfigPostsMetrics),
	fx.Supply(&metrics.PostsMetrics{}),

	grpc.Module,
	rmq.Module,
)
