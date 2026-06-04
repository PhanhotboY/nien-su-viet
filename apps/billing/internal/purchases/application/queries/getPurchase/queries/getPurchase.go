package aquery

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPurchaseQuery struct {
	*adto.GetPurchaseReqDto
}

func NewGetPurchaseQuery(req any) (*GetPurchaseQuery, error) {
	var typedReq = new(adto.GetPurchaseReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.NewValidationGrpcError(err.Error(), "NewGetPurchaseQuery")
	}

	return &GetPurchaseQuery{typedReq}, nil
}
