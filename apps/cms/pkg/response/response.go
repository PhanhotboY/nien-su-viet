package response

import (
	"context"
	"net/http"
)

// StandardRes response format
type APIResponse[T any] struct {
	Code    int         `json:"code" example:"200" doc:"HTTP status code"`        // HTTP status code
	Message string      `json:"message" example:"success" doc:"Response message"` // Message
	Data    T           `json:"data" doc:"Response data"`                         // data return null not show
	Error   interface{} `json:"error,omitempty" doc:"Response error"`             // Error return null not show
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
			Error:   err,
		},
	}
}

type HandlerFunc[T any, R any] func(context.Context, *T) (*APIBodyResponse[R], error)

func Wrap[T any, R any](handler HandlerFunc[T, R]) func(context.Context, *T) (*APIBodyResponse[R], error) {
	return func(ctx context.Context, input *T) (*APIBodyResponse[R], error) {
		res, err := handler(ctx, input)
		if err != nil {
			if apiErr, ok := err.(*APIError); ok {
				return ErrorResponse[R](apiErr.Code, apiErr.Message, apiErr.Err), nil
			} else {
				return ErrorResponse[R](http.StatusInternalServerError, "Internal Server Error", err), nil
			}
		}

		return res, nil
	}
}
