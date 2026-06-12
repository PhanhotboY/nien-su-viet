package webhook

import (
	"context"

	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/infrastructure/zalopay"
	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	googleGrpc "google.golang.org/grpc"

	createInboxEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands"
	createInboxEventDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/dto"
)

var Module = fx.Module(
	"zalopayWebhook",

	fx.Invoke(
		NewZalopayGrpcServiceServer,
	))

type zalopayGrpcServiceServer struct {
	logger        logger.Logger
	zalopayClient *zalopay.Client

	createInboxEvent createInboxEvent.CreateInboxEventCmdHandler
}

func NewZalopayGrpcServiceServer(
	zalopayGrpcServer grpcServer.GrpcServer,
	logger logger.Logger,
	zpClient *zalopay.Client,

	createInboxEvent createInboxEvent.CreateInboxEventCmdHandler,
) error {
	zalopayGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
		billing_service.RegisterZaloPayWebhookServiceServer(server,
			&zalopayGrpcServiceServer{
				logger:           logger,
				zalopayClient:    zpClient,
				createInboxEvent: createInboxEvent,
			},
		)
	})
	return nil
}

func (s *zalopayGrpcServiceServer) HandleCallback(
	ctx context.Context,
	req *billing_service.CallbackPayload,
) (*billing_service.CallbackResponse, error) {
	payload := zalopay.NewCallbackPayload(req)

	// validate callback payload
	if err := s.zalopayClient.VerifyCallback([]byte(payload.Data), payload.Mac); err != nil {
		s.logger.Error("invalid callback payload: ", err)
		return zalopay.NewCallbackResponse(-1, "MAC not equal").ToGrpcResponse(), nil
	}

	// TODO: insert inbox event
	createInboxEventCmd, err := createInboxEvent.NewCreateInboxEventCmd(createInboxEventDto.CreateInboxEventReqDto{
		EventType:       "zalopay_payment_success",
		Provider:        "zalopay",
		ExternalEventID: payload.GetData().AppTransID,
		Payload:         payload.Data,
		Signature:       payload.Mac,
	})
	if err != nil {
		s.logger.Error("failed to create createInboxEventCmd: ", err)
		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}
	if _, err := s.createInboxEvent.Handle(ctx, *createInboxEventCmd); err != nil {
		s.logger.Error("failed to handle createInboxEventCmd: ", err)
		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}

	// return success response to ZaloPay server
	return zalopay.NewCallbackResponse(1, "success").ToGrpcResponse(), nil
}
