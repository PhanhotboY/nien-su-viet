package response

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

// StandardRes response format
type APIResponse[T any] struct {
	Code       int         `json:"statusCode" example:"200" doc:"HTTP status code"`                        // HTTP status code
	Message    string      `json:"message" example:"success" doc:"Response message"`                       // Message
	Data       T           `json:"data" doc:"Response data"`                                               // data return null not show
	Pagination *Pagination `json:"pagination,omitempty" doc:"Pagination info,omitempty"`                   // Pagination info
	Timestamp  int64       `json:"timestamp" example:"1625247600" doc:"Response timestamp in Unix format"` // Response timestamp
}

type Pagination struct {
	Total      int `json:"total" example:"100" doc:"Total number of items"`     // Total number of items
	Limit      int `json:"limit" example:"10" doc:"Number of items per page"`   // Number of items per page
	Page       int `json:"page" example:"1" doc:"Current page number"`          // Current page number
	TotalPages int `json:"totalPages" example:"10" doc:"Total number of pages"` // Total number of pages
}

type APIBodyResponse[T any] struct {
	Status int `json:"-" doc:"HTTP status code"` // HTTP status code
	Body   APIResponse[T]
}

type OperationResult struct {
	Id      string `json:"id" doc:"Id of the document affected by the action"`
	Success bool   `json:"success"` // Indicates if the operation was successful
}

func PaginatedSuccessResponse[T any](code int, data T, pagination *Pagination) *APIBodyResponse[T] {
	return &APIBodyResponse[T]{
		Status: code,
		Body: APIResponse[T]{
			Code:       code,
			Message:    "success",
			Data:       data,
			Pagination: pagination,
		},
	}
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

func OperationSuccessResponse(code int, id string) *APIBodyResponse[OperationResult] {
	return SuccessResponse(code, OperationResult{Id: id, Success: true})
}

func NewError(code int, message string) huma.StatusError {
	return huma.NewError(code, message)
}

type HandlerFunc[T any, R any] func(context.Context, *T) (*APIBodyResponse[R], error)

func Wrap[T any, R any](handler HandlerFunc[T, R]) func(context.Context, *T) (*APIBodyResponse[R], error) {
	return func(ctx context.Context, input *T) (*APIBodyResponse[R], error) {
		res, err := handler(ctx, input)
		if err != nil {
			if _, ok := err.(huma.StatusError); ok {
				return nil, err
			}
			fmt.Printf("error handling request: %s", err.Error())
			if strings.Contains(err.Error(), "record not found") {
				return nil, huma.NewError(404, "Record not found")
			}
			return nil, huma.NewError(500, "Internal server error: "+err.Error())
		}

		res.Body.Timestamp = time.Now().Unix()
		return res, nil
	}
}
