package types

import (
	"encoding/json"
	"time"
)

type IMessage interface {
	GetMessageId() string
	GetCreated() time.Time
	GetData() json.RawMessage
	GetPattern() string
	SetPattern(string)
}

type IMessageParser[E any, T any] interface {
	SetRawData(string) error
	SetData(E) error
	ParseData() (T, error)
}

// EventEnvelope compatible
type Message struct {
	// for Nest.js microservices RMQ transport compatibility
	Pattern string `json:"pattern,omitempty"`
	// Outbox event id
	MessageId  string         `json:"event_id,omitempty"`
	Created    time.Time      `json:"occurred_at"`
	Attributes map[string]any `json:"attributes,omitempty"` // for custom attributes
	// Entity's name that incurred the event, e.g., Order, User, etc.
	AggregateType string `json:"aggregate_type,omitempty"`
	// Entity's ID that incurred the event, e.g., Order ID, User ID, etc.
	AggregateId string          `json:"aggregate_id,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"` // parsed to Nest.js' @Payload() decorator
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

func (m *Message) GetPattern() string {
	return m.Pattern
}

func (m *Message) SetPattern(pattern string) {
	m.Pattern = pattern
}
