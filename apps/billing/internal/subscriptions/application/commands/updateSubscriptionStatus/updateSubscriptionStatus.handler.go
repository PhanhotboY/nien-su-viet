package acmd

import (
	"context"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/updateSubscriptionStatus/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/repository"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

type UpdateSubscriptionStatusHandler interface {
	grpcTypes.GrpcHandler[*UpdateSubscriptionStatusCommand, *adto.UpdateSubscriptionStatusResDto]
}

type updateSubscriptionStatusHandler struct {
	logger logger.Logger
	cache  drepo.SubscriptionCacheRepo
	db     drepo.SubscriptionDbRepo
}

func NewUpdateSubscriptionStatusHandler(l logger.Logger, cache drepo.SubscriptionCacheRepo, db drepo.SubscriptionDbRepo) UpdateSubscriptionStatusHandler {
	return &updateSubscriptionStatusHandler{l, cache, db}
}

func (h *updateSubscriptionStatusHandler) Handle(ctx context.Context, command *UpdateSubscriptionStatusCommand) (*adto.UpdateSubscriptionStatusResDto, error) {
	subscriptionId, err := h.db.UpdateSubscription(ctx, command.ID, command.MapToEntity())
	if err != nil {
		h.logger.Errorf("Failed to update subscription status: %v", err)
		return nil, err
	}

	// Invalidate cache after creating subscription
	err = h.cache.DeleteAllSubscriptions(ctx)
	if err != nil {
		// Todo: Handle cache deletion failure
		h.logger.Warnf("Failed to delete all subscriptions cache after updating subscription status: %v", err)
	}

	return adto.NewUpdateSubscriptionStatusResDto(subscriptionId, true, "Subscription status updated successfully"), nil
}
