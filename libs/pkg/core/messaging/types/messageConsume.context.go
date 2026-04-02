package types

import (
	"time"
)

type MessageConsumeContext interface {
	Message() IMessage
	ContentType() string
	Timestamp() time.Time
	DeliveryTag() uint64
	CorrelationId() string
}

type messageConsumeContext struct {
	message       IMessage
	contentType   string
	timestamp     time.Time
	deliveryTag   uint64
	correlationID string
}

func NewMessageConsumeContext(
	message IMessage,
	contentType string,
	timestamp time.Time,
	deliveryTag uint64,
	correlationID string,
) MessageConsumeContext {
	return &messageConsumeContext{
		message:       message,
		contentType:   contentType,
		timestamp:     timestamp,
		deliveryTag:   deliveryTag,
		correlationID: correlationID,
	}
}

func (m *messageConsumeContext) ContentType() string {
	return m.contentType
}

func (m *messageConsumeContext) Timestamp() time.Time {
	return m.timestamp
}

func (m *messageConsumeContext) DeliveryTag() uint64 {
	return m.deliveryTag
}

func (m *messageConsumeContext) CorrelationId() string {
	return m.correlationID
}

func (m *messageConsumeContext) Message() IMessage {
	return m.message
}
