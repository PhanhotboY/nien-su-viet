package grpc

import (
	"context"

	"github.com/go-playground/validator"
	pb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type PostGrpcServiceServer struct {
	// postMetrics *contracts.PostsMetrics
	logger    logger.Logger
	validator *validator.Validate
	pb.UnimplementedPostsServiceServer
}

var grpcMetricsAttr = api.WithAttributes(
	attribute.Key("MetricsType").String("Grpc"),
)

func NewPostGrpcServiceServer(
	logger logger.Logger,
	validator *validator.Validate,
	// postMetrics *contracts.PostsMetrics,
) *PostGrpcServiceServer {
	return &PostGrpcServiceServer{
		// postMetrics: postMetrics,
		logger:    logger,
		validator: validator,
	}
}

// ============================================================
// COMMAND HANDLERS
// ============================================================

func (p *PostGrpcServiceServer) CreatePost(
	ctx context.Context,
	req *pb.CreatePostCommand,
) (*pb.CreatePostCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "CreatePost"))
	// p.postMetrics.CreatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	// For now, return stub response
	return &pb.CreatePostCommandResponse{
		Id:      "",
		Success: true,
		Message: "Post created successfully",
	}, nil
}

func (p *PostGrpcServiceServer) UpdatePost(
	ctx context.Context,
	req *pb.UpdatePostCommand,
) (*pb.UpdatePostCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UpdatePost"))
	// p.postMetrics.UpdatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.UpdatePostCommandResponse{
		Success: true,
		Message: "Post updated successfully",
	}, nil
}

func (p *PostGrpcServiceServer) PublishPost(
	ctx context.Context,
	req *pb.PublishPostCommand,
) (*pb.PublishPostCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "PublishPost"))
	// p.postMetrics.PublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.PublishPostCommandResponse{
		Success: true,
		Message: "Post published successfully",
	}, nil
}

func (p *PostGrpcServiceServer) UnpublishPost(
	ctx context.Context,
	req *pb.UnpublishPostCommand,
) (*pb.UnpublishPostCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UnpublishPost"))
	// p.postMetrics.UnpublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.UnpublishPostCommandResponse{
		Success: true,
		Message: "Post unpublished successfully",
	}, nil
}

func (p *PostGrpcServiceServer) DeletePost(
	ctx context.Context,
	req *pb.DeletePostCommand,
) (*pb.DeletePostCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePost"))
	// p.postMetrics.DeletePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.DeletePostCommandResponse{
		Success: true,
		Message: "Post deleted successfully",
	}, nil
}

func (p *PostGrpcServiceServer) IncrementPostViews(
	ctx context.Context,
	req *pb.IncrementPostViewsCommand,
) (*pb.IncrementPostViewsCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostViews"))
	// p.postMetrics.IncrementPostViewsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.IncrementPostViewsCommandResponse{
		Views:   0,
		Success: true,
	}, nil
}

func (p *PostGrpcServiceServer) IncrementPostLikes(
	ctx context.Context,
	req *pb.IncrementPostLikesCommand,
) (*pb.IncrementPostLikesCommandResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostLikes"))
	// p.postMetrics.IncrementPostLikesGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.IncrementPostLikesCommandResponse{
		Likes:   0,
		Success: true,
	}, nil
}

// ============================================================
// QUERY HANDLERS
// ============================================================

func (p *PostGrpcServiceServer) GetPost(
	ctx context.Context,
	req *pb.GetPostQuery,
) (*pb.GetPostQueryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPost"))
	// p.postMetrics.GetPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostQueryResponse{
		Post:  nil,
		Found: false,
	}, nil
}

func (p *PostGrpcServiceServer) GetPostBySlug(
	ctx context.Context,
	req *pb.GetPostBySlugQuery,
) (*pb.GetPostBySlugQueryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostBySlug"))
	// p.postMetrics.GetPostBySlugGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostBySlugQueryResponse{
		Post:  nil,
		Found: false,
	}, nil
}

func (p *PostGrpcServiceServer) ListPosts(
	ctx context.Context,
	req *pb.ListPostsQuery,
) (*pb.ListPostsQueryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "ListPosts"))
	// p.postMetrics.ListPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.ListPostsQueryResponse{
		Posts:      nil,
		Pagination: nil,
	}, nil
}

func (p *PostGrpcServiceServer) GetPostsByCategory(
	ctx context.Context,
	req *pb.GetPostsByCategoryQuery,
) (*pb.GetPostsByCategoryQueryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByCategory"))
	// p.postMetrics.GetPostsByCategoryGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostsByCategoryQueryResponse{
		Posts:      nil,
		Pagination: nil,
	}, nil
}

func (p *PostGrpcServiceServer) GetPostsByAuthor(
	ctx context.Context,
	req *pb.GetPostsByAuthorQuery,
) (*pb.GetPostsByAuthorQueryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByAuthor"))
	// p.postMetrics.GetPostsByAuthorGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostsByAuthorQueryResponse{
		Posts:      nil,
		Pagination: nil,
	}, nil
}

func (p *PostGrpcServiceServer) GetPopularPosts(
	ctx context.Context,
	req *pb.GetPopularPostsQuery,
) (*pb.GetPopularPostsQueryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPopularPosts"))
	// p.postMetrics.GetPopularPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPopularPostsQueryResponse{
		Posts: nil,
	}, nil
}
