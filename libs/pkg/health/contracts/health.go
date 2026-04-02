package contracts

import "context"

type Health interface {
	CheckHealth(ctx context.Context) error
	GetHealthName() string
}
