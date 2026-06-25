package event

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	eventspb "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/events"
	grpcUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/utils"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"

	subscriptionEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
	subscriptionhelper "github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/helper"
)

type SubscriptionCreatedEvent interface {
	MessageEvent[subscriptionEntity.Subscription, *eventspb.SubscriptionCreated]
}

type subscriptionCreatedEvent struct {
	*types.Message
}

func NewSubscriptionCreatedEvent(msg types.IMessage) SubscriptionCreatedEvent {
	if msg == nil {
		msg = types.NewMessage(uuid.NewString(), []byte("{}"))
	}
	if m, ok := msg.(*subscriptionCreatedEvent); ok {
		msg = m.Message
	}

	msg.SetPattern(utils.GetMessageName(subscriptionCreatedEvent{}))
	return &subscriptionCreatedEvent{
		Message: msg.(*types.Message),
	}
}

func (e *subscriptionCreatedEvent) SetRawData(data string) error {
	if err := dtoUtil.ValidateStruct(data, &eventspb.SubscriptionCreated{}); err != nil {
		return err
	}
	e.Data = json.RawMessage(data)
	return nil
}

func (e *subscriptionCreatedEvent) SetData(data subscriptionEntity.Subscription) error {
	eventData := &eventspb.SubscriptionCreated{
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
	}

	jsonData, err := jsonUtils.MarshalToJsonString(eventData)
	fmt.Println("SubscriptionCreatedEvent SetData jsonData:", jsonData)
	if err != nil {
		return err
	}
	e.Data = json.RawMessage(jsonData)
	return nil
}

func (e *subscriptionCreatedEvent) ParseData() (*eventspb.SubscriptionCreated, error) {
	data := new(eventspb.SubscriptionCreated)
	if err := jsonUtils.UnmarshalJson(string(e.Data), data); err != nil {
		return nil, err
	}

	return data, nil
}
