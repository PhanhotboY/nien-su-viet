package grpc

import (
	"context"

	"github.com/go-playground/validator"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	getAllPostsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getAllPosts/v1/queries"
	getPopularPostsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPopularPosts/v1/queries"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPost/v1/queries"
	getPublishedPostsQuery "github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/application/query/getPublishedPosts/v1/queries"
	pb "github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/grpc/genproto"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type PostsGrpcServerHandler struct {
	logger    logger.Logger
	validator *validator.Validate

	getPublishedPostsHandler getPublishedPostsQuery.GetPublishedPostsHandler
	getAllPostsHandler       getAllPostsQuery.GetAllPostsHandler
	getPopularPostsHandler   getPopularPostsQuery.GetPopularPostsHandler
	getPostHandler           getPostQuery.GetPostHandler
}

func NewPostsGrpcServerHandler(
	logger logger.Logger,
	validator *validator.Validate,

	getPublishedPostsHandler getPublishedPostsQuery.GetPublishedPostsHandler,
	getAllPostsHandler getAllPostsQuery.GetAllPostsHandler,
	getPopularPostsHandler getPopularPostsQuery.GetPopularPostsHandler,
	getPostHandler getPostQuery.GetPostHandler,
) *PostsGrpcServerHandler {
	return &PostsGrpcServerHandler{
		logger:    logger,
		validator: validator,

		getPublishedPostsHandler: getPublishedPostsHandler,
		getAllPostsHandler:       getAllPostsHandler,
		getPopularPostsHandler:   getPopularPostsHandler,
		getPostHandler:           getPostHandler,
	}
}

// ============================================================
// QUERY HANDLERS
// ============================================================

func (p *PostsGrpcServerHandler) GetPost(
	ctx context.Context,
	req *pb.GetPostRequest,
) (*pb.GetPostResponse, error) {
	p.logger.Infof("[PostService] Handle get post query: %+v", req)
	// span := trace.SpanFromContext(ctx)
	// span.SetAttributes(attribute.String("rpc.method", "GetPost"))

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
