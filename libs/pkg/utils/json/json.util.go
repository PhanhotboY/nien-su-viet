package jsonUtils

import (
	"encoding/json"
	"fmt"
)

func MarshalToJsonString(input any) (string, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func UnmarshalJson[T any](input string, output *T) error {
	if err := json.Unmarshal([]byte(input), output); err != nil {
		return fmt.Errorf("error unmarshal data for type '%T': %v", output, err)
	}

	return nil
}
