package contracts

import (
	"go.opentelemetry.io/otel/metric"
)

type PostsMetrics struct {
	SuccessGrpcRequests metric.Float64Counter
	ErrorGrpcRequests   metric.Float64Counter

	// Query Metrics
	GetPostsGrpcRequests           metric.Float64Counter
	GetPostByIdGrpcRequests        metric.Float64Counter
	GetPostGrpcRequests            metric.Float64Counter
	GetPostBySlugGrpcRequests      metric.Float64Counter
	ListPostsGrpcRequests          metric.Float64Counter
	GetPostsByCategoryGrpcRequests metric.Float64Counter
	GetPostsByAuthorGrpcRequests   metric.Float64Counter
	GetPopularPostsGrpcRequests    metric.Float64Counter
	SearchPostsGrpcRequests        metric.Float64Counter

	// Command Metrics
	CreatePostGrpcRequests         metric.Float64Counter
	UpdatePostGrpcRequests         metric.Float64Counter
	DeletePostGrpcRequests         metric.Float64Counter
	PublishPostGrpcRequests        metric.Float64Counter
	UnpublishPostGrpcRequests      metric.Float64Counter
	IncrementPostViewsGrpcRequests metric.Float64Counter
	IncrementPostLikesGrpcRequests metric.Float64Counter

	// RMQ Metrics
	SuccessRmqMessages metric.Float64Counter
	ErrorRmqMessages   metric.Float64Counter

	CreatePostRmqMessages    metric.Float64Counter
	DeletePostRmqMessages    metric.Float64Counter
	UpdatePostRmqMessages    metric.Float64Counter
	PublishPostRmqMessages   metric.Float64Counter
	UnpublishPostRmqMessages metric.Float64Counter
}
