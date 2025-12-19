package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterMediaRoutes(api huma.API, handler *MediaHandler) {
	// GET /api/v1/media - Get media list
	huma.Register(api, huma.Operation{
		OperationID:   "get-media-list",
		Method:        http.MethodGet,
		Path:          "/api/v1/media",
		Summary:       "Get media list",
		Description:   "Retrieve all media items ordered by creation time (desc)",
		Tags:          []string{"media"},
		DefaultStatus: http.StatusOK,
	}, handler.GetMediaList)

	// POST /api/v1/media - Upload media
	// huma.Register(api, huma.Operation{
	// 	OperationID:   "upload-media",
	// 	Method:        http.MethodPost,
	// 	Path:          "/api/v1/media",
	// 	Summary:       "Upload media",
	// 	Description:   "Upload a new media file (max 10MB)",
	// 	Tags:          []string{"media"},
	// 	DefaultStatus: http.StatusOK,
	// 	MaxBodyBytes:  10 << 20, // 10MB limit
	// }, handler.UploadMedia)
}
