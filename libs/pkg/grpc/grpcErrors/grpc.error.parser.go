package grpcerrors

import (
	"regexp"
	"strings"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/constants"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// mapErrorMessageToStatusCode maps error messages to gRPC status codes based on content patterns
func mapErrorMessageToStatusCode(errorMessage string) codes.Code {
	lowerMsg := strings.ToLower(errorMessage)

	if strings.Contains(lowerMsg, "already exists") ||
		strings.Contains(lowerMsg, "duplicate") ||
		strings.Contains(lowerMsg, "unique constraint") ||
		strings.Contains(lowerMsg, "uq_") {
		return codes.AlreadyExists
	}

	if strings.Contains(lowerMsg, "not found") ||
		strings.Contains(lowerMsg, "does not exist") {
		return codes.NotFound
	}

	if strings.Contains(lowerMsg, "validation") ||
		strings.Contains(lowerMsg, "invalid") ||
		strings.Contains(lowerMsg, "required") ||
		strings.Contains(lowerMsg, "must be") ||
		strings.Contains(lowerMsg, "should not be empty") {
		return codes.InvalidArgument
	}

	if strings.Contains(lowerMsg, "timeout") ||
		strings.Contains(lowerMsg, "connection") ||
		strings.Contains(lowerMsg, "etimedout") {
		return codes.DeadlineExceeded
	}

	if strings.Contains(lowerMsg, "unauthorized") ||
		strings.Contains(lowerMsg, "access denied") {
		return codes.Unauthenticated
	}

	if strings.Contains(lowerMsg, "forbidden") {
		return codes.PermissionDenied
	}

	if strings.Contains(lowerMsg, "insufficient") ||
		strings.Contains(lowerMsg, "not enough") {
		return codes.InvalidArgument
	}

	return codes.Internal
}

// cleanupErrorMessage cleans up error messages for better readability
func cleanupErrorMessage(message string) string {
	// Handle TypeORM constraint errors
	if strings.Contains(message, "unique constraint") {
		lowerMsg := strings.ToLower(message)
		if strings.Contains(lowerMsg, "sku") ||
			strings.Contains(message, "UQ_5ec10f972b1fa4f1e60d66d28bc") {
			return "Record with this SKU already exists"
		}
		if strings.Contains(lowerMsg, "email") {
			return "User with this email already exists"
		}
		if strings.Contains(lowerMsg, "username") {
			return "User with this username already exists"
		}
		if strings.Contains(lowerMsg, "name") {
			return "Record with this name already exists"
		}
		if strings.Contains(lowerMsg, "phone") {
			return "User with this phone number already exists"
		}
		if strings.Contains(lowerMsg, "order_number") {
			return "Order with this number already exists"
		}
		if strings.Contains(lowerMsg, "transaction_id") {
			return "Transaction with this ID already exists"
		}
		return "Duplicate entry found"
	}

	// Handle common database errors
	if strings.Contains(message, "QueryFailedError") {
		return "Database operation failed"
	}

	if strings.Contains(message, "duplicate key value violates") {
		return "Duplicate entry found"
	}

	// Handle validation errors
	if strings.Contains(strings.ToLower(message), "should not be empty") {
		re := regexp.MustCompile(`\b\w+\b should not be empty`)
		message = re.ReplaceAllStringFunc(message, func(match string) string {
			parts := strings.Split(match, " ")
			if len(parts) > 0 {
				return parts[0] + " is required"
			}
			return match
		})
	}

	// Handle insufficient stock/points errors
	lowerMsg := strings.ToLower(message)
	if strings.Contains(lowerMsg, "insufficient stock") {
		return "Insufficient stock available"
	}

	if strings.Contains(lowerMsg, "insufficient points") {
		return "Insufficient points available"
	}

	if strings.Contains(lowerMsg, "not enough stock") {
		return "Not enough stock available"
	}

	if strings.Contains(lowerMsg, "not enough points") {
		return "Not enough points available"
	}

	// Remove technical prefixes
	re := regexp.MustCompile(`(?i)^(Error: |QueryFailedError: |ValidationError: )`)
	message = re.ReplaceAllString(message, "")

	// Remove IP addresses
	re = regexp.MustCompile(`(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5]):([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])`)
	message = re.ReplaceAllString(message, "")

	// Remove URLs
	re = regexp.MustCompile(`https?://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&/=]*)`)
	message = re.ReplaceAllString(message, "")

	return strings.TrimSpace(message)
}

// ParseError parses an error and returns a GrpcErr with appropriate status code and cleaned message
func ParseError(err error) GrpcErr {
	if err == nil {
		return nil
	}

	errorMessage := err.Error()
	cleanedMessage := cleanupErrorMessage(errorMessage)

	// First check if it's already a gRPC error
	if st, ok := status.FromError(err); ok {
		return NewGrpcError(st.Code(), constants.ErrInternalServerErrorTitle, cleanedMessage, "")
	}

	// Try to map by error message
	statusCode := mapErrorMessageToStatusCode(errorMessage)

	// Get the title based on status code
	var title string
	switch statusCode {
	case codes.AlreadyExists:
		title = constants.ErrConflictTitle
	case codes.NotFound:
		title = constants.ErrNotFoundTitle
	case codes.InvalidArgument:
		title = constants.ErrBadRequestTitle
	case codes.Unauthenticated:
		title = constants.ErrUnauthorizedTitle
	case codes.PermissionDenied:
		title = constants.ErrForbiddenTitle
	case codes.DeadlineExceeded:
		title = "Request Timeout"
	default:
		title = constants.ErrInternalServerErrorTitle
	}

	return NewGrpcError(statusCode, title, cleanedMessage, "")
}
