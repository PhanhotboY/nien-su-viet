package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type PostUpdatedEvent struct {
	types.Message
}

func NewPostUpdatedEvent() PostUpdatedEvent {
	return PostUpdatedEvent{
		Message: types.Message{},
	}
}
