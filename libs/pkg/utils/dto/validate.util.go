package dtoUtil

import (
	"encoding/json"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
)

func ValidateDto[T any](data json.RawMessage, dto T, logger logger.Logger) *T {
	res := new(T)
	if err := json.Unmarshal(data, res); err != nil {
		logger.Errorf("error unmarshal data for DTO '%T': %v", dto, err)
	}
	return res
}

func ValidateConsumeContextData[T any](ctx types.MessageConsumeContext, dto T, l logger.Logger) *T {
	return ValidateDto(ctx.Message().GetData(), dto, l)
}

func ValidateStruct[T any](input any, dto T, logger logger.Logger) (*T, error) {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		logger.Errorf("error marshal input data for DTO '%T': %v", dto, err)
		return nil, err
	}

	res := new(T)
	if err := json.Unmarshal(inputBytes, res); err != nil {
		logger.Errorf("error unmarshal data for DTO '%T': %v", dto, err)
		return nil, err
	}

	return res, nil
}
