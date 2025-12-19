package http

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/application/service"
	mediaDto "github.com/phanhotboy/nien-su-viet/apps/cms/internal/media/controller/dto"
)

type MediaHandler struct {
	service service.MediaService
}

func NewMediaHandler(sv service.MediaService) *MediaHandler {
	return &MediaHandler{service: sv}
}

// GetMediaListInput represents the input for getting media list
type GetMediaListInput struct{}

// GetMediaListOutput represents the output for getting media list
type GetMediaListOutput struct {
	Body mediaDto.MediaListResponse
}

// UploadMediaInput represents the input for uploading media
type UploadMediaInput struct {
	File huma.FormFile `form:"file" doc:"Image file to upload"`
}

// UploadMediaOutput represents the output for uploading media
type UploadMediaOutput struct {
	Body mediaDto.OperationResponse
}

// GetMediaList retrieves all media items
func (h *MediaHandler) GetMediaList(ctx context.Context, input *GetMediaListInput) (*GetMediaListOutput, error) {
	mediaList, err := h.service.GetMediaList(ctx)
	if err != nil {
		return nil, err
	}

	return &GetMediaListOutput{
		Body: mediaDto.MediaListResponse{
			Code:    200,
			Message: "success",
			Data:    mediaList,
		},
	}, nil
}

// UploadMedia uploads a new media file
// func (h *MediaHandler) UploadMedia(ctx context.Context, input *UploadMediaInput) (*UploadMediaOutput, error) {
// 	// Get the uploaded file from the multipart form
// 	file := input.File
// 	if file.File == nil {
// 		return nil, fmt.Errorf("no file uploaded with field name 'media'")
// 	}

// 	defer file.Close()

// 	// Check file size (10MB limit)
// 	if file.Size > 10<<20 {
// 		return nil, fmt.Errorf("file size exceeds 10MB limit")
// 	}

// 	fmt.Printf("Uploaded File: %s\n", file.Filename)
// 	fmt.Printf("File Size: %d\n", file.Size)
// 	fmt.Printf("MIME Header: %v\n", file.ContentType)

// 	// Save the file
// 	dst, err := h.service.CreateFile(file.Filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("error saving the file: %w", err)
// 	}
// 	defer dst.Close()

// 	// Copy the uploaded file to the destination
// 	if _, err := dst.ReadFrom(file); err != nil {
// 		return nil, fmt.Errorf("error writing file: %w", err)
// 	}

// 	return &UploadMediaOutput{
// 		Body: mediaDto.OperationResponse{
// 			Code:    200,
// 			Message: "success",
// 			Data: mediaDto.OperationResult{
// 				Success: true,
// 			},
// 		},
// 	}, nil
// }
