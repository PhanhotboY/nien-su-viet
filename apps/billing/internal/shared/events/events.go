package event

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
)

type MessageEvent[T any] interface {
	types.IMessage
	types.IMessageParser[T]
}
