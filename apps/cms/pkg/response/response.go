package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StandardRes response format
type APIResponse[T any] struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    T           `json:"data,omitempty"`  // data return null not show
	Error   interface{} `json:"error,omitempty"` // Error return null not show
}

func SuccessResponse[T any](c *gin.Context, data T) {
	c.JSON(200, APIResponse[T]{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, APIResponse[any]{
		Code:    code,
		Message: message,
		Error:   err,
	})
}

type HandlerFunc[T any] func(ctx *gin.Context) (res T, err error)

func Wrap[T any](handler HandlerFunc[T]) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		res, err := handler(ctx)
		if err != nil {
			if apiErr, ok := err.(*APIError); ok {
				ErrorResponse(ctx, apiErr.StatusCode, apiErr.Message, apiErr.Err)
			} else {
				ErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error", err)
			}
			return
		}
		SuccessResponse(ctx, res)
	}
}
