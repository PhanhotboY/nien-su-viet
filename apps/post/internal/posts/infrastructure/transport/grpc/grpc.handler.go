package grpc

import (
	"context"

	"github.com/go-playground/validator"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/metrics"
	pb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	commonpb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto/common"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type PostsGrpcServerHandler struct {
	postMetrics *metrics.PostsMetrics
	logger      logger.Logger
	validator   *validator.Validate
	pb.UnimplementedPostsServiceServer
}

// var grpcMetricsAttr = api.WithAttributes(
// 	attribute.Key("MetricsType").String("Grpc"),
// )

func NewPostsGrpcServerHandler(
	logger logger.Logger,
	validator *validator.Validate,
	postMetrics *metrics.PostsMetrics,
) *PostsGrpcServerHandler {
	return &PostsGrpcServerHandler{
		postMetrics: postMetrics,
		logger:      logger,
		validator:   validator,
	}
}

// ============================================================
// COMMAND HANDLERS
// ============================================================

func (p *PostsGrpcServerHandler) CreatePost(
	ctx context.Context,
	req *pb.CreatePostRequest,
) (*pb.CreatePostResponse, error) {
	p.logger.Info("[PostService] Handle create post command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "CreatePost"))
	// p.postMetrics.CreatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	// For now, return stub response
	return &pb.CreatePostResponse{
		Data: &commonpb.OperationMetadata{
			Id:      "",
			Success: true,
		},
	}, nil
}

func (p *PostsGrpcServerHandler) UpdatePost(
	ctx context.Context,
	req *pb.UpdatePostRequest,
) (*pb.UpdatePostResponse, error) {
	p.logger.Info("[PostService] Handle update post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UpdatePost"))
	// p.postMetrics.UpdatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.UpdatePostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostsGrpcServerHandler) PublishPost(
	ctx context.Context,
	req *pb.PublishPostRequest,
) (*pb.PublishPostResponse, error) {
	p.logger.Info("[PostService] Handle publish post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "PublishPost"))
	// p.postMetrics.PublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.PublishPostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostsGrpcServerHandler) UnpublishPost(
	ctx context.Context,
	req *pb.UnpublishPostRequest,
) (*pb.UnpublishPostResponse, error) {
	p.logger.Info("[PostService] Handle unpublish post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UnpublishPost"))
	// p.postMetrics.UnpublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.UnpublishPostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostsGrpcServerHandler) DeletePost(
	ctx context.Context,
	req *pb.DeletePostRequest,
) (*pb.DeletePostResponse, error) {
	p.logger.Info("[PostService] Handle delete post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePost"))
	// p.postMetrics.DeletePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.DeletePostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostsGrpcServerHandler) DeletePosts(
	ctx context.Context,
	req *pb.DeletePostsRequest,
) (*pb.DeletePostsResponse, error) {
	p.logger.Info("[PostService] Handle delete posts command", "post_ids_count", len(req.GetPostIds()))
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePosts"))
	// p.postMetrics.DeletePostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	resultID := ""
	if len(req.GetPostIds()) > 0 {
		resultID = req.GetPostIds()[0]
	}

	return &pb.DeletePostsResponse{
		Data: &commonpb.OperationMetadata{Id: resultID, Success: true},
	}, nil
}

func (p *PostsGrpcServerHandler) IncrementPostViews(
	ctx context.Context,
	req *pb.IncrementPostViewsRequest,
) (*pb.IncrementPostViewsResponse, error) {
	p.logger.Info("[PostService] Handle increment post views command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostViews"))
	// p.postMetrics.IncrementPostViewsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.IncrementPostViewsResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostsGrpcServerHandler) IncrementPostLikes(
	ctx context.Context,
	req *pb.IncrementPostLikesRequest,
) (*pb.IncrementPostLikesResponse, error) {
	p.logger.Info("[PostService] Handle increment post likes command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostLikes"))
	// p.postMetrics.IncrementPostLikesGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.IncrementPostLikesResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

// ============================================================
// QUERY HANDLERS
// ============================================================

func (p *PostsGrpcServerHandler) GetPost(
	ctx context.Context,
	req *pb.GetPostRequest,
) (*pb.GetPostResponse, error) {
	p.logger.Info("[PostService] Handle get post query", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPost"))
	// p.postMetrics.GetPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostResponse{
		Data: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPublishedPosts(
	ctx context.Context,
	req *pb.GetPublishedPostsRequest,
) (*pb.GetPublishedPostsResponse, error) {
	p.logger.Info("[PostService] Handle get published posts query")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPublishedPosts"))
	// p.postMetrics.GetPublishedPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPublishedPostsResponse{
		Data: make([]*pb.Post, 0),
	}, nil
}

func (p *PostsGrpcServerHandler) GetAllPosts(
	ctx context.Context,
	req *pb.GetAllPostsRequest,
) (*pb.GetAllPostsResponse, error) {
	p.logger.Info("[PostService] Handle get all posts query", "page", req.GetPage(), "limit", req.GetLimit())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetAllPosts"))
	// p.postMetrics.ListPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetAllPostsResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPostsByCategory(
	ctx context.Context,
	req *pb.GetPostsByCategoryRequest,
) (*pb.GetPostsByCategoryResponse, error) {
	p.logger.Info("[PostService] Handle get posts by category query", "category_id", req.GetCategoryId(), "page", req.GetPage(), "limit", req.GetLimit())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByCategory"))
	// p.postMetrics.GetPostsByCategoryGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostsByCategoryResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPostsByAuthor(
	ctx context.Context,
	req *pb.GetPostsByAuthorRequest,
) (*pb.GetPostsByAuthorResponse, error) {
	p.logger.Info("[PostService] Handle get posts by author query", "author_id", req.GetAuthorId(), "page", req.GetPage(), "limit", req.GetLimit())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByAuthor"))
	// p.postMetrics.GetPostsByAuthorGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostsByAuthorResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPopularPosts(
	ctx context.Context,
	req *pb.GetPopularPostsRequest,
) (*pb.GetPopularPostsResponse, error) {
	p.logger.Info("[PostService] Handle get popular posts query", "limit", req.GetLimit(), "days_ago", req.GetDaysAgo())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPopularPosts"))
	// p.postMetrics.GetPopularPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPopularPostsResponse{
		Data: nil,
	}, nil
}
