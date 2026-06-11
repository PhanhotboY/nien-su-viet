package dtoUtil

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ParseValidationErrors(errors validator.ValidationErrors) error {
	if len(errors) == 0 {
		return nil
	}
	err := errors[0]
	validation := err.ActualTag()
	if validation == "required" {
		validation = "not empty"
	}
	return fmt.Errorf("validation failed: %s must be %s", err.Namespace(), validation)
}
