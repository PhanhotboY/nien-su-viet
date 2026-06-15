package worker

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/entity"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/inbox_events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/bus"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	cworker "github.com/phanhotboy/nien-su-viet/libs/pkg/core/worker"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
)

type inboxEventsWorker struct {
	logger logger.Logger
	repo   drepo.InBoxEventDbRepo
	bus    bus.Bus
	db     dbcontracts.TxContextDb
}

func NewInboxEventsWorker(
	l logger.Logger,
	repo drepo.InBoxEventDbRepo,
	bus bus.Bus,
	db dbcontracts.TxContextDb,
) cworker.Worker {
	worker := inboxEventsWorker{l, repo, bus, db}

	return cworker.NewBackgroundWorker(worker.Start, worker.Stop)
}

func (w inboxEventsWorker) Start(ctx context.Context) error {
	w.logger.Debug("Starting Inbox events worker...")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-ticker.C:
			w.logger.Debug("Publishing Inbox events")
			processCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			w.processBatch(processCtx)
		}
	}
}

func (w inboxEventsWorker) Stop(ctx context.Context) error {
	w.logger.Debug("Stopping Inbox Events worker...")
	return nil
}

func (w inboxEventsWorker) processBatch(ctx context.Context) {
	var wg sync.WaitGroup

	w.db.RunInTx(ctx, func(ctx context.Context, tx dbcontracts.TxContextDb) error {
		events, err := w.repo.FetchPending(ctx, 100)
		if err != nil {
			w.logger.Error("fetch error:", err)
			return err
		}

		for _, event := range events {
			wg.Add(1)

			wg.Go(func() {
				defer wg.Done()

				if err := w.handleEvent(ctx, event); err != nil {
					w.logger.Error("publish failed:", err)
				}
			})
		}

		return nil
	})

	wg.Wait()
}

func (w inboxEventsWorker) handleEvent(
	ctx context.Context,
	event *entity.InboxEvent,
) error {
	eventMsg := &types.Message{
		MessageId: event.ID.String(),
		Created:   event.CreatedAt,
		Data:      json.RawMessage(event.Payload),
	}
	err := w.bus.PublishMessage(
		ctx,
		eventMsg,
	)

	if err != nil {
		return w.repo.MarkFailed(ctx, event.ID.String(), err)
	}

	return w.repo.MarkPublished(ctx, event.ID.String())
}
