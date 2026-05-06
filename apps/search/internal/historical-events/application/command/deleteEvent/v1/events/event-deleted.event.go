package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type HistoricalEventDeletedEvent struct {
	types.Message
}

func NewHistoricalEventDeletedEvent() HistoricalEventDeletedEvent {
	return HistoricalEventDeletedEvent{
		Message: types.Message{},
	}
}
