package posts

import (
	"fmt"
	"strings"

	"go.opentelemetry.io/otel/metric"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
)

func configPostsMetrics(
	cfg settings.Config,
	meter metric.Meter,
) (*contracts.PostsMetrics, error) {
	if meter == nil {
		return nil, nil
	}

	serverOptions := cfg.Server

	metricsMap := map[string]metric.Int64Counter{
		"create_post_grpc_requests":          nil,
		"update_post_grpc_requests":          nil,
		"delete_post_grpc_requests":          nil,
		"delete_posts_grpc_requests":         nil,
		"increment_post_views_grpc_requests": nil,
		"increment_post_likes_grpc_requests": nil,
		"publish_post_grpc_requests":         nil,
		"unpublish_post_grpc_requests":       nil,

		"get_post_grpc_requests":            nil,
		"get_all_posts_grpc_requests":       nil,
		"get_published_posts_grpc_requests": nil,
		"get_popular_posts_grpc_requests":   nil,
	}
	for metricName := range metricsMap {
		counter, err := meter.Int64Counter(
			fmt.Sprintf("%s_%s_total", serverOptions.ServiceName, metricName),
			metric.WithDescription(fmt.Sprintf("The total number of %s", strings.ReplaceAll(metricName, "_", " "))),
		)
		if err != nil {
			return nil, err
		}
		metricsMap[metricName] = counter
	}

	return &contracts.PostsMetrics{
		CreatePostGrpcRequests:         metricsMap["create_post_grpc_requests"],
		UpdatePostGrpcRequests:         metricsMap["update_post_grpc_requests"],
		DeletePostGrpcRequests:         metricsMap["delete_post_grpc_requests"],
		DeletePostsGrpcRequests:        metricsMap["delete_posts_grpc_requests"],
		IncrementPostViewsGrpcRequests: metricsMap["increment_post_views_grpc_requests"],
		IncrementPostLikesGrpcRequests: metricsMap["increment_post_likes_grpc_requests"],
		PublishPostGrpcRequests:        metricsMap["publish_post_grpc_requests"],
		UnpublishPostGrpcRequests:      metricsMap["unpublish_post_grpc_requests"],

		GetPostGrpcRequests:           metricsMap["get_post_grpc_requests"],
		GetAllPostsGrpcRequests:       metricsMap["get_all_posts_grpc_requests"],
		GetPublishedPostsGrpcRequests: metricsMap["get_published_posts_grpc_requests"],
		GetPopularPostsGrpcRequests:   metricsMap["get_popular_posts_grpc_requests"],
	}, nil
}
