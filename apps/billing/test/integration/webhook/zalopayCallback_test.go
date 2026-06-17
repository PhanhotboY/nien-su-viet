//go:build integration || webhook

package v1

import (
	"context"
	"slices"
	"testing"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/billing/test/integration/shared/helper"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	ieEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"
	ieRepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/repository"
	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/webhook"
)

func TestZaloPayCallback(t *testing.T) {
	var (
		logger testlogger.TestLogger

		data              = "{\"app_id\":2553,\"app_trans_id\":\"200904_2553_1598435687208\",\"app_time\":1599189392817,\"app_user\":\"demo\",\"amount\":10000,\"embed_data\":\"{\\\"merchantinfo\\\":\\\"embeddata123\\\",\\\"promotioninfo\\\":\\\"\\\"}\",\"item\":\"[{\\\"itemid\\\":\\\"knb\\\",\\\"itemname\\\":\\\"kim nguyen bao\\\",\\\"itemprice\\\":198400,\\\"itemquantity\\\":1}]\",\"zp_trans_id\":200904000000389,\"server_time\":1599189413498,\"channel\":38,\"merchant_user_id\":\"7ZMSl3nEg5sOUJzOLSoUFT8xKNQVaLOLXHB--8Eytqc\",\"user_fee_amount\":0,\"discount_amount\":0}"
		mac               = "d8d33baf449b31d7f9b94fa50d7c942c08cd4d83f28fa185557da21acb104f67"
		paymentType int32 = 1
		eventType         = utils.GetMessageName(event.NewPaymentSucceededEvent(nil))

		zalopayGrpcServiceServer webhook.ZalopayGrpcServiceServer
		inboxEventRepo           ieRepo.InBoxEventDbRepo
	)

	testhelper.GetDIServices(t, &logger, &zalopayGrpcServiceServer, &inboxEventRepo)

	t.Run("Handle Success ZaloPay callback", func(t *testing.T) {
		ctx := context.Background()

		res, err := zalopayGrpcServiceServer.HandleCallback(ctx, &billing_service.CallbackPayload{
			Data: data,
			Mac:  mac,
			Type: paymentType,
		})
		if err != nil {
			logger.TestFatalf("Failed to handle ZaloPay callback: ", err)
		}

		if res == nil {
			logger.TestFatalf("Expected non-nil response, got nil")
		}

		if res.ReturnCode != 1 {
			logger.TestFatalf("Expected return code 1, got %d", res.ReturnCode)
		}

		inboxEvents, err := inboxEventRepo.FindByEventType(ctx, eventType)
		if err != nil {
			logger.TestFatalf("Failed to find inbox events by event type: ", err)
		}

		if len(inboxEvents) == 0 {
			logger.TestFatalf("Expected at least one inbox event, got 0")
		}

		eventIdx := slices.IndexFunc(inboxEvents, func(event *ieEntity.InboxEvent) bool {
			return event.EventType == eventType &&
				event.ExternalEventID == "200904_2553_1598435687208" &&
				event.Signature == mac &&
				event.RetryCount == 0 &&
				event.Status == ieEntity.INBOX_EVENT_STATUS_PENDING
		})
		if eventIdx == -1 {
			logger.TestFatalf("Expected to find inbox event with correct data, but did not find one")
		}
	})

	t.Run("Handle Invalid signature ZaloPay callback", func(t *testing.T) {
		ctx := context.Background()
		invalidMac := mac + "invalid"

		res, err := zalopayGrpcServiceServer.HandleCallback(ctx, &billing_service.CallbackPayload{
			Data: data,
			Mac:  invalidMac,
			Type: paymentType,
		})
		if err != nil {
			logger.TestFatalf("Failed to handle ZaloPay callback: ", err)
		}

		if res == nil {
			logger.TestFatalf("Expected non-nil response, got nil")
		}

		if res.ReturnCode != -1 {
			logger.TestFatalf("Expected return code -1, got %d", res.ReturnCode)
		}

		inboxEvents, err := inboxEventRepo.FindByEventType(ctx, eventType)
		if err != nil {
			logger.TestFatalf("Failed to find inbox events by event type: ", err)
		}

		if len(inboxEvents) == 0 {
			// No inbox events should be created for invalid callback
			return
		}

		eventIdx := slices.IndexFunc(inboxEvents, func(event *ieEntity.InboxEvent) bool {
			return event.EventType == eventType &&
				event.ExternalEventID == "200904_2553_1598435687208" &&
				event.Signature == invalidMac &&
				event.RetryCount == 0 &&
				event.Status == ieEntity.INBOX_EVENT_STATUS_PENDING
		})
		if eventIdx != -1 {
			logger.TestFatalf("Expected not to find inbox event with correct data for invalid callback, but found one")
		}
	})
}
