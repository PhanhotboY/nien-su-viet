package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type HistoricalEventCreatedEvent struct {
	types.Message
}

func NewHistoricalEventCreatedEvent() HistoricalEventCreatedEvent {
	return HistoricalEventCreatedEvent{
		Message: types.Message{},
	}
}
