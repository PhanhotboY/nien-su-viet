package contracts

import (
	"go.opentelemetry.io/otel/metric"
)

type PostsMetrics struct {
	SuccessGrpcRequests metric.Int64Counter
	FailedGrpcRequests  metric.Int64Counter

	CreatePostGrpcRequests         metric.Int64Counter
	UpdatePostGrpcRequests         metric.Int64Counter
	DeletePostGrpcRequests         metric.Int64Counter
	DeletePostsGrpcRequests        metric.Int64Counter
	IncrementPostViewsGrpcRequests metric.Int64Counter
	IncrementPostLikesGrpcRequests metric.Int64Counter
	PublishPostGrpcRequests        metric.Int64Counter
	UnpublishPostGrpcRequests      metric.Int64Counter

	GetPostGrpcRequests           metric.Int64Counter
	GetAllPostsGrpcRequests       metric.Int64Counter
	GetPublishedPostsGrpcRequests metric.Int64Counter
	GetPopularPostsGrpcRequests   metric.Int64Counter
}
