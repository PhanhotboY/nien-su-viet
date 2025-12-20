package request

type APIBodyRequest[T any] struct {
	Body T
}
