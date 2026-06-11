package zalopay

import (
	"encoding/json"
	"fmt"
	"time"
)

func ParseCallback(body []byte) (*CallbackData, error) {
	var payload CallbackPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, fmt.Errorf("zalopay: decode callback wrapper: %w", err)
	}
	if payload.Data == "" {
		return nil, fmt.Errorf("zalopay: missing callback data")
	}

	var data CallbackData
	if err := json.Unmarshal([]byte(payload.Data), &data); err != nil {
		return nil, fmt.Errorf("zalopay: decode callback data: %w", err)
	}
	return &data, nil
}

func GenerateAppTransID(prefix string, t time.Time) string {
	return fmt.Sprintf("%s%d", prefix, t.Unix())
}
