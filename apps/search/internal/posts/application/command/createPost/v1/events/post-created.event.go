package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type PostCreatedEvent struct {
	types.Message
}

func NewPostCreatedEvent() PostCreatedEvent {
	return PostCreatedEvent{
		Message: types.Message{},
	}
}
