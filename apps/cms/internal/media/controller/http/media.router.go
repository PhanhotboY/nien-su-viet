package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

func RegisterMediaRoutes(api huma.API, handler *MediaHandler) {
	media := huma.NewGroup(api, "/media")
	// GET /api/v1/media - Get media list
	huma.Register(media, huma.Operation{
		OperationID:   "get-media-list",
		Method:        http.MethodGet,
		Path:          "",
		Summary:       "Get media list",
		Description:   "Retrieve all media items ordered by creation time (desc)",
		Tags:          []string{"media"},
		DefaultStatus: http.StatusOK,
	}, response.Wrap(handler.GetMediaList))

	// POST /api/v1/media - Upload media
	huma.Register(media, huma.Operation{
		OperationID:   "upload-media",
		Method:        http.MethodPost,
		Path:          "",
		Summary:       "Upload media",
		Description:   "Upload a new media file (max 10MB)",
		Tags:          []string{"media"},
		DefaultStatus: http.StatusCreated,
		MaxBodyBytes:  10 << 20, // 10MB limit
	}, response.Wrap(handler.UploadMedia))
}
