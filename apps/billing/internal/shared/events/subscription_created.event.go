package event

import (
	"encoding/json"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	eventspb "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/events"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"

	subscriptionEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
	subscriptionhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/helper"
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
	MessageEvent[subscriptionEntity.Subscription, *eventspb.EventEnvelope_SubscriptionCreated]
}

type subscriptionCreatedEvent struct {
	*types.Message
}

func NewSubscriptionCreatedEvent(msg types.IMessage) SubscriptionCreatedEvent {
	message := new(types.Message)
	if m, ok := msg.(*types.Message); ok {
		message = m
	}

	return &subscriptionCreatedEvent{
		Message: message,
	}
}

func (e *subscriptionCreatedEvent) SetRawData(data string) error {
	if err := dtoUtil.ValidateStruct(data, &eventspb.EventEnvelope_SubscriptionCreated{}); err != nil {
		return err
	}
	e.Data = json.RawMessage(data)
	return nil
}

func (e *subscriptionCreatedEvent) SetData(data subscriptionEntity.Subscription) error {
	eventData := &eventspb.EventEnvelope_SubscriptionCreated{
		SubscriptionCreated: &eventspb.SubscriptionCreated{
			Subscription: &eventspb.SubscriptionSnapshot{
				SubscriptionId:     data.ID.String(),
				UserId:             data.UserID,
				PlanId:             data.PlanID.String(),
				Status:             subscriptionhelper.ToGrpcStatus(data.Status),
				CurrentPeriodStart: grpcUtils.TimeToTimestamp(&data.CurrentPeriodStart),
				CurrentPeriodEnd:   grpcUtils.TimeToTimestamp(&data.CurrentPeriodEnd),
				CancelAtPeriodEnd:  data.CancelAtPeriodEnd,
				CanceledAt:         grpcUtils.TimeToTimestamp(data.CanceledAt),
				ExpiredAt:          grpcUtils.TimeToTimestamp(data.ExpiredAt),
				CreatedAt:          grpcUtils.TimeToTimestamp(&data.CreatedAt),
				UpdatedAt:          grpcUtils.TimeToTimestamp(&data.UpdatedAt),
			},
		},
	}

	jsonData, err := jsonUtils.MarshalToJsonString(eventData)
	if err != nil {
		return err
	}
	e.Data = json.RawMessage(jsonData)
	return nil
}

func (e *subscriptionCreatedEvent) ParseData() (*eventspb.EventEnvelope_SubscriptionCreated, error) {
	data := new(eventspb.EventEnvelope_SubscriptionCreated)
	if err := jsonUtils.UnmarshalJson(string(e.Data), data); err != nil {
		return nil, err
	}

	return data, nil
}
