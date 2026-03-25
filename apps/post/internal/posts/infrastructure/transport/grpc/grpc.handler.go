package grpc

import (
	"context"

	"github.com/go-playground/validator"
	"go.opentelemetry.io/otel/attribute"
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
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/metrics"
	pb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type PostsGrpcServerHandler struct {
	postMetrics *metrics.PostsMetrics
	logger      logger.Logger
	validator   *validator.Validate

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

// var grpcMetricsAttr = api.WithAttributes(
// 	attribute.Key("MetricsType").String("Grpc"),
// )

func NewPostsGrpcServerHandler(
	logger logger.Logger,
	validator *validator.Validate,
	postMetrics *metrics.PostsMetrics,

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
		postMetrics: postMetrics,
		logger:      logger,
		validator:   validator,

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
	// p.postMetrics.CreatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := createPostCommand.NewCreatePostCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid create post command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.createPostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to create post: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.CreatePostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) UpdatePost(
	ctx context.Context,
	req *pb.UpdatePostRequest,
) (*pb.UpdatePostResponse, error) {
	p.logger.Info("[PostService] Handle update post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UpdatePost"))
	// p.postMetrics.UpdatePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := updatePostCommand.NewUpdatePostCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid update post command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.updatePostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to update post: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.UpdatePostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) PublishPost(
	ctx context.Context,
	req *pb.PublishPostRequest,
) (*pb.PublishPostResponse, error) {
	p.logger.Info("[PostService] Handle publish post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "PublishPost"))
	// p.postMetrics.PublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := publishPostCommand.NewPublishPostCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid publish post command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.publishPostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to publish post: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.PublishPostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) UnpublishPost(
	ctx context.Context,
	req *pb.UnpublishPostRequest,
) (*pb.UnpublishPostResponse, error) {
	p.logger.Info("[PostService] Handle unpublish post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "UnpublishPost"))
	// p.postMetrics.UnpublishPostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := unpublishPostCommand.NewUnpublishPostCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid unpublish post command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.unpublishPostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to unpublish post: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.UnpublishPostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) DeletePost(
	ctx context.Context,
	req *pb.DeletePostRequest,
) (*pb.DeletePostResponse, error) {
	p.logger.Info("[PostService] Handle delete post command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePost"))
	// p.postMetrics.DeletePostGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := deletePostCommand.NewDeletePostCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid delete post command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.deletePostHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to delete post: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.DeletePostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) DeletePosts(
	ctx context.Context,
	req *pb.DeletePostsRequest,
) (*pb.DeletePostsResponse, error) {
	p.logger.Info("[PostService] Handle delete posts command", "post_ids_count", len(req.GetPostIds()))
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "DeletePosts"))
	// p.postMetrics.DeletePostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := deletePostsCommand.NewDeletePostsCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid delete posts command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.deletePostsHandler.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to delete posts: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.DeletePostsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) IncrementPostViews(
	ctx context.Context,
	req *pb.IncrementPostViewsRequest,
) (*pb.IncrementPostViewsResponse, error) {
	p.logger.Info("[PostService] Handle increment post views command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostViews"))
	// p.postMetrics.IncrementPostViewsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := incrementPostViewsCommand.NewIncrementPostViewsCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid increment post views command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.incrementPostViews.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to increment post views: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.IncrementPostViewsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) IncrementPostLikes(
	ctx context.Context,
	req *pb.IncrementPostLikesRequest,
) (*pb.IncrementPostLikesResponse, error) {
	p.logger.Info("[PostService] Handle increment post likes command", "post_id", req.GetId())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "IncrementPostLikes"))
	// p.postMetrics.IncrementPostLikesGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	// Create command
	cmd, err := incrementPostLikesCommand.NewIncrementPostLikesCommandWithValidation(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid increment post likes command: %v", err)
		return nil, err
	}

	// Handle command
	res, err := p.incrementPostLikes.Handle(ctx, cmd)
	if err != nil {
		p.logger.Errorf("[PostService] Failed to increment post likes: %v", err)
		return nil, err
	}

	return dtoUtil.ValidateStruct(res, pb.IncrementPostLikesResponse{}, p.logger)
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

	query, err := getPostQuery.NewGetPostQueryWithValidation(req)
	if err != nil {
		p.logger.Error("[PostService] Invalid get post query", "error", err)
		return &pb.GetPostResponse{Data: nil}, nil
	}

	data, err := p.getPostHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Error("[PostService] Failed to handle get post query", "error", err)
		return &pb.GetPostResponse{Data: nil}, nil
	}

	return dtoUtil.ValidateStruct(data, pb.GetPostResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) GetPublishedPosts(
	ctx context.Context,
	req *pb.GetPublishedPostsRequest,
) (*pb.GetPublishedPostsResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPublishedPosts"))
	// p.postMetrics.GetPublishedPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	query, err := getPublishedPostsQuery.NewGetPublishedPostsQueryWithValidation(req)
	if err != nil {
		p.logger.Error("[PostService] Invalid get published posts query", "error", err)
		return &pb.GetPublishedPostsResponse{
			Data:       nil,
			Pagination: nil,
		}, nil
	}
	data, err := p.getPublishedPostsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Error("[PostService] Failed to handle get published posts query", "error", err)
		return &pb.GetPublishedPostsResponse{
			Data:       nil,
			Pagination: nil,
		}, nil
	}
	p.logger.Warnf("[PostService] Get published posts query result %+v", data)

	return dtoUtil.ValidateStruct(data, pb.GetPublishedPostsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) GetAllPosts(
	ctx context.Context,
	req *pb.GetAllPostsRequest,
) (*pb.GetAllPostsResponse, error) {
	p.logger.Info("[PostService] Handle get all posts query", "page", req.GetPage(), "limit", req.GetLimit())
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetAllPosts"))
	// p.postMetrics.ListPostsGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	query, err := getAllPostsQuery.NewGetAllPostsQueryWithValidation(req)
	if err != nil {
		p.logger.Error("[PostService] Invalid get all posts query", "error", err)
		return &pb.GetAllPostsResponse{
			Data:       nil,
			Pagination: nil,
		}, nil
	}

	data, err := p.getAllPostsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Error("[PostService] Failed to handle get all posts query", "error", err)
		return &pb.GetAllPostsResponse{
			Data:       nil,
			Pagination: nil,
		}, nil
	}

	return dtoUtil.ValidateStruct(data, pb.GetAllPostsResponse{}, p.logger)
}

func (p *PostsGrpcServerHandler) GetPostsByCategory(
	ctx context.Context,
	req *pb.GetPostsByCategoryRequest,
) (*pb.GetPostsByCategoryResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByCategory"))
	// p.postMetrics.GetPostsByCategoryGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

	return &pb.GetPostsByCategoryResponse{
		Data:       nil,
		Pagination: nil,
	}, nil
}

func (p *PostsGrpcServerHandler) GetPostsByAuthor(
	ctx context.Context,
	req *pb.GetPostsByAuthorRequest,
) (*pb.GetPostsByAuthorResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("rpc.method", "GetPostsByAuthor"))
	// p.postMetrics.GetPostsByAuthorGrpcRequests.Add(ctx, 1, grpcMetricsAttr)

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

	query, err := getPopularPostsQuery.NewGetPopularPostsQueryWithValidation(req)
	if err != nil {
		p.logger.Error("[PostService] Invalid get popular posts query", "error", err)
		return &pb.GetPopularPostsResponse{
			Data: nil,
		}, nil
	}

	data, err := p.getPopularPostsHandler.Handle(ctx, query)
	if err != nil {
		p.logger.Error("[PostService] Failed to handle get popular posts query", "error", err)
		return &pb.GetPopularPostsResponse{
			Data: nil,
		}, nil
	}

	return dtoUtil.ValidateStruct(data, pb.GetPopularPostsResponse{}, p.logger)
}
