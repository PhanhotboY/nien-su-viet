package posts

import (
	"fmt"

	"github.com/phanhotboy/nien-su-viet/apps/post/config"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/configurations/posts/infrastructure"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/contracts"

	"go.opentelemetry.io/otel/metric"
	api "go.opentelemetry.io/otel/metric"
	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/

var PostsServiceModule = fx.Module(
	"postsModule",
	// Shared Modules
	config.Module,
	infrastructure.Module,

	// Features Modules
	posts.Module,

	// Other provides
	fx.Provide(configPostsMetrics),
)

// ref: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go

func configPostsMetrics(
	cfg *config.Config,
	meter metric.Meter,
) (*contracts.PostsMetrics, error) {
	if meter == nil {
		return nil, nil
	}

	appOptions := cfg.AppOptions
	successGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_success_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of success grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	errorGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_error_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of error grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	createPostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_create_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of create Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	updatePostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_update_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of update Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPostByIdGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_Post_by_id_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get Post by id grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPostsGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_Posts_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get Posts grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	searchPostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_search_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of search Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	deletePostRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_delete_Post_rabbitmq_messages_total", appOptions.ServiceName),
		api.WithDescription("The total number of delete Post rabbirmq messages"),
	)
	if err != nil {
		return nil, err
	}

	deletePostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_delete_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of delete Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	publishPostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_publish_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of publish Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	unpublishPostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_unpublish_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of unpublish Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPostGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_Post_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get Post grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPostBySlugGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_Post_by_slug_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get Post by slug grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	listPostsGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_list_Posts_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of list Posts grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPostsByCategoryGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_Posts_by_category_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get Posts by category grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPostsByAuthorGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_Posts_by_author_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get Posts by author grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	getPopularPostsGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_get_popular_Posts_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of get popular Posts grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	incrementPostViewsGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_increment_Post_views_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of increment Post views grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	incrementPostLikesGrpcRequests, err := meter.Float64Counter(
		fmt.Sprintf("%s_increment_Post_likes_grpc_requests_total", appOptions.ServiceName),
		api.WithDescription("The total number of increment Post likes grpc requests"),
	)
	if err != nil {
		return nil, err
	}

	createPostRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_create_Post_rabbitmq_messages_total", appOptions.ServiceName),
		api.WithDescription("The total number of create Post rabbirmq messages"),
	)
	if err != nil {
		return nil, err
	}

	updatePostRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_update_Post_rabbitmq_messages_total", appOptions.ServiceName),
		api.WithDescription("The total number of update Post rabbirmq messages"),
	)
	if err != nil {
		return nil, err
	}

	publishPostRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_publish_Post_rabbitmq_messages_total", appOptions.ServiceName),
		api.WithDescription("The total number of publish Post rabbitmq messages"),
	)
	if err != nil {
		return nil, err
	}

	unpublishPostRabbitMQMessages, err := meter.Float64Counter(
		fmt.Sprintf("%s_unpublish_Post_rabbitmq_messages_total", appOptions.ServiceName),
		api.WithDescription("The total number of unpublish Post rabbitmq messages"),
	)
	if err != nil {
		return nil, err
	}

	return &contracts.PostsMetrics{
		SuccessGrpcRequests:            successGrpcRequests,
		ErrorGrpcRequests:              errorGrpcRequests,
		CreatePostGrpcRequests:         createPostGrpcRequests,
		UpdatePostGrpcRequests:         updatePostGrpcRequests,
		DeletePostGrpcRequests:         deletePostGrpcRequests,
		PublishPostGrpcRequests:        publishPostGrpcRequests,
		UnpublishPostGrpcRequests:      unpublishPostGrpcRequests,
		GetPostByIdGrpcRequests:        getPostByIdGrpcRequests,
		GetPostsGrpcRequests:           getPostsGrpcRequests,
		GetPostGrpcRequests:            getPostGrpcRequests,
		GetPostBySlugGrpcRequests:      getPostBySlugGrpcRequests,
		ListPostsGrpcRequests:          listPostsGrpcRequests,
		GetPostsByCategoryGrpcRequests: getPostsByCategoryGrpcRequests,
		GetPostsByAuthorGrpcRequests:   getPostsByAuthorGrpcRequests,
		GetPopularPostsGrpcRequests:    getPopularPostsGrpcRequests,
		IncrementPostViewsGrpcRequests: incrementPostViewsGrpcRequests,
		IncrementPostLikesGrpcRequests: incrementPostLikesGrpcRequests,
		SearchPostsGrpcRequests:        searchPostGrpcRequests,
		DeletePostRmqMessages:          deletePostRabbitMQMessages,
		CreatePostRmqMessages:          createPostRabbitMQMessages,
		UpdatePostRmqMessages:          updatePostRabbitMQMessages,
		PublishPostRmqMessages:         publishPostRabbitMQMessages,
		UnpublishPostRmqMessages:       unpublishPostRabbitMQMessages,
	}, nil
}
