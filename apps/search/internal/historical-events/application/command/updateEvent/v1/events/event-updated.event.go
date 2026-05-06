package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type HistoricalEventUpdatedEvent struct {
	types.Message
}

func NewHistoricalEventUpdatedEvent() HistoricalEventUpdatedEvent {
	return HistoricalEventUpdatedEvent{
		Message: types.Message{},
	}
}
