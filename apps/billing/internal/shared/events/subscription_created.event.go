package event

import (
	"encoding/json"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
)

type SubscriptionCreatedEventData struct {
	ID string `json:"id"`

	UserID string `json:"user_id"`

	PlanID string `json:"plan_id"`

	Status int32 `json:"status"`

	// Very common query: find the current subscription that covers now
	CurrentPeriodStart time.Time `json:"current_period_start"`
	CurrentPeriodEnd   time.Time `json:"current_period_end"`

	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`

	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	ExpiredAt  *time.Time `json:"expired_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubscriptionCreatedEvent interface {
	MessageEvent[*SubscriptionCreatedEventData]
}

type subscriptionCreatedEvent struct {
	*types.Message
}

func NewSubscriptionCreatedEvent(msg types.IMessage) SubscriptionCreatedEvent {
	return &subscriptionCreatedEvent{
		Message: msg.(*types.Message),
	}
}

func (e *subscriptionCreatedEvent) SetData(data string) error {
	if err := dtoUtil.ValidateStruct(data, &SubscriptionCreatedEventData{}); err != nil {
		return err
	}
	e.Data = json.RawMessage(data)
	return nil
}

func (e *subscriptionCreatedEvent) ParseData() (*SubscriptionCreatedEventData, error) {
	data := new(SubscriptionCreatedEventData)
	if err := jsonUtils.UnmarshalJson(string(e.Data), data); err != nil {
		return nil, err
	}

	return data, nil
}
