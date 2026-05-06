package queries

import (
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getAllEvents/v1/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetAllEventsQuery struct {
	dto.GetAllEventsQueryReq
}

func NewGetAllEventsQuery(
	req any,
) (*GetAllEventsQuery, error) {
	typedReq := dto.NewGetAllEventsQueryReqWithDefaultValue()
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetAllEventsQuery")
	}

	return &GetAllEventsQuery{
		GetAllEventsQueryReq: *typedReq,
	}, nil
}
