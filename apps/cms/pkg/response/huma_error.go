package response

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
)

// StandardRes response format
type errorResponse struct {
	Code      int    `json:"statusCode" example:"500" doc:"HTTP status code"`                        // HTTP status code
	Message   string `json:"message" example:"Internal Server Error" doc:"Error message"`            // Message
	Timestamp int64  `json:"timestamp" example:"1625247600" doc:"Response timestamp in Unix format"` // Response timestamp
}

func (e *errorResponse) Error() string {
	return e.Message
}

func (e *errorResponse) GetStatus() int {
	return e.Code
}
func SerializeError(status int, message string, errs ...error) huma.StatusError {
	var details string
	for i, err := range errs {
		if i == 0 {
			details = details + ": "
		}
		details = details + err.Error()
	}
	return &errorResponse{
		Code:      status,
		Message:   message + details,
		Timestamp: time.Now().Unix(),
	}
}
