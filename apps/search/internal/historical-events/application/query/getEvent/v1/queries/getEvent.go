package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getEvent/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetEventQuery struct {
	dto.GetEventQueryReq
}

func NewGetEventQuery(
	req any,
) (*GetEventQuery, error) {
	typedReq := new(dto.GetEventQueryReq)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetEventQuery")
	}

	return &GetEventQuery{
		GetEventQueryReq: *typedReq,
	}, nil
}
