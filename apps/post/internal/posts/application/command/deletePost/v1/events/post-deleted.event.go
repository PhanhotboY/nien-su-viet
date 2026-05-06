package event

import (
	"encoding/json"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/events"
)

type PostDeletedEvent struct {
	types.Message
}

func NewPostDeletedEvent(id string) (*PostDeletedEvent, error) {
	eventData := events.PostDeletedEvent{
		Id: id,
	}

	data, err := json.Marshal(eventData)
	if err != nil {
		return nil, err
	}
	return &PostDeletedEvent{
		Message: *types.NewMessage(data),
	}, nil
}
