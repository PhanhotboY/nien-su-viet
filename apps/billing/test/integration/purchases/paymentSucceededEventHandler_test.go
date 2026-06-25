//go:build integration || purchases

package v1

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	testhelper "github.com/phanhotboy/nien-su-viet/apps/billing/test/integration/shared/helper"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	paymentSucceededEventHandler "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/infrastructure/transport/rabbitmq"

	createPurchaseCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase"
	createPurchaseDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase/dto"

	createPlanCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan"
	createPlanDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan/dto"

	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
)

func TestPaymentSucceededEventHandler(t *testing.T) {
	var (
		logger testlogger.TestLogger

		paymentSucceededHandler paymentSucceededEventHandler.PaymentSucceededEventHandler
		createPurchaseHandler   createPurchaseCmd.CreatePurchaseHandler
		createPlanHandler       createPlanCmd.CreatePlanHandler
	)

	testhelper.GetDIServices(t, &logger, &paymentSucceededHandler, &createPurchaseHandler, &createPlanHandler)

	paymentSucceededEvent := event.NewPaymentSucceededEvent(nil)
	ctx := context.Background()
	trueValue := true
	userId := "user-123"

	createPlanCmd, err := createPlanCmd.NewCreatePlanCommand(createPlanDto.CreatePlanReqDto{
		Code:     "test-payment-succeeded-event-handler-plan",
		Name:     "Test payment succeeded event handler",
		Price:    *sdto.NewMoneyDto(10000, "VND"),
		IsActive: &trueValue,
	})
	if err != nil {
		logger.TestFatalf("Expect create plan command to be created, but got error: ", err)
	}

	createPlanRes, err := createPlanHandler.Handle(ctx, createPlanCmd)
	if err != nil {
		logger.TestFatalf("Expect plan to be created, but got error: ", err)
	}

	createPurchaseCmd, err := createPurchaseCmd.NewCreatePurchaseCommand(createPurchaseDto.CreatePurchaseReqDto{
		PlanID:         createPlanRes.GetData().ID,
		UserID:         userId,
		IdempotencyKey: uuid.NewString(),
	})
	if err != nil {
		logger.TestFatalf("Expect create purchase command to be created, but got error: ", err)
	}

	createPurchaseRes, err := createPurchaseHandler.Handle(ctx, createPurchaseCmd)
	if err != nil {
		logger.TestFatalf("Expect purchase to be created, but got error: ", err)
	}

	logger.Infof("purchase created with id %s", createPurchaseRes.GetData().Id)
	paymentAttempt := createPurchaseRes.GetData().PaymentAttempt.GetData()
	err = paymentSucceededEvent.SetData(&event.PaymentSucceededEventData{
		Provider:   paymentAttempt.Provider,
		AppID:      1,
		AppTransID: paymentAttempt.ProviderTransactionID,
		AppTime:    time.Now().UnixMilli(),
		Amount:     paymentAttempt.Amount,
		AppUser:    userId,
		EmbedData: event.EmbedData{
			PlanID:     createPlanRes.GetData().ID,
			PurchaseID: createPurchaseRes.GetData().Id,
		},
		Item:           []event.Item{},
		UseFeeAmount:   0,
		DiscountAmount: 0,
	})
	if err != nil {
		logger.TestFatalf("Expect event data to be set, but got error: ", err)
	}
	consumeCtx := types.NewMessageConsumeContext(paymentSucceededEvent, "applicatin/json", time.Now(), 1, "")

	t.Run("Handle payment succeeded event", func(t *testing.T) {
		err := paymentSucceededHandler(ctx, consumeCtx)
		if err != nil {
			logger.TestFatalf("Expected successfully handle PaymentSucceededEvent, but got error: %v", err)
		}
	})

	t.Run("Handle duplicated payment succeeded event", func(t *testing.T) {
		err := paymentSucceededHandler(ctx, consumeCtx)
		if err == nil {
			logger.TestFatalf("Expected failed to handle PaymentSucceededEvent, but got no error")
		}
	})
}
