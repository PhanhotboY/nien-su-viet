package prepo // persistence repository

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/repository"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

type inBoxEventDbRepo struct {
	db dbcontracts.TxContextDb
}

func NewInBoxEventDbRepo(db dbcontracts.TxContextDb) drepo.InBoxEventDbRepo {
	return &inBoxEventDbRepo{db: db}
}

func (r *inBoxEventDbRepo) Insert(ctx context.Context, event *entity.InboxEvent) error {
	return r.db.WithTxIfExists(ctx).DB().Create(event).Error
}

func (r *inBoxEventDbRepo) UpdateStatus(ctx context.Context, id string, status entity.InboxEventStatus) error {
	return r.db.WithTxIfExists(ctx).DB().Model(&entity.InboxEvent{}).Where("id = ?", id).Update("status", status).Error
}

func (r *inBoxEventDbRepo) FindByID(ctx context.Context, id string) (*entity.InboxEvent, error) {
	var event entity.InboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().First(&event, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *inBoxEventDbRepo) FindPendingEvents(ctx context.Context, limit int) ([]*entity.InboxEvent, error) {
	var events []*entity.InboxEvent
	if err := r.db.WithTxIfExists(ctx).DB().Where("status = ?", entity.INBOX_EVENT_STATUS_PENDING).Limit(limit).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
