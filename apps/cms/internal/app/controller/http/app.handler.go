package http

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/application/service"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/app/controller/dto"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/request"
	"github.com/phanhotboy/nien-su-viet/apps/cms/pkg/response"
)

type AppHandler struct {
	service service.AppService
}

func NewAppHandler(sv service.AppService) *AppHandler {
	return &AppHandler{service: sv}
}

// GetAppInfo retrieves app information
func (h *AppHandler) GetAppInfo(ctx context.Context, input *struct{}) (*GetAppRes, error) {
	result, err := h.service.GetAppInfo(ctx)
	if err != nil {
		return nil, err
	}

	app := &dto.AppData{}
	app.FromEntity(result)

	return response.SuccessResponse(200, *app), nil
}

// UpdateAppInfo updates app information
func (h *AppHandler) UpdateAppInfo(ctx context.Context, input *UpdateAppBodyReq) (*UpdateAppRes, error) {
	err := h.service.UpdateAppInfo(ctx, &input.Body)
	if err != nil {
		return nil, err
	}

	return response.SuccessResponse(201, response.OperationResult{
		Success: true,
	}), nil
}

type GetAppQueryReq struct{}
type GetAppRes = response.APIBodyResponse[dto.AppData]

type UpdateAppBodyReq = request.APIBodyRequest[dto.AppUpdateReq]
type UpdateAppRes = response.APIBodyResponse[response.OperationResult]
