package response

import (
	"context"
	"time"
)

// StandardRes response format
type APIResponse[T any] struct {
	Code      int    `json:"statusCode" example:"200" doc:"HTTP status code"`                        // HTTP status code
	Message   string `json:"message" example:"success" doc:"Response message"`                       // Message
	Data      T      `json:"data" doc:"Response data"`                                               // data return null not show
	Timestamp int64  `json:"timestamp" example:"1625247600" doc:"Response timestamp in Unix format"` // Response timestamp
}

type APIBodyResponse[T any] struct {
	Status int `json:"-" doc:"HTTP status code"` // HTTP status code
	Body   APIResponse[T]
}

type OperationResult struct {
	Success bool `json:"success"` // Indicates if the operation was successful
}

func SuccessResponse[T any](code int, data T) *APIBodyResponse[T] {
	return &APIBodyResponse[T]{
		Status: code,
		Body: APIResponse[T]{
			Code:    code,
			Message: "success",
			Data:    data,
		},
	}
}

func OperationSuccessResponse(code int) *APIBodyResponse[OperationResult] {
	return SuccessResponse(code, OperationResult{Success: true})
}

func ErrorResponse[T any](code int, message string, err interface{}) *APIBodyResponse[T] {
	return &APIBodyResponse[T]{
		Status: code,
		Body: APIResponse[T]{
			Code:    code,
			Message: message,
		},
	}
}

type HandlerFunc[T any, R any] func(context.Context, *T) (*APIBodyResponse[R], error)

func Wrap[T any, R any](handler HandlerFunc[T, R]) func(context.Context, *T) (*APIBodyResponse[R], error) {
	return func(ctx context.Context, input *T) (*APIBodyResponse[R], error) {
		res, err := handler(ctx, input)
		if err != nil {
			return nil, err
		}

		res.Body.Timestamp = time.Now().Unix()
		return res, nil
	}
}
