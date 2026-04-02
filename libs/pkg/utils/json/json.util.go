package jsonUtils

import (
	"encoding/json"
)

func MarshalToJsonString(input any) (string, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
