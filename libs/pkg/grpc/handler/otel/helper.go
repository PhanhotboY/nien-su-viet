package otel

import "google.golang.org/grpc/codes"

func mapGrpcCodeToHttpStatusCode(code codes.Code) int {
	switch code {
	case codes.OK: // OK
		return 200

	case codes.FailedPrecondition: // FAILED_PRECONDITION
		fallthrough
	case codes.InvalidArgument: // INVALID_ARGUMENT
		fallthrough
	case codes.OutOfRange: // OUT_OF_RANGE
		return 400

	case codes.Unauthenticated: // UNAUTHENTICATED
		return 401
	case codes.PermissionDenied: // PERMISSION_DENIED
		return 403
	case codes.NotFound: // NOT_FOUND
		return 404

	case codes.DeadlineExceeded: // DEADLINE_EXCEEDED
		fallthrough
	case codes.Canceled: // CANCELLED
		return 408

	case codes.AlreadyExists: // ALREADY_EXISTS
		fallthrough
	case codes.Aborted: // ABORTED
		return 409
	case codes.ResourceExhausted: // RESOURCE_EXHAUSTED
		return 429

	case codes.Unimplemented: // UNIMPLEMENTED
		return 501
	case codes.Unavailable: // UNAVAILABLE
		return 503
	case codes.Internal: // INTERNAL
		fallthrough
	case codes.DataLoss: // DATA_LOSS
		fallthrough
	case codes.Unknown: // UNKNOWN
		fallthrough
	default:
		return 500 // Default to 500 for unknown codes
	}
}
