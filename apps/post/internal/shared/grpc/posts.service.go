package grpc

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	pb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	commonpb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto/common"
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

func (p *PostGrpcServiceServer) UpdatePost(
	ctx context.Context,
	req *pb.UpdatePostRequest,
) (*pb.UpdatePostResponse, error) {
	p.logger.Info("[PostService] Handle update post command")
	return nil, fmt.Errorf("Test error")

	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UpdatePost"))
	// p.postMetrics.UpdatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.UpdatePostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostGrpcServiceServer) PublishPost(
	ctx context.Context,
	req *pb.PublishPostRequest,
) (*pb.PublishPostResponse, error) {
	p.logger.Info("[PostService] Handle publish post command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "PublishPost"))
	// p.postMetrics.PublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.PublishPostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostGrpcServiceServer) UnpublishPost(
	ctx context.Context,
	req *pb.UnpublishPostRequest,
) (*pb.UnpublishPostResponse, error) {
	p.logger.Info("[PostService] Handle unpublish post command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UnpublishPost"))
	// p.postMetrics.UnpublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.UnpublishPostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostGrpcServiceServer) DeletePost(
	ctx context.Context,
	req *pb.DeletePostRequest,
) (*pb.DeletePostResponse, error) {
	p.logger.Info("[PostService] Handle delete post command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePost"))
	// p.postMetrics.DeletePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.DeletePostResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostGrpcServiceServer) DeletePosts(
	ctx context.Context,
	req *pb.DeletePostsRequest,
) (*pb.DeletePostsResponse, error) {
	p.logger.Info("[PostService] Handle delete posts command")
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

func (p *PostGrpcServiceServer) IncrementPostViews(
	ctx context.Context,
	req *pb.IncrementPostViewsRequest,
) (*pb.IncrementPostViewsResponse, error) {
	p.logger.Info("[PostService] Handle increment post views command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostViews"))
	// p.postMetrics.IncrementPostViewsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send command to mediatr
	return &pb.IncrementPostViewsResponse{
		Data: &commonpb.OperationMetadata{Id: req.GetId(), Success: true},
	}, nil
}

func (p *PostGrpcServiceServer) IncrementPostLike(
	ctx context.Context,
	req *pb.IncrementPostLikesRequest,
) (*pb.IncrementPostLikesResponse, error) {
	p.logger.Info("[PostService] Handle increment post likes command")
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

func (p *PostGrpcServiceServer) GetPost(
	ctx context.Context,
	req *pb.GetPostRequest,
) (*pb.GetPostResponse, error) {
	p.logger.Info("[PostService] Handle get post command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPost"))
	// p.postMetrics.GetPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostResponse{
		Data: nil,
	}, nil
}

func (p *PostGrpcServiceServer) GetPublishedPosts(
	ctx context.Context,
	req *pb.GetPublishedPostsRequest,
) (*pb.GetPublishedPostsResponse, error) {
	p.logger.Info("[PostService] Handle get published posts command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPublishedPosts"))
	// p.postMetrics.GetPostBySlugGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPublishedPostsResponse{
		Data: []*pb.Post{
			&pb.Post{
				Id:          "post-1",
				Title:       "Sample Post",
				Slug:        "sample-post",
				Content:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				AuthorId:    "author-1",
				PublishedAt: nil,
				Views:       100,
				Likes:       10,
			},
		},
		Pagination: &commonpb.PaginationMetadata{
			Total:      1,
			Limit:      10,
			Page:       1,
			TotalPages: 1,
		},
	}, nil
}

func (p *PostGrpcServiceServer) GetAllPosts(
	ctx context.Context,
	req *pb.GetAllPostsRequest,
) (*pb.GetAllPostsResponse, error) {
	p.logger.Info("[PostService] Handle get all posts command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetAllPosts"))
	// p.postMetrics.ListPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetAllPostsResponse{
		Data: nil,
		Pagination: &commonpb.PaginationMetadata{
			Total:      0,
			Limit:      req.GetLimit(),
			Page:       1,
			TotalPages: 0,
		},
	}, nil
}

func (p *PostGrpcServiceServer) GetPostsByCategory(
	ctx context.Context,
	req *pb.GetPostsByCategoryRequest,
) (*pb.GetPostsByCategoryResponse, error) {
	p.logger.Info("[PostService] Handle get posts by category command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByCategory"))
	// p.postMetrics.GetPostsByCategoryGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostsByCategoryResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostGrpcServiceServer) GetPostsByAuthor(
	ctx context.Context,
	req *pb.GetPostsByAuthorRequest,
) (*pb.GetPostsByAuthorResponse, error) {
	p.logger.Info("[PostService] Handle get posts by author command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByAuthor"))
	// p.postMetrics.GetPostsByAuthorGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPostsByAuthorResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostGrpcServiceServer) GetPopularPosts(
	ctx context.Context,
	req *pb.GetPopularPostsRequest,
) (*pb.GetPopularPostsResponse, error) {
	p.logger.Info("[PostService] Handle get popular posts command")
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPopularPosts"))
	// p.postMetrics.GetPopularPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// TODO: Send query to mediatr
	return &pb.GetPopularPostsResponse{
		Data: nil,
	}, nil
}
