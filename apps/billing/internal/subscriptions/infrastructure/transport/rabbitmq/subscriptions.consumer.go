package trmq

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq/bus"

	event "github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/events"
)

func NewPurchasesConsumer(
	b bus.RabbitmqBus,
	logger logger.Logger,
	db dbcontracts.TxContextDb,

	userAddedToOrganizationEventHandler UserAddedToOrganizationEventHandler,
) error {
	b.ConnectConsumerHandler(
		event.NewUserAddedToOrganizationEvent(nil),
		userAddedToOrganizationEventHandler,
	)

	return nil
}
