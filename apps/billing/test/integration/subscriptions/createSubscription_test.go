//go:build integration || subscriptions

package v1

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
	acmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription"
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription/dto"
	subscriptionhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/helper"

	oeEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
	oeRepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/repository"

	createPlanCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan"
	createPlanDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan/dto"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	planhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/helper"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/billing/test/integration/shared/helper"
)

func TestCreateSubscription(t *testing.T) {
	var (
		createSubscriptionHandler acmd.CreateSubscriptionHandler
		outboxEventRepo           oeRepo.OutboxEventDbRepo
		createPlanHandler         createPlanCmd.CreatePlanHandler

		logger testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &logger, &createSubscriptionHandler, &outboxEventRepo, &createPlanHandler)
	subscriptionCreatedEvent := event.NewSubscriptionCreatedEvent(nil)
	subscriptionCreatedEventType := utils.GetMessageName(subscriptionCreatedEvent)

	t.Run("Create subscription with all required fields", func(t *testing.T) {
		ctx := context.Background()
		trueValue := true
		billingInterval := int32(billing_service.BillingInterval_BILLING_INTERVAL_MONTH)

		createPlanCmd, err := createPlanCmd.NewCreatePlanCommand(createPlanDto.CreatePlanReqDto{
			Code: "premium",
			Name: "Premium Membership",
			Price: sdto.MoneyDto{
				Amount:   100_000,
				Currency: "VND",
			},
			BillingInterval: billingInterval,
			IsActive:        &trueValue,
		})
		if err != nil {
			logger.TestFatalf("Failed to create create plan command: %v", err)
		}

		createPlanRes, err := createPlanHandler.Handle(ctx, createPlanCmd)
		if err != nil {
			logger.TestFatalf("Failed to create plan: ", err)
		}

		planId := createPlanRes.GetData().ID

		periodStart := time.Now()
		createSubscriptionCmd, err := acmd.NewCreateSubscriptionCommand(adto.CreateSubscriptionReqDto{
			UserID:             "user-123",
			PlanID:             planId,
			Status:             int32(billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE),
			CurrentPeriodStart: periodStart,
			CurrentPeriodEnd:   subscriptionhelper.CalculateSubscriptionEndDate(periodStart, planhelper.ToEntityInterval(&billingInterval, entity.BILLING_INTERVAL_MONTH)),
			CancelAtPeriodEnd:  false,
		})
		if err != nil {
			logger.TestFatalf("Failed to create create subscription command: %v", err)
		}
		createSubscriptionRes, err := createSubscriptionHandler.Handle(ctx, createSubscriptionCmd)
		if err != nil {
			logger.TestFatalf("Failed to create subscription: ", err)
		}

		if createSubscriptionRes.GetData().ID == "" {
			logger.TestFatalf("Expected subscription ID to be generated, got empty string")
		}
		if createSubscriptionRes.GetData().Success != true {
			logger.TestFatalf("Expected subscription to be created successfully, got: ", createSubscriptionRes.GetData().Success)
		}

		if outboxEventRepo == nil {
			logger.TestFatalf("Expected outbox event repository to be initialized")
		}
		outboxEvents, err := outboxEventRepo.FindByEventType(ctx, subscriptionCreatedEventType)
		if err != nil {
			logger.TestFatalf("Failed to find outbox events by event type: ", err)
		}

		if len(outboxEvents) == 0 {
			logger.TestFatalf("Expected at least one outbox event, got 0")
		}

		eventIdx := slices.IndexFunc(outboxEvents, func(event *oeEntity.OutboxEvent) bool {
			return event.EventType == subscriptionCreatedEventType &&
				event.RetryCount == 0 &&
				event.AggregateID.String() == createSubscriptionRes.GetData().ID &&
				event.AggregateType == "subscription" &&
				event.Status == oeEntity.OUTBOX_EVENT_STATUS_PENDING
		})
		if eventIdx == -1 {
			logger.TestFatalf("Expected to find outbox event with correct data, but did not find one")
		}
	})

	t.Run("Create subscription with invalid period", func(t *testing.T) {
		ctx := context.Background()

		trueValue := true
		billingInterval := int32(billing_service.BillingInterval_BILLING_INTERVAL_MONTH)
		createPlanCmd, err := createPlanCmd.NewCreatePlanCommand(createPlanDto.CreatePlanReqDto{
			Code: "premium1",
			Name: "Premium Membership",
			Price: sdto.MoneyDto{
				Amount:   100_000,
				Currency: "VND",
			},
			BillingInterval: billingInterval,
			IsActive:        &trueValue,
		})
		if err != nil {
			logger.TestFatalf("Failed to create create plan command: %v", err)
		}

		createPlanRes, err := createPlanHandler.Handle(ctx, createPlanCmd)
		if err != nil {
			logger.TestFatalf("Failed to create plan: ", err)
		}

		planId := createPlanRes.GetData().ID

		periodStart := time.Now()
		_, err = acmd.NewCreateSubscriptionCommand(adto.CreateSubscriptionReqDto{
			UserID:             "user-123",
			PlanID:             planId,
			Status:             int32(billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE),
			CurrentPeriodStart: periodStart,
			CurrentPeriodEnd:   time.Now().Add(-time.Hour), // Invalid period end (before start)
			CancelAtPeriodEnd:  false,
		})
		if err == nil {
			logger.TestFatalf("Expected error when creating subscription command with invalid period, but got no error")
		}

	})

	t.Run("Create subscription with not found plan", func(t *testing.T) {
		ctx := context.Background()
		billingInterval := int32(billing_service.BillingInterval_BILLING_INTERVAL_MONTH)

		periodStart := time.Now()
		createSubscriptionCmd, err := acmd.NewCreateSubscriptionCommand(adto.CreateSubscriptionReqDto{
			UserID:             "user-123",
			PlanID:             uuid.New().String(),
			Status:             int32(billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE),
			CurrentPeriodStart: periodStart,
			CurrentPeriodEnd:   subscriptionhelper.CalculateSubscriptionEndDate(periodStart, planhelper.ToEntityInterval(&billingInterval, entity.BILLING_INTERVAL_MONTH)),
			CancelAtPeriodEnd:  false,
		})
		if err != nil {
			logger.TestFatalf("Failed to create create subscription command: %v", err)
		}
		_, err = createSubscriptionHandler.Handle(ctx, createSubscriptionCmd)
		if err == nil {
			logger.TestFatalf("Expected error when creating subscription with not found plan, but got no error")
		}
	})
}
