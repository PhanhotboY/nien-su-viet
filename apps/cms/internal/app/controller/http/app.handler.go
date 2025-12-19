package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type AppHandler struct {
	service service.AppService
}

func NewAppHandler(sv service.AppService) *AppHandler {
	return &AppHandler{service: sv}
}

// GetAppInfoInput represents the input for getting app info
type GetAppInfoInput struct{}

// UpdateAppInfoInput represents the input for updating app info
type UpdateAppInfoInput struct {
	Body dto.AppUpdateReq
}

// UpdateAppInfoOutput represents the output for updating app info
type UpdateAppInfoOutput struct {
	Body response.APIOperationResponse
}

// GetAppInfo retrieves app information
func (h *AppHandler) GetAppInfo(ctx context.Context, input *GetAppInfoInput) (*response.APIBodyResponse[dto.AppData], error) {
	result, err := h.service.GetAppInfo(ctx)
	if err != nil {
		return nil, err
	}

	app := &dto.AppData{}
	app.FromEntity(result)

	return response.SuccessResponse(200, *app), nil
}

// UpdateAppInfo updates app information
func (h *AppHandler) UpdateAppInfo(ctx context.Context, input *UpdateAppInfoInput) (*response.APIBodyResponse[response.OperationResult], error) {
	err := h.service.UpdateAppInfo(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.SuccessResponse(201, response.OperationResult{
		Success: true,
	}), nil
}
