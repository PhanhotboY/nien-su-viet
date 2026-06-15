package types

import (
	"encoding/json"
	"time"
)

type IMessage interface {
	GetMessageId() string
	GetCreated() time.Time
	GetData() json.RawMessage
}

type IMessageParser[T any] interface {
	ParseData() (T, error)
}

type Message struct {
	MessageId string    `json:"messageId,omitempty"`
	Created   time.Time `json:"created"`
	isMessage bool
	Data      json.RawMessage `json:"data,omitempty"`
}

func NewMessage(id string, data json.RawMessage) IMessage {
	return &Message{MessageId: id, Created: time.Now(), Data: data}
}

func (m *Message) GetMessageId() string {
	return m.MessageId
}

func (m *Message) GetCreated() time.Time {
	return m.Created
}

func (m *Message) GetData() json.RawMessage {
	return m.Data
}
