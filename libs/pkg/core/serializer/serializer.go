package serializer

import "github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"

type SerializedMessage struct {
	Data        []byte
	ContentType string
}

type MessageSerializer interface {
	Serialize(message types.IMessage) (*SerializedMessage, error)
	Deserialize(data []byte, contentType string) (types.IMessage, error)
	ContentType() string
}
