package grpcUtils

import (
	"encoding/json"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func UnmarshalProtoMessage[T proto.Message](input any, message T, logger logger.Logger) (T, error) {
	res := new(T) // nil pointer of type T, to return in case of error

	inputBytes, err := json.Marshal(input)
	if err != nil {
		logger.Errorf("error marshal input data for DTO '%T': %v", message, err)
		return *res, err
	}

	unmarshalOptions := protojson.UnmarshalOptions{
		AllowPartial:   true, // Allow partial messages
		DiscardUnknown: true, // Ignore unknown fields in the input JSON
	}
	if err := unmarshalOptions.Unmarshal(inputBytes, message); err != nil {
		logger.Errorf("error unmarshal data for DTO '%T': %v", message, err)
		return *res, err
	}

	return message, nil
}
