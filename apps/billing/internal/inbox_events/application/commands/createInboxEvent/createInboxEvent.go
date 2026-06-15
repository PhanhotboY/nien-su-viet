package acmd // application command

import (
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/createInboxEvent/dto"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type CreateInboxEventCmd struct {
	*adto.CreateInboxEventReqDto
}

func NewCreateInboxEventCmd(req any) (*CreateInboxEventCmd, error) {
	typedReq := new(adto.CreateInboxEventReqDto)
	if err := dtoUtil.ValidateStruct(req, typedReq); err != nil {
		return nil, grpcerrors.ParseError(err)
	}

	return &CreateInboxEventCmd{CreateInboxEventReqDto: typedReq}, nil
}
