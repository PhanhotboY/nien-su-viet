package types

import "context"

type GrpcHandler[T any, R any] interface {
	Handle(ctx context.Context, req T) (R, error)
}
