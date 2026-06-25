package trmq

import (
	"context"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	updateSubscriptionStatus "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/updateSubscriptionStatus"
)

type UserAddedToOrganizationEventHandler = consumer.ConsumerHandler

func NewUserAddedToOrganizationEventHandler(
	logger logger.Logger,
	db dbcontracts.TxContextDb,
	updateSubscriptionStatusHandler updateSubscriptionStatus.UpdateSubscriptionStatusHandler,
) UserAddedToOrganizationEventHandler {
	return func(ctx context.Context, consumeCtx types.MessageConsumeContext) error {
		eventMsg := event.NewUserAddedToOrganizationEvent(consumeCtx.Message())

		consumeData, err := eventMsg.ParseData()
		if err != nil {
			return err
		}
		logger.Infof("Received UserAddedToOrganizationEvent with payload: %+v", consumeData)

		if !consumeData.IsPremium {
			logger.Infof("Organization %s is not premium. No subscription to update.", consumeData.OrganizationId)
			return nil
		}

		// Start transaction
		return db.RunInTx(ctx, func(ctx context.Context, txCtx dbcontracts.TxContextDb) error {
			// Update subscription status to active
			// cmd, err := updateSubscriptionStatus.NewUpdateSubscriptionStatusCommand(
			// 	updateSubscriptionStatusDto.UpdateSubscriptionStatusReqDto{
			// 		ID:     consumeData.OrganizationId,
			// 		Status: billing_service.SubscriptionStatus_ACTIVE,
			// 	},
			// )
			return nil
		})
	}
}
