package event

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type PostDeletedEvent struct {
	types.Message
}

func NewPostDeletedEvent() PostDeletedEvent {
	return PostDeletedEvent{
		Message: types.Message{},
	}
}
