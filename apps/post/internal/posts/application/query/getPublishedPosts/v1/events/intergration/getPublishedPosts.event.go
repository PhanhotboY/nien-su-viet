package events

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	uuid "github.com/satori/go.uuid"
)

type GetPublishedPostsEvent struct {
	*types.Message
}

func NewGetPublishedPostsEvent() types.IMessage {
	return &GetPublishedPostsEvent{
		Message: types.NewMessage(uuid.NewV4().String(), "posts.created", nil),
	}
}
