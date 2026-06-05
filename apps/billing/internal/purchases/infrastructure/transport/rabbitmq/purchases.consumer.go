package trmq

import (
	"context"

	"go.uber.org/fx"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"

	updatePurchaseStatus "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/updatePurchaseStatus/commands"
)

var ConsumerModule = fx.Module(
	"purchasesRmqConsumerModule",
	fx.Invoke(NewPurchasesConsumer),
)

func NewPurchasesConsumer(
	b bus.RabbitmqBus,
	logger logger.Logger,

	updatePurchaseStatusHandler updatePurchaseStatus.UpdatePurchaseStatusHandler,
) error {
	b.ConnectConsumerHandler(
		event.NewPaymentSucceededEvent(),
		func(ctx context.Context, consumeCtx types.MessageConsumeContext) error {
			logger.Infof("Received PaymentSucceededEvent with payload: %v", string(consumeCtx.Message().GetData()))

			updatePurchaseStatusCmd, err := updatePurchaseStatus.NewUpdatePurchaseStatusCommand(consumeCtx.Message().GetData())
			if err != nil {
				logger.Errorf("Failed to convert consume context to UpdatePurchaseStatusCommand: %v", err)
				return err
			}
			_, err = updatePurchaseStatusHandler.Handle(ctx, updatePurchaseStatusCmd)
			if err != nil {
				logger.Errorf("Failed to handle UpdatePurchaseStatusCommand: %v", err)
				return err
			}
			return nil
		},
	)

	return nil
}
