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
		return nil
	}
	return res
}

func ValidateConsumeContextData[T any](ctx types.MessageConsumeContext, dto T, l logger.Logger) *T {
	return ValidateDto(ctx.Message().GetData(), dto, l)
}

func ValidateStruct[T any](input any, dto *T) error {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error marshal input data for DTO '%T': %v", dto, err)
	}

	if err := json.Unmarshal(inputBytes, dto); err != nil {
		return fmt.Errorf("error unmarshal data for DTO '%T': %v", dto, err)
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			return ParseValidationErrors(validateErrs)
		}
		return fmt.Errorf("error validating DTO '%T': %v", dto, err)
	}
	return nil
}
