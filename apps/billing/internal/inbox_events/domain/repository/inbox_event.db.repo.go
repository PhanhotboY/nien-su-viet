package drepo // domain repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"
)

type InBoxEventDbRepo interface {
	Insert(ctx context.Context, event *entity.InboxEvent) error
	UpdateStatus(ctx context.Context, id string, status entity.InboxEventStatus) error
	FindByID(ctx context.Context, id string) (*entity.InboxEvent, error)

	FetchPending(ctx context.Context, limit int) ([]*entity.InboxEvent, error)
	MarkPublished(ctx context.Context, id string) error
	MarkFailed(ctx context.Context, id string, err error) error
}
