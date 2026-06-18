package acmd

import (
	"context"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/application/commands/createSubscription/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"

	createOutboxEvent "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/application/commands/createOutboxEvent"
	createOutboxEventDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/application/commands/createOutboxEvent/dto"
)

type CreateSubscriptionHandler interface {
	grpcTypes.GrpcHandler[*CreateSubscriptionCommand, adto.CreateSubscriptionResDto]
}

type createSubscriptionHandler struct {
	logger           logger.Logger
	db               dbcontracts.TxContextDb
	subscriptionRepo drepo.SubscriptionDbRepo

	createOutboxEventHandler createOutboxEvent.CreateOutboxEventHandler
}

func NewCreateSubscriptionHandler(
	l logger.Logger,
	db dbcontracts.TxContextDb,
	subscriptionRepo drepo.SubscriptionDbRepo,
	createOutboxEventHandler createOutboxEvent.CreateOutboxEventHandler,
) CreateSubscriptionHandler {
	return &createSubscriptionHandler{l, db, subscriptionRepo, createOutboxEventHandler}
}

func (h *createSubscriptionHandler) Handle(ctx context.Context, command *CreateSubscriptionCommand) (adto.CreateSubscriptionResDto, error) {
	h.logger.Info("Handling CreateSubscriptionCommand", "command", command.CreateSubscriptionReqDto)

	subscriptionId := ""
	err := h.db.RunInTx(ctx, func(ctx context.Context, txCtx dbcontracts.TxContextDb) error {
		id, err := h.subscriptionRepo.CreateSubscription(ctx, command.MapToEntity())
		if err != nil {
			h.logger.Errorf("Failed to create subscription: %v", err)
			return grpcerrors.NewInternalServerGrpcError("Failed to create subscription", "CreateSubscriptionHandler")
		}

		subscription, err := h.subscriptionRepo.GetSubscriptionByID(ctx, id)
		if err != nil {
			h.logger.Errorf("Failed to retrieve created subscription: %v", err)
			return grpcerrors.NewInternalServerGrpcError("Failed to retrieve created subscription", "CreateSubscriptionHandler")
		}

		subscriptionCreatedEvent := event.NewSubscriptionCreatedEvent(nil)
		err = subscriptionCreatedEvent.SetData(*subscription)
		if err != nil {
			h.logger.Errorf("Failed to set data for SubscriptionCreatedEvent: %v", err)
			return grpcerrors.NewInternalServerGrpcError("Failed to create subscription created event", "CreateSubscriptionHandler")
		}
		createOutboxEventCmd, err := createOutboxEvent.NewCreateOutboxEventCommand(
			createOutboxEventDto.CreateOutboxEventReqDto{
				EventType:     utils.GetMessageName(subscriptionCreatedEvent),
				Payload:       string(subscriptionCreatedEvent.GetData()),
				AggregateType: "subscription",
				AggregateID:   id,
			})
		if err != nil {
			h.logger.Errorf("Failed to create CreateOutboxEventCommand: %v", err)
			return grpcerrors.NewInternalServerGrpcError("Failed to create outbox event command", "CreateSubscriptionHandler")
		}
		_, err = h.createOutboxEventHandler.Handle(ctx, createOutboxEventCmd)
		if err != nil {
			h.logger.Errorf("Failed to handle CreateOutboxEventCommand: %v", err)
			return grpcerrors.NewInternalServerGrpcError("Failed to create outbox event", "CreateSubscriptionHandler")
		}

		subscriptionId = id
		return nil
	})
	if (err != nil) || (subscriptionId == "") {
		h.logger.Errorf("Transaction failed: %v", err)
		return nil, err
	}

	h.logger.Infof("Subscription created successfully with ID: %s", subscriptionId)
	return adto.NewCreateSubscriptionResDto(subscriptionId, true, "Subscription created successfully"), nil
}
