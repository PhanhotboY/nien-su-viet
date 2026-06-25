package webhook

import (
	"context"
	"strings"

	"go.uber.org/fx"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/infrastructure/zalopay"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	grpcServer "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
	googleGrpc "google.golang.org/grpc"

	createInboxEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/createInboxEvent"
	createInboxEventDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/application/commands/createInboxEvent/dto"
)

type ZalopayGrpcServiceServer interface {
	HandleCallback(ctx context.Context, req *billing_service.CallbackPayload) (*billing_service.CallbackResponse, error)
}

var Module = fx.Module(
	"zalopayWebhook",

	fx.Provide(NewZalopayGrpcServiceServer),

	fx.Invoke(
		func(
			zalopayGrpcServer grpcServer.GrpcServer,
			logger logger.Logger,
			zpClient *zalopay.Client,

			zalopayGrpcServiceServer ZalopayGrpcServiceServer,

			createInboxEvent createInboxEvent.CreateInboxEventCmdHandler,
		) error {
			zalopayGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				billing_service.RegisterZaloPayWebhookServiceServer(server,
					zalopayGrpcServiceServer,
				)
			})
			return nil
		},
	))

type zalopayGrpcServiceServer struct {
	logger        logger.Logger
	zalopayClient *zalopay.Client

	createInboxEvent createInboxEvent.CreateInboxEventCmdHandler
}

func NewZalopayGrpcServiceServer(
	logger logger.Logger,
	zpClient *zalopay.Client,

	createInboxEvent createInboxEvent.CreateInboxEventCmdHandler,
) ZalopayGrpcServiceServer {
	return zalopayGrpcServiceServer{
		logger:           logger,
		zalopayClient:    zpClient,
		createInboxEvent: createInboxEvent,
	}
}

func (s zalopayGrpcServiceServer) HandleCallback(
	ctx context.Context,
	req *billing_service.CallbackPayload,
) (*billing_service.CallbackResponse, error) {
	payload := zalopay.NewCallbackPayload(req)

	// validate callback payload
	if err := s.zalopayClient.VerifyCallback([]byte(payload.Data), payload.Mac); err != nil {
		s.logger.Error("invalid callback payload: ", err)
		return zalopay.NewCallbackResponse(-1, "MAC not equal").ToGrpcResponse(), nil
	}

	var embedData event.EmbedData
	var itemData []event.Item
	if err := jsonUtils.UnmarshalJson(payload.GetData().EmbedData, &embedData); err != nil {
		s.logger.Error("failed to unmarshal embed data: ", err)
		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}
	if err := jsonUtils.UnmarshalJson(payload.GetData().Item, &itemData); err != nil {
		s.logger.Error("failed to unmarshal item data: ", err)
		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}
	paymentSucceededEvent := event.NewPaymentSucceededEvent(nil)
	if err := paymentSucceededEvent.SetData(&event.PaymentSucceededEventData{
		Provider:       "zalopay",
		AppID:          payload.GetData().AppID,
		AppTransID:     payload.GetData().AppTransID,
		AppTime:        payload.GetData().AppTime,
		Amount:         payload.GetData().Amount,
		AppUser:        payload.GetData().AppUser,
		EmbedData:      embedData,
		Item:           itemData,
		UseFeeAmount:   payload.GetData().UseFeeAmount,
		DiscountAmount: payload.GetData().DiscountAmount,
	}); err != nil {
		s.logger.Error("failed to set raw data for paymentSucceededEvent: ", err)
		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}
	// TODO: insert inbox event
	createInboxEventCmd, err := createInboxEvent.NewCreateInboxEventCmd(createInboxEventDto.CreateInboxEventReqDto{
		EventType:       utils.GetMessageName(event.NewPaymentSucceededEvent(nil)),
		Provider:        "zalopay",
		ExternalEventID: payload.GetData().AppTransID,
		Payload:         string(paymentSucceededEvent.GetData()),
		Signature:       payload.Mac,
	})
	if err != nil {
		s.logger.Error("failed to create createInboxEventCmd: ", err)
		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}
	if _, err := s.createInboxEvent.Handle(ctx, *createInboxEventCmd); err != nil {
		s.logger.Error("failed to handle createInboxEventCmd: ", err)
		// duplicate key means the event has been processed before, we can return success response to ZaloPay server
		if strings.Contains(err.Error(), "duplicate key") && strings.Contains(err.Error(), "idx_inbox_provider_event") {
			return zalopay.NewCallbackResponse(1, "success").ToGrpcResponse(), nil
		}

		return zalopay.NewCallbackResponse(0, "internal error").ToGrpcResponse(), nil
	}

	// return success response to ZaloPay server
	return zalopay.NewCallbackResponse(1, "success").ToGrpcResponse(), nil
}
