package adto // application command dto

import "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"

type CreateInboxEventReqDto struct {
	EventType       string `json:"event_type" validate:"required"`
	Provider        string `json:"provider" validate:"required"`
	ExternalEventID string `json:"external_event_id" validate:"required"`

	Payload   string `json:"payload" validate:"required"`
	Signature string `json:"signature"`
}

func (d *CreateInboxEventReqDto) MapToEntity() *entity.InboxEvent {
	return &entity.InboxEvent{
		EventType:       d.EventType,
		Provider:        d.Provider,
		ExternalEventID: d.ExternalEventID,
		Payload:         []byte(d.Payload),
		Signature:       d.Signature,
		Status:          entity.INBOX_EVENT_STATUS_PENDING,
	}
}
