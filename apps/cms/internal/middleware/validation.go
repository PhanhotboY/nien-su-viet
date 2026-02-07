package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type contextKey string

const validatorKey contextKey = "validator"

func ValidatorMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			validate := validator.New()

			// set the middleware in context
			ctx := context.WithValue(r.Context(), validatorKey, validate)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// GetValidator retrieves the validator from the request context
func GetValidator(r *http.Request) *validator.Validate {
	fmt.Printf("Retrieving validator from context: %v\n", r.Context().Value(validatorKey))
	if val := r.Context().Value(validatorKey); val != nil {
		if v, ok := val.(*validator.Validate); ok {
			return v
		}
	}
	return nil
}
