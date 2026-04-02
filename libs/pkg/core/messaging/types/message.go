package types

import (
	"encoding/json"
	"time"
)

type IMessage interface {
	GetMessageId() string
	GetCreated() time.Time
	GetData() json.RawMessage
	GetRoutingKey() string
}

type Message struct {
	MessageId  string    `json:"messageId,omitempty"`
	Created    time.Time `json:"created,omitempty"`
	isMessage  bool
	Data       json.RawMessage `json:"data,omitempty"`
	RoutingKey string          `json:"pattern,omitempty"`
}

func NewMessage(messageId string, routingKey string, data json.RawMessage) *Message {
	return &Message{MessageId: messageId, Created: time.Now(), Data: data, RoutingKey: routingKey}
}

func (m Message) GetMessageId() string {
	return m.MessageId
}

func (m Message) GetCreated() time.Time {
	return m.Created
}

func (m Message) GetData() json.RawMessage {
	return m.Data
}

func (m Message) GetRoutingKey() string {
	return m.RoutingKey
}
