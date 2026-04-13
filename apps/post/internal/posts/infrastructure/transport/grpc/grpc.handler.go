package grpc

import (
	"context"

	"github.com/go-playground/validator"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

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
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/contracts"
	pb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type PostsGrpcServerHandler struct {
	postsMetrics *contracts.PostsMetrics
	logger       logger.Logger
	validator    *validator.Validate

	createPostHandler  createPostCommand.CreatePostHandler
	deletePostHandler  deletePostCommand.DeletePostHandler
	deletePostsHandler deletePostsCommand.DeletePostsHandler
	updatePostHandler  updatePostCommand.UpdatePostHandler

	publishPostHandler   publishPostCommand.PublishPostHandler
	unpublishPostHandler unpublishPostCommand.UnpublishPostHandler
	incrementPostViews   incrementPostViewsCommand.IncrementPostViewsHandler
	incrementPostLikes   incrementPostLikesCommand.IncrementPostLikesHandler

	getPublishedPostsHandler getPublishedPostsQuery.GetPublishedPostsHandler
	getAllPostsHandler       getAllPostsQuery.GetAllPostsHandler
	getPopularPostsHandler   getPopularPostsQuery.GetPopularPostsHandler
	getPostHandler           getPostQuery.GetPostHandler

	pb.UnimplementedPostsServiceServer
}

var grpcMetricsAttr = metric.WithAttributes(
	attribute.Key("MetricsType").String("Grpc"),
)

func NewPostsGrpcServerHandler(
	logger logger.Logger,
	validator *validator.Validate,
	postsMetrics *contracts.PostsMetrics,

	createPostHandler createPostCommand.CreatePostHandler,
	deletePostHandler deletePostCommand.DeletePostHandler,
	deletePostsHandler deletePostsCommand.DeletePostsHandler,
	updatePostHandler updatePostCommand.UpdatePostHandler,

	publishPostHandler publishPostCommand.PublishPostHandler,
	unpublishPostHandler unpublishPostCommand.UnpublishPostHandler,
	incrementPostViewsHandler incrementPostViewsCommand.IncrementPostViewsHandler,
	incrementPostLikesHandler incrementPostLikesCommand.IncrementPostLikesHandler,

	getPublishedPostsHandler getPublishedPostsQuery.GetPublishedPostsHandler,
	getAllPostsHandler getAllPostsQuery.GetAllPostsHandler,
	getPopularPostsHandler getPopularPostsQuery.GetPopularPostsHandler,
	getPostHandler getPostQuery.GetPostHandler,
) *PostsGrpcServerHandler {
	return &PostsGrpcServerHandler{
		postsMetrics: postsMetrics,
		logger:       logger,
		validator:    validator,

		createPostHandler:  createPostHandler,
		deletePostHandler:  deletePostHandler,
		deletePostsHandler: deletePostsHandler,
		updatePostHandler:  updatePostHandler,

		publishPostHandler:   publishPostHandler,
		unpublishPostHandler: unpublishPostHandler,
		incrementPostViews:   incrementPostViewsHandler,
		incrementPostLikes:   incrementPostLikesHandler,

		getPublishedPostsHandler: getPublishedPostsHandler,
		getAllPostsHandler:       getAllPostsHandler,
		getPopularPostsHandler:   getPopularPostsHandler,
		getPostHandler:           getPostHandler,
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
	p.postsMetrics.CreatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := createPostCommand.NewCreatePostCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid create post command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.createPostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to create post: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.CreatePostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) UpdatePost(
	ctx context.Context,
	req *pb.UpdatePostRequest,
) (*pb.UpdatePostResponse, error) {
	p.logger.Infof("[PostService] Handle update post command for: %s", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UpdatePost"))
	p.postsMetrics.UpdatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := updatePostCommand.NewUpdatePostCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid update post command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.updatePostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to update post: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.UpdatePostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) PublishPost(
	ctx context.Context,
	req *pb.PublishPostRequest,
) (*pb.PublishPostResponse, error) {
	p.logger.Infof("[PostService] Handle publish post command for: %s", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "PublishPost"))
	p.postsMetrics.PublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := publishPostCommand.NewPublishPostCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid publish post command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.publishPostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to publish post: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.PublishPostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) UnpublishPost(
	ctx context.Context,
	req *pb.UnpublishPostRequest,
) (*pb.UnpublishPostResponse, error) {
	p.logger.Infof("[PostService] Handle unpublish post command for: %s", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UnpublishPost"))
	p.postsMetrics.UnpublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := unpublishPostCommand.NewUnpublishPostCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid unpublish post command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.unpublishPostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to unpublish post: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.UnpublishPostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) DeletePost(
	ctx context.Context,
	req *pb.DeletePostRequest,
) (*pb.DeletePostResponse, error) {
	p.logger.Infof("[PostService] Handle delete post command for: %s", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePost"))
	p.postsMetrics.DeletePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := deletePostCommand.NewDeletePostCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid delete post command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.deletePostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to delete post: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.DeletePostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) DeletePosts(
	ctx context.Context,
	req *pb.DeletePostsRequest,
) (*pb.DeletePostsResponse, error) {
	p.logger.Infof("[PostService] Handle delete posts command: %s", len(req.GetPostIds()))
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePosts"))
	p.postsMetrics.DeletePostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := deletePostsCommand.NewDeletePostsCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid delete posts command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.deletePostsHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to delete posts: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.DeletePostsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) IncrementPostViews(
	ctx context.Context,
	req *pb.IncrementPostViewsRequest,
) (*pb.IncrementPostViewsResponse, error) {
	p.logger.Info("[PostService] Handle increment post views command for: %s", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostViews"))
	p.postsMetrics.IncrementPostViewsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := incrementPostViewsCommand.NewIncrementPostViewsCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid increment post views command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.incrementPostViews.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to increment post views: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.IncrementPostViewsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) IncrementPostLikes(
	ctx context.Context,
	req *pb.IncrementPostLikesRequest,
) (*pb.IncrementPostLikesResponse, error) {
	p.logger.Infof("[PostService] Handle increment post likes command for: %s", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostLikes"))
	p.postsMetrics.IncrementPostLikesGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := incrementPostLikesCommand.NewIncrementPostLikesCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid increment post likes command: %s", err.Error())
		return nil, err
	}

	// Handle command
	res, err := p.incrementPostLikes.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to increment post likes: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(res, &pb.IncrementPostLikesResponse{}, p.logger)
}

// ============================================================
// QUERY HANDLERS
// ============================================================

func (p *PostsGrpcServerHandler) GetPost(
	ctx context.Context,
	req *pb.GetPostRequest,
) (*pb.GetPostResponse, error) {
	p.logger.Infof("[PostService] Handle get post query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPost"))
	p.postsMetrics.GetPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	query, err := getPostQuery.NewGetPostQuery(req)
	if err != nil {
		p.logger.Error("[PostService] Invalid get post query", "error", err)
		return nil, err
	}

	data, err := p.getPostHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to handle get post query: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(data, &pb.GetPostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) GetPublishedPosts(
	ctx context.Context,
	req *pb.GetPublishedPostsRequest,
) (*pb.GetPublishedPostsResponse, error) {
	p.logger.Infof("[PostService] Handle get published posts query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPublishedPosts"))
	p.postsMetrics.GetPublishedPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	query, err := getPublishedPostsQuery.NewGetPublishedPostsQuery(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid get published posts query: %s", err.Error())
		return nil, err
	}
	data, err := p.getPublishedPostsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to handle get published posts query: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(data, &pb.GetPublishedPostsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) GetAllPosts(
	ctx context.Context,
	req *pb.GetAllPostsRequest,
) (*pb.GetAllPostsResponse, error) {
	p.logger.Infof("[PostService] Handle get all posts query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetAllPosts"))
	p.postsMetrics.GetAllPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	query, err := getAllPostsQuery.NewGetAllPostsQuery(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid get all posts query: %s", err.Error())
		return nil, err
	}

	data, err := p.getAllPostsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to handle get all posts query: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(data, &pb.GetAllPostsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) GetPostsByCategory(
	ctx context.Context,
	req *pb.GetPostsByCategoryRequest,
) (*pb.GetPostsByCategoryResponse, error) {
	p.logger.Infof("[PostService] Handle get posts by category query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByCategory"))
	// p.postsMetrics.GetPostsByCategoryGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	return &pb.GetPostsByCategoryResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPostsByAuthor(
	ctx context.Context,
	req *pb.GetPostsByAuthorRequest,
) (*pb.GetPostsByAuthorResponse, error) {
	p.logger.Infof("[PostService] Handle get posts by author query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByAuthor"))
	// p.postsMetrics.GetPostsByAuthorGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	return &pb.GetPostsByAuthorResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPopularPosts(
	ctx context.Context,
	req *pb.GetPopularPostsRequest,
) (*pb.GetPopularPostsResponse, error) {
	p.logger.Infof("[PostService] Handle get popular posts query: %+v", req)
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPopularPosts"))
	p.postsMetrics.GetPopularPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	query, err := getPopularPostsQuery.NewGetPopularPostsQuery(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid get popular posts query: %s", err.Error())
		return nil, err
	}

	data, err := p.getPopularPostsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to handle get popular posts query: %s", err.Error())
		return nil, err
	}

	return grpcUtils.UnmarshalProtoMessage(data, &pb.GetPopularPostsResponse{}, p.logger)
}
