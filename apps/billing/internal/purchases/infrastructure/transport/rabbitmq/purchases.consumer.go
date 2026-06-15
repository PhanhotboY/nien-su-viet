package trmq

import (
	"context"
	"time"

	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"

	updatePurchase "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase"
	updatePurchaseDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchase/dto"

	updatePaymentAttemptStatus "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus"
	updatePaymentAttemptStatusDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/updatePaymentAttemptStatus/dto"

	createPaymentTransaction "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction"
	createPaymentTransactionDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/application/commands/createPaymentTransaction/dto"

	createProcessedEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent"
	createProcessedEventDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/processed_events/application/commands/createProcessedEvent/dto"

	createSubscription "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription"
	createSubscriptionDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription/dto"
)

var ConsumerModule = fx.Module(
	"purchasesRmqConsumerModule",
	fx.Invoke(NewPurchasesConsumer),
)

func NewPurchasesConsumer(
	b bus.RabbitmqBus,
	logger logger.Logger,
	db dbcontracts.TxContextDb,

	updatePurchaseHandler updatePurchase.UpdatePurchaseHandler,
	updatePaymentAttemptStatusHandler updatePaymentAttemptStatus.UpdatePaymentAttemptStatusHandler,
	createPaymentTransactionHandler createPaymentTransaction.CreatePaymentTransactionHandler,
	createProcessedEventHandler createProcessedEvent.CreateProcessedEventHandler,
	createSubscriptionHandler createSubscription.CreateSubscriptionHandler,
) error {
	b.ConnectConsumerHandler(
		event.NewPaymentSucceededEvent(nil),
		func(ctx context.Context, consumeCtx types.MessageConsumeContext) error {
			eventMsg := event.NewPaymentSucceededEvent(consumeCtx.Message())

			consumeData, err := eventMsg.ParseData()
			if err != nil {
				return nil
			}
			logger.Infof("Received PaymentSucceededEvent with payload: %+v", consumeData)

			// Start transaction
			return db.RunInTx(ctx, func(ctx context.Context, txCtx dbcontracts.TxContextDb) error {
				// Insert Subscription status = 'pending' -> wait for user's role to be updated by Auth Service
				createSubscriptionCmd, _ := createSubscription.NewCreateSubscriptionCommand(
					createSubscriptionDto.CreateSubscriptionReqDto{
						PlanID: consumeData.EmbedData.PlanID,
						Status: int32(billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_PENDING),
					},
				)
				createSubscriptionRes, err := createSubscriptionHandler.Handle(ctx, createSubscriptionCmd)
				if err != nil {
					logger.Errorf("Failed to create subscription: %v", err)
				}

				// Update Purchase status = 'completed'
				subscriptionID := createSubscriptionRes.GetData().ID
				purchaseStatusCompleted := int32(billing_service.PurchaseStatus_PURCHASE_STATUS_COMPLETED)
				updatePurchaseCmd, _ := updatePurchase.NewUpdatePurchaseCommand(
					updatePurchaseDto.UpdatePurchaseReqDto{
						ID:             consumeData.EmbedData.PurchaseID,
						Status:         &purchaseStatusCompleted,
						SubscriptionID: &subscriptionID,
					},
				)
				_, err = updatePurchaseHandler.Handle(ctx, updatePurchaseCmd)
				if err != nil {
					logger.Errorf("Failed to update purchase: %v", err)
					return err
				}

				// Update PurchaseAttempt status = 'success'
				updatePaymentAttemptStatusCmd, _ := updatePaymentAttemptStatus.NewUpdatePaymentAttemptStatusCommand(
					updatePaymentAttemptStatusDto.UpdatePaymentAttemptStatusReqDto{
						PurchaseID: consumeData.EmbedData.PurchaseID,
						Status:     int32(billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_SUCCEEDED),
					},
				)
				updatePARes, err := updatePaymentAttemptStatusHandler.Handle(ctx, updatePaymentAttemptStatusCmd)
				if err != nil {
					logger.Errorf("Failed to update payment attempt status: %v", err)
				}

				// Insert PaymentTransaction status = 'success'
				now := time.Now()
				createPTCmd, _ := createPaymentTransaction.NewCreatePaymentTransactionCmd(
					createPaymentTransactionDto.CreatePaymentTransactionReqDto{
						PaymentAttemptID:  updatePARes.GetData().ID,
						Type:              int32(billing_service.PaymentTransactionType_PAYMENT_TRANSACTION_TYPE_PAYMENT),
						Status:            int32(billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_SUCCEEDED),
						ProviderReference: consumeData.AppTransID,
						ProcessedAt:       &now,
					},
				)
				_, err = createPaymentTransactionHandler.Handle(ctx, *createPTCmd)
				if err != nil {
					logger.Errorf("Failed to create payment transaction: %v", err)
				}

				// Insert ProcessedEvent type='payment_succeeded', consumer_name=''
				createProcessedEventCmd, _ := createProcessedEvent.NewCreateProcessedEventCommand(
					createProcessedEventDto.CreateProcessedEventReqDto{
						ConsumerName: "purchases.consumer",
						MessageID:    eventMsg.GetMessageId(),
						ProcessedAt:  time.Now(),
					},
				)
				_, err = createProcessedEventHandler.Handle(ctx, createProcessedEventCmd)
				if err != nil {
					logger.Errorf("Failed to create processed event: %v", err)
				}

				return nil
			})
		},
	)

	return nil
}
