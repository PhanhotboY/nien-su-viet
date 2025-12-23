package http

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type MediaHandler struct {
	service service.MediaService
}

func NewMediaHandler(sv service.MediaService) *MediaHandler {
	return &MediaHandler{service: sv}
}

// UploadMediaInput represents the input for uploading media
type UploadMediaInput struct {
	File huma.FormFile `formData:"file" doc:"Image file to upload" required:"true" contentType:"application/octet-stream"`
}

// UploadMediaOutput represents the output for uploading media
type UploadMediaOutput struct {
	Body dto.OperationResponse
}

type GetMediaListInput struct{}

// GetMediaList retrieves all media items
func (h *MediaHandler) GetMediaList(ctx context.Context, input *GetMediaListInput) (*response.APIBodyResponse[[]dto.MediaRes], error) {
	mediaList, err := h.service.GetMediaList(ctx)
	if err != nil {
		return nil, err
	}

	return response.SuccessResponse(200, mediaList), nil
}

// UploadMedia uploads a new media file
func (h *MediaHandler) UploadMedia(ctx context.Context, input *UploadMediaInput) (*response.APIBodyResponse[response.OperationResult], error) {
	// Get the uploaded file from the multipart form
	file := input.File
	if file.File == nil {
		return nil, fmt.Errorf("no file uploaded")
	}

	defer file.Close()

	// Check file size (10MB limit)
	if file.Size > 10<<20 {
		return nil, fmt.Errorf("file size exceeds 10MB limit")
	}

	fmt.Printf("Uploaded File: %s\n", file.Filename)
	fmt.Printf("File Size: %d\n", file.Size)
	fmt.Printf("MIME Header: %v\n", file.ContentType)

	// Save the file
	dst, err := h.service.CreateFile(file.Filename)
	if err != nil {
		return nil, fmt.Errorf("error saving the file: %w", err)
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := dst.ReadFrom(file); err != nil {
		return nil, fmt.Errorf("error writing file: %w", err)
	}

	return response.OperationSuccessResponse(201), nil
}
