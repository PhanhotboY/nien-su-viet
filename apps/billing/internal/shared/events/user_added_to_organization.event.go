package event

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/utils"
	eventspb "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/events"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
	jsonUtils "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/json"
)

type UserAddedToOrganizationEventData struct {
	UserId         string `json:"user_id"`
	OrganizationId string `json:"organization_id"`
	IsPremium      bool   `json:"is_premium"`
}

type UserAddedToOrganizationEvent interface {
	MessageEvent[UserAddedToOrganizationEventData, *eventspb.UserAddedToOrganization]
}

type userAddedToOrganizationEvent struct {
	*types.Message
}

func NewUserAddedToOrganizationEvent(msg types.IMessage) UserAddedToOrganizationEvent {
	if msg == nil {
		msg = types.NewMessage(uuid.NewString(), []byte("{}"))
	}
	if m, ok := msg.(*userAddedToOrganizationEvent); ok {
		msg = m.Message
	}

	msg.SetPattern(utils.GetMessageName(userAddedToOrganizationEvent{}))
	return &userAddedToOrganizationEvent{
		Message: msg.(*types.Message),
	}
}

func (e *userAddedToOrganizationEvent) SetRawData(data string) error {
	if err := dtoUtil.ValidateStruct(data, &eventspb.UserAddedToOrganization{}); err != nil {
		return err
	}
	e.Data = json.RawMessage(data)
	return nil
}

func (e *userAddedToOrganizationEvent) SetData(data UserAddedToOrganizationEventData) error {
	eventData := &eventspb.UserAddedToOrganization{
		UserId:         data.UserId,
		OrganizationId: data.OrganizationId,
		IsPremium:      data.IsPremium,
	}

	jsonData, err := jsonUtils.MarshalToJsonString(eventData)
	fmt.Println("UserAddedToOrganizationEvent SetData jsonData:", jsonData)
	if err != nil {
		return err
	}
	e.Data = json.RawMessage(jsonData)
	return nil
}

func (e *userAddedToOrganizationEvent) ParseData() (*eventspb.UserAddedToOrganization, error) {
	var eventData eventspb.UserAddedToOrganization
	if err := jsonUtils.UnmarshalJson(string(e.Data), &eventData); err != nil {
		return nil, err
	}
	return &eventData, nil
}
