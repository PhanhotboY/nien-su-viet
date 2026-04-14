package dtoUtil

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
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

func ValidateStruct[T any](input any, dto T) (*T, error) {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error marshal input data for DTO '%T': %v", dto, err)
	}

	res := new(T)
	if err := json.Unmarshal(inputBytes, res); err != nil {
		return nil, fmt.Errorf("error unmarshal data for DTO '%T': %v", dto, err)
	}

	validate := validator.New()
	err = validate.Struct(res)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			return nil, fmt.Errorf("error validating input: %s", validateErrs[0].Error())
		}
		return nil, fmt.Errorf("error validating DTO '%T': %v", dto, err)
	}
	return res, nil
}
