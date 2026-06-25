package trmq

import (
	"context"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	updatePaymentAttemptStatus "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus"
	updatePaymentAttemptStatusDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus/dto"
	createPaymentTransaction "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction"
	createPaymentTransactionDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction/dto"
	createProcessedEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent"
	createProcessedEventDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent/dto"
	updatePurchase "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase"
	updatePurchaseDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase/dto"
	sdto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/dto"
	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
	createSubscription "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription"
	createSubscriptionDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription/dto"
	subscriptionhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/helper"

	getPlanById "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById"
	getPlanByIdDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById/dto"

	getPaymentAttemptByProvider "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttemptByProvider"
	getPaymentAttemptByProviderDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttemptByProvider/dto"
)

type PaymentSucceededEventHandler = consumer.ConsumerHandler

func NewPaymentSucceededEventHandler(
	logger logger.Logger,
	db dbcontracts.TxContextDb,

	updatePurchaseHandler updatePurchase.UpdatePurchaseHandler,
	updatePaymentAttemptStatusHandler updatePaymentAttemptStatus.UpdatePaymentAttemptStatusHandler,
	createPaymentTransactionHandler createPaymentTransaction.CreatePaymentTransactionHandler,
	createProcessedEventHandler createProcessedEvent.CreateProcessedEventHandler,
	createSubscriptionHandler createSubscription.CreateSubscriptionHandler,
	getPlanByIdHandler getPlanById.GetPlanByIdHandler,
	getPaymentAttemptByProviderHandler getPaymentAttemptByProvider.GetPaymentAttemptByProviderHandler,
) PaymentSucceededEventHandler {
	return func(ctx context.Context, consumeCtx types.MessageConsumeContext) error {

		eventMsg := event.NewPaymentSucceededEvent(consumeCtx.Message())

		consumeData, err := eventMsg.ParseData()
		if err != nil {
			return err
		}
		logger.Infof("Received PaymentSucceededEvent with payload: %+v", consumeData)

		// Start transaction
		return db.RunInTx(ctx, func(ctx context.Context, txCtx dbcontracts.TxContextDb) error {
			// Insert ProcessedEvent first for fast fail if the event has been processed before
			createProcessedEventCmd, err := createProcessedEvent.NewCreateProcessedEventCommand(
				createProcessedEventDto.CreateProcessedEventReqDto{
					ConsumerName: "purchases.consumer",
					MessageID:    eventMsg.GetMessageId(),
					ProcessedAt:  time.Now(),
				},
			)
			if err != nil {
				logger.Errorf("Failed to create command: %v", err)
				return err
			}
			_, err = createProcessedEventHandler.Handle(ctx, createProcessedEventCmd)
			if err != nil {
				logger.Errorf("Failed to create processed event: %v", err)
				return err
			}

			getPlanByIdQuery, err := getPlanById.NewGetPlanByIdQuery(getPlanByIdDto.GetPlanByIdReqDto{
				Id: consumeData.EmbedData.PlanID,
			})
			if err != nil {
				logger.Errorf("Failed to create query: %v", err)
				return err
			}
			getPlanByIdRes, err := getPlanByIdHandler.Handle(ctx, getPlanByIdQuery)
			if err != nil {
				logger.Errorf("Failed to get plan by ID: %v", err)
				return err
			}
			plan := getPlanByIdRes.GetData()

			// Insert Subscription status = 'pending' -> wait for user's role to be updated by Auth Service
			periodStart := time.Now()
			createSubscriptionCmd, err := createSubscription.NewCreateSubscriptionCommand(
				createSubscriptionDto.CreateSubscriptionReqDto{
					UserID:             consumeData.AppUser,
					PlanID:             plan.Id.String(),
					Status:             int32(billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_PENDING),
					CurrentPeriodStart: periodStart,
					CurrentPeriodEnd:   subscriptionhelper.CalculateSubscriptionEndDate(periodStart, plan.BillingInterval),
				},
			)
			if err != nil {
				logger.Errorf("Failed to create command: %v", err)
				return err
			}
			createSubscriptionRes, err := createSubscriptionHandler.Handle(ctx, createSubscriptionCmd)
			if err != nil {
				logger.Errorf("Failed to create subscription: %v", err)
				return err
			}

			// Get PaymentAttempt by ID
			getPAByProviderQuery, err := getPaymentAttemptByProvider.NewGetPaymentAttemptByProviderQuery(getPaymentAttemptByProviderDto.GetPaymentAttemptByProviderReqDto{
				Provider:              consumeData.Provider,
				ProviderTransactionId: consumeData.AppTransID,
			})
			if err != nil {
				logger.Errorf("Failed to create query: %v", err)
				return err
			}
			getPAByProviderRes, err := getPaymentAttemptByProviderHandler.Handle(ctx, getPAByProviderQuery)
			if err != nil {
				logger.Errorf("Failed to get payment attempt: %v", err)
				return err
			}
			paymentAttempt := getPAByProviderRes.GetData()

			// Update Purchase status = 'completed'
			subscriptionID := createSubscriptionRes.GetData().ID
			purchaseStatusCompleted := int32(billing_service.PurchaseStatus_PURCHASE_STATUS_COMPLETED)
			updatePurchaseCmd, err := updatePurchase.NewUpdatePurchaseCommand(
				updatePurchaseDto.UpdatePurchaseReqDto{
					ID:             paymentAttempt.PurchaseID,
					Status:         &purchaseStatusCompleted,
					SubscriptionID: &subscriptionID,
				},
			)
			if err != nil {
				logger.Errorf("Failed to create command: %v", err)
				return err
			}
			_, err = updatePurchaseHandler.Handle(ctx, updatePurchaseCmd)
			if err != nil {
				logger.Errorf("Failed to update purchase: %v", err)
				return err
			}

			// Update PurchaseAttempt status = 'success'
			updatePaymentAttemptStatusCmd, err := updatePaymentAttemptStatus.NewUpdatePaymentAttemptStatusCommand(
				updatePaymentAttemptStatusDto.UpdatePaymentAttemptStatusReqDto{
					ID:         paymentAttempt.ID,
					PurchaseID: paymentAttempt.PurchaseID,
					Status:     int32(billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_SUCCEEDED),
				},
			)
			if err != nil {
				logger.Errorf("Failed to create command: %v", err)
				return err
			}
			updatePARes, err := updatePaymentAttemptStatusHandler.Handle(ctx, updatePaymentAttemptStatusCmd)
			if err != nil {
				logger.Errorf("Failed to update payment attempt status: %v", err)
				return err
			}

			// Insert PaymentTransaction status = 'success'
			now := time.Now()
			createPTCmd, err := createPaymentTransaction.NewCreatePaymentTransactionCmd(
				createPaymentTransactionDto.CreatePaymentTransactionReqDto{
					PaymentAttemptID:  updatePARes.GetData().ID,
					Type:              int32(billing_service.PaymentTransactionType_PAYMENT_TRANSACTION_TYPE_PAYMENT),
					Status:            int32(billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_SUCCEEDED),
					ProviderReference: consumeData.AppTransID,
					ProcessedAt:       &now,
					Price:             *sdto.NewMoneyDto(consumeData.Amount, ""),
					Metadata:          consumeData.EmbedData.String(),
				},
			)
			if err != nil {
				logger.Errorf("Failed to create command: %v", err)
				return err
			}
			_, err = createPaymentTransactionHandler.Handle(ctx, *createPTCmd)
			if err != nil {
				logger.Errorf("Failed to create payment transaction: %v", err)
				return err
			}

			return nil
		})
	}
}
