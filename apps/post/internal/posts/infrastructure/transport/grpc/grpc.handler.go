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

	pb "github.com/phanhotboy/nien-su-viet/apps/post/internal/shared/grpc/genproto"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type PostsGrpcServerHandler struct {
	logger    logger.Logger
	validator *validator.Validate

	createPostHandler  createPostCommand.CreatePostHandler
	deletePostHandler  deletePostCommand.DeletePostHandler
	deletePostsHandler deletePostsCommand.DeletePostsHandler
	updatePostHandler  updatePostCommand.UpdatePostHandler

	publishPostHandler   publishPostCommand.PublishPostHandler
	unpublishPostHandler unpublishPostCommand.UnpublishPostHandler
	incrementPostViews   incrementPostViewsCommand.IncrementPostViewsHandler
	incrementPostLikes   incrementPostLikesCommand.IncrementPostLikesHandler
}

func NewPostsGrpcServerHandler(
	logger logger.Logger,
	validator *validator.Validate,

	createPostHandler createPostCommand.CreatePostHandler,
	deletePostHandler deletePostCommand.DeletePostHandler,
	deletePostsHandler deletePostsCommand.DeletePostsHandler,
	updatePostHandler updatePostCommand.UpdatePostHandler,

	publishPostHandler publishPostCommand.PublishPostHandler,
	unpublishPostHandler unpublishPostCommand.UnpublishPostHandler,
	incrementPostViewsHandler incrementPostViewsCommand.IncrementPostViewsHandler,
	incrementPostLikesHandler incrementPostLikesCommand.IncrementPostLikesHandler,

) pb.PostsServiceServer {
	return &PostsGrpcServerHandler{
		logger:    logger,
		validator: validator,

		createPostHandler:  createPostHandler,
		deletePostHandler:  deletePostHandler,
		deletePostsHandler: deletePostsHandler,
		updatePostHandler:  updatePostHandler,

		publishPostHandler:   publishPostHandler,
		unpublishPostHandler: unpublishPostHandler,
		incrementPostViews:   incrementPostViewsHandler,
		incrementPostLikes:   incrementPostLikesHandler,
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

	// Create command
	cmd, err := createPostCommand.NewCreatePostCommand(req)
	if err != nil {
		p.logger.Errorf("[PostService] Invalid create post command: %s", err.Error())
		return nil, err
	}

	// Handle command
	p.logger.Debugf("[PostService] Created create post command: %+v", cmd)
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
