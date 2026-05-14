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

type Message struct {
	MessageId string    `json:"messageId,omitempty"`
	Created   time.Time `json:"created"`
	isMessage bool
	Data      json.RawMessage `json:"data,omitempty"`
}

func NewMessage(data json.RawMessage) *Message {
	return &Message{Created: time.Now(), Data: data}
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
