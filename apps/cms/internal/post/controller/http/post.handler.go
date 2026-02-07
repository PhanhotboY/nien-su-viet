package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/post/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/request"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(sv service.PostService) *PostHandler {
	return &PostHandler{service: sv}
}

// GetPostInfo retrieves post information
func (h *PostHandler) GetPublicPosts(ctx context.Context, input *dto.GetPublicPostsQueryReq) (*GetPostListRes, error) {
	result, err := h.service.GetPublishedPosts(ctx, input)
	if err != nil {
		return nil, err
	}

	posts := make([]dto.PostBrief, 0)
	for _, r := range result {
		post := &dto.PostBrief{}
		post.FromEntity(r)
		posts = append(posts, *post)
	}

	return response.PaginatedSuccessResponse(200, posts, &response.Pagination{
		Total:      1,
		Limit:      10,
		Page:       1,
		TotalPages: 1,
	}), nil
}

// GetPostInfo retrieves post information
func (h *PostHandler) FindAllPosts(ctx context.Context, input *dto.FindPostsQueryReq) (*GetPostListRes, error) {
	result, err := h.service.FindPosts(ctx, input)
	if err != nil {
		return nil, err
	}

	posts := make([]dto.PostBrief, 0)
	for _, r := range result {
		post := &dto.PostBrief{}
		post.FromEntity(r)
		posts = append(posts, *post)
	}

	return response.PaginatedSuccessResponse(200, posts, &response.Pagination{
		Total:      1,
		Limit:      10,
		Page:       1,
		TotalPages: 1,
	}), nil
}

func (h *PostHandler) CreatePost(ctx context.Context, input *request.APIBodyRequest[dto.PostCreateReq]) (
	*response.APIBodyResponse[response.OperationResult], error,
) {
	id, err := h.service.CreatePost(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(201, id), nil
}

// PostIDInput represents input with a path parameter ID
type PostIDInput struct {
	ID string `path:"id" maxLength:"100" example:"item_123" doc:"Post ID"`
}

func (h *PostHandler) GetPostByIDOrSlug(ctx context.Context, input *struct {
	IdOrSlug string `path:"idOrSlug" doc:"Post id or slug"`
}) (*GetPostRes, error) {
	result, err := h.service.FindPostByIdOrSlug(ctx, input.IdOrSlug)
	if err != nil {
		return nil, err
	}

	post := &dto.PostDetail{}
	post.FromEntity(result)
	return response.SuccessResponse(200, *post), nil
}

func (h *PostHandler) PublishPost(ctx context.Context, input *struct {
	PostIDInput
}) (*response.APIBodyResponse[response.OperationResult], error) {
	id, err := h.service.PublishPost(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return response.OperationSuccessResponse(200, id), nil
}

func (h *PostHandler) UnpublishPost(ctx context.Context, input *struct {
	PostIDInput
}) (*response.APIBodyResponse[response.OperationResult], error) {
	id, err := h.service.UnpublishPost(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(200, id), nil
}

// UpdatePostInfo updates post information
func (h *PostHandler) UpdatePost(ctx context.Context, input *struct {
	PostIDInput
	UpdatePostBodyReq
}) (*UpdatePostRes, error) {
	id, err := h.service.UpdatePost(ctx, input.ID, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(201, id), nil
}

func (h *PostHandler) DeletePost(ctx context.Context, input *struct {
	PostIDInput
}) (*response.APIBodyResponse[response.OperationResult], error) {
	id, err := h.service.DeletePost(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return response.OperationSuccessResponse(200, id), nil
}

type GetPostRes = response.APIBodyResponse[dto.PostDetail]
type GetPostListRes = response.APIBodyResponse[[]dto.PostBrief]

type UpdatePostBodyReq = request.APIBodyRequest[dto.PostUpdateReq]
type UpdatePostRes = response.APIBodyResponse[response.OperationResult]
