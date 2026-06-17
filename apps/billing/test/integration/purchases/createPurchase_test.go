//go:build integration

package v1

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	acmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase"
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase/dto"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"

	getPurchaseQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase"
	getPurchaseDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/queries/getPurchase/dto"

	createPlanCmd "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan"
	createPlanDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/commands/createPlan/dto"

	PAEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/billing/test/integration/shared/helper"
)

func TestCreatePurchase(t *testing.T) {
	var (
		createPurchaseHandler acmd.CreatePurchaseHandler
		getPurchaseHandler    getPurchaseQuery.GetPurchaseHandler
		createPlanHandler     createPlanCmd.CreatePlanHandler

		log testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPurchaseHandler, &getPurchaseHandler, &createPlanHandler, &log)

	t.Run("Create purchase with all required fields", func(t *testing.T) {
		ctx := context.Background()
		trueValue := true

		createPlanCmd, err := createPlanCmd.NewCreatePlanCommand(createPlanDto.CreatePlanReqDto{
			Code: "premium",
			Name: "Premium Membership",
			Price: sdto.MoneyDto{
				Amount:   100_000,
				Currency: "VND",
			},
			BillingInterval: int32(billing_service.BillingInterval_BILLING_INTERVAL_MONTH),
			IsActive:        &trueValue,
		})
		if err != nil {
			log.TestFatalf("Failed to create create plan command: %v", err)
		}

		createPlanRes, err := createPlanHandler.Handle(ctx, createPlanCmd)
		if err != nil {
			log.TestFatalf("Failed to create plan: ", err)
		}

		planId := createPlanRes.GetData().ID

		cmd, err := acmd.NewCreatePurchaseCommand(
			&adto.CreatePurchaseReqDto{
				UserID:         "user-123",
				PlanID:         planId,
				IdempotencyKey: uuid.NewString(),
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create purchase command: %v", err)
		}

		res, err := createPurchaseHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected purchase to be created, but got error: %v", err)
		}

		// Verify response contains ID
		if res.GetData().Id == "" {
			log.TestErrorf("Expected purchase ID to be generated, but got empty ID")
		}

		// Get the purchase to verify it was created correctly
		getCmd, err := getPurchaseQuery.NewGetPurchaseQuery(getPurchaseDto.GetPurchaseReqDto{
			PurchaseId: res.GetData().Id,
		})
		if err != nil {
			log.TestFatalf("Failed to create get purchase query: %v", err)
		}

		getRes, err := getPurchaseHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get created purchase, but got error: %v", err)
		}

		// Verify purchase is not published by default
		if _, err = uuid.Parse(getRes.GetData().ID); err != nil {
			log.TestErrorf("Expected purchase id to be uuid, but got: %s", getRes.GetData().ID)
		}

		paymentAttempt := res.GetData().PaymentAttempt.GetData()
		if paymentAttempt.Status != PAEntity.PAYMENT_ATTEMPT_STATUS_PENDING {
			log.TestErrorf("Expected payment attempt status to be '%s', but got: %s", PAEntity.PAYMENT_ATTEMPT_STATUS_PENDING, paymentAttempt.Status)
		}
		if paymentAttempt.PurchaseID != getRes.GetData().ID {
			log.TestErrorf("Expected payment attempt to be linked to purchase ID '%s', but got: %s", getRes.GetData().ID, paymentAttempt.PurchaseID)
		}

		if paymentAttempt.CheckoutURL == "" {
			log.TestErrorf("Expected payment attempt to have a checkout URL, but got empty string")
		}
	})
}
