package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterPostRoutes(api huma.API, handler *PostHandler) {
	app := huma.NewGroup(api, "/posts")
	// GET /api/v1/posts - Get published posts list
	huma.Register(app, huma.Operation{
		OperationID:   "get-published-posts",
		Method:        http.MethodGet,
		Path:          "",
		Summary:       "Get published posts",
		Description:   "Retrieve brief published post list",
		Tags:          []string{"posts"},
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Successful response with published posts list",
			},
		},
	}, response.Wrap(handler.GetPublicPosts))

	// GET /api/v1/posts/all - Get posts list
	huma.Register(app, huma.Operation{
		OperationID:   "get-all-posts",
		Method:        http.MethodGet,
		Path:          "/all",
		Summary:       "Get all posts",
		Description:   "Retrieve brief all post list",
		Tags:          []string{"posts"},
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Successful response with all posts list",
			},
		},
	}, response.Wrap(handler.FindAllPosts))

	// GET /api/v1/posts/{id} - Get post by ID or slug
	huma.Register(app, huma.Operation{
		OperationID:   "get-post-by-id-or-slug",
		Method:        http.MethodGet,
		Path:          "/{idOrSlug}",
		Summary:       "Get post by ID",
		Description:   "Retrieve detailed post information by ID or slug",
		Tags:          []string{"posts"},
		DefaultStatus: http.StatusOK,
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Successful response with post information",
			},
		},
	}, response.Wrap(handler.GetPostByIDOrSlug))

	// POST /api/v1/posts - Create post
	huma.Register(app, huma.Operation{
		OperationID:   "create-post",
		Method:        http.MethodPost,
		Path:          "",
		Summary:       "Create post",
		Description:   "Create a new post",
		Tags:          []string{"posts"},
		DefaultStatus: http.StatusCreated,
	}, response.Wrap(handler.CreatePost))

	// PUT /api/v1/posts/{id} - Update post
	huma.Register(app, huma.Operation{
		OperationID:   "update-post",
		Method:        http.MethodPut,
		Path:          "/{id}",
		Summary:       "Update post",
		Description:   "Update post information by ID",
		Tags:          []string{"posts"},
		DefaultStatus: http.StatusOK,
	}, response.Wrap(handler.UpdatePost))

	// DELETE /api/v1/posts/{id} - Delete post
	huma.Register(app, huma.Operation{
		OperationID:   "delete-post",
		Method:        http.MethodDelete,
		Path:          "/{id}",
		Summary:       "Delete post",
		Description:   "Delete post by ID",
		Tags:          []string{"posts"},
		DefaultStatus: http.StatusOK,
	}, response.Wrap(handler.DeletePost))
}
