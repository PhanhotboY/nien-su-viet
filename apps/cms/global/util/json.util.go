package util

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// jsonbValue is a generic helper for JSONB Value method
func JsonbValue(v interface{}) (driver.Value, error) {
	if v == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

// jsonbScan is a generic helper for JSONB Scan method
func JsonbScan[T any](dest *T, value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, dest)
}
