package json

import (
	stdjson "encoding/json"
	"fmt"
	"reflect"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/serializer"
)

const contentTypeJSON = "application/json"

type DefaultJsonSerializer struct{}

func NewDefaultJsonSerializer() *DefaultJsonSerializer {
	return &DefaultJsonSerializer{}
}

func (d *DefaultJsonSerializer) Serialize(v interface{}) ([]byte, error) {
	return stdjson.Marshal(v)
}

func (d *DefaultJsonSerializer) Deserialize(data []byte, v interface{}) error {
	return stdjson.Unmarshal(data, v)
}

type defaultMessageJSONSerializer struct {
	jsonSerializer *DefaultJsonSerializer
}

func NewDefaultMessageJsonSerializer(jsonSerializer *DefaultJsonSerializer) serializer.MessageSerializer {
	if jsonSerializer == nil {
		jsonSerializer = NewDefaultJsonSerializer()
	}

	return &defaultMessageJSONSerializer{jsonSerializer: jsonSerializer}
}

func (d *defaultMessageJSONSerializer) ContentType() string {
	return contentTypeJSON
}

func (d *defaultMessageJSONSerializer) Serialize(message types.IMessage) (*serializer.SerializedMessage, error) {
	data, err := d.jsonSerializer.Serialize(message)
	if err != nil {
		return nil, err
	}

	return &serializer.SerializedMessage{Data: data, ContentType: contentTypeJSON}, nil
}

func (d *defaultMessageJSONSerializer) Deserialize(data []byte, contentType string) (types.IMessage, error) {
	if contentType != "" && contentType != contentTypeJSON {
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}

	var target interface{}
	target = &types.Message{}
	if err := d.jsonSerializer.Deserialize(data, target); err != nil {
		return nil, err
	}

	if msg, ok := target.(types.IMessage); ok {
		return msg, nil
	}

	if v := reflect.ValueOf(target); v.Kind() == reflect.Pointer && !v.IsNil() {
		if msg, ok := v.Elem().Interface().(types.IMessage); ok {
			return msg, nil
		}
	}

	return nil, fmt.Errorf("deserialized payload is not IMessage: %s")
}
