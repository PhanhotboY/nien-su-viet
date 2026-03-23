package producer

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
)

type Producer interface {
	PublishMessage(ctx context.Context, message types.IMessage) error
	PublishMessageWithTopicName(ctx context.Context, message types.IMessage, topicOrExchangeName string) error
	IsProduced(h func(message types.IMessage))
}
