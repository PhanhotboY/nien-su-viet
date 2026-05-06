import { Inject, Injectable, Logger } from '@nestjs/common';
import { tap } from 'rxjs';
import { RpcException, ClientRMQ } from '@nestjs/microservices';
import { trace } from '@opentelemetry/api';
import { isUUID } from 'class-validator';
import { status } from '@grpc/grpc-js';

import {
  getExcerpt,
  WithSpan,
  recordOperationTiming,
  commonUtils,
} from '@phanhotboy/nsv-common';
import type {
  CreateHistoricalEventRequest,
  DeleteHistoricalEventRequest,
  UpdateHistoricalEventRequest,
} from '@phanhotboy/genproto/historical_event_service/historical_events';
import { IHistoricalEventDbRepo } from '@historical-event/modules/historical-event/domain/repo/historical-event.db.repo';
import { RMQ } from '@phanhotboy/constants';
import { HistoricalEventCreatedEvent } from './application/events/historical-event-created.event';
import { HistoricalEventUpdatedEvent } from './application/events/historical-event-updated.event';
import { HistoricalEventDeletedEvent } from './application/events/historical-event-deleted.event';
import { toEventDateType } from '@historical-event/helper/dateType.helper';

const tracerName = 'historical-event-service';
@Injectable()
export class HistoricalEventService {
  constructor(
    private readonly logger: Logger,
    @Inject(IHistoricalEventDbRepo)
    private readonly eventPgRepo: IHistoricalEventDbRepo,
    @Inject(RMQ.TOPIC_EVENTS_EXCHANGE) private readonly rmqClient: ClientRMQ,
  ) {}

  @WithSpan(tracerName, 'historical_event.create', {
    'operation.type': 'create',
  })
  async createEvent(payload: CreateHistoricalEventRequest) {
    const event = await recordOperationTiming(tracerName, 'prisma.create', () =>
      this.eventPgRepo.createHistoricalEvent(payload),
    );

    const createdEvent = new HistoricalEventCreatedEvent({
      ...event,
      createdAt: event.createdAt.toISOString(),
      updatedAt: event.updatedAt.toISOString(),
    });
    const routingKey = commonUtils.getRoutingKey(HistoricalEventCreatedEvent);
    this.rmqClient
      .emit(routingKey, createdEvent)
      .pipe(
        tap(() => {
          this.logger.log(`Emitted ${routingKey} for event ID: ${event.id}`);
        }),
      )
      .subscribe({
        error: (err) => {
          this.logger.error(
            `Failed to emit ${routingKey} for event ID: ${event.id}`,
            err,
          );
        },
      });

    return { id: event.id, success: true };
  }

  @WithSpan(tracerName, 'historical_event.get_by_id', {
    'operation.type': 'read',
  })
  private async getEventById({ id }: { id: string }) {
    const span = trace.getActiveSpan();

    if (isUUID(id, '4') === false) {
      const exception = new RpcException({
        message: 'ID không hợp lệ',
        code: status.INVALID_ARGUMENT,
      });
      span?.recordException(exception);
      throw exception;
    }

    return await recordOperationTiming(tracerName, 'prisma.findUnique', () =>
      this.eventPgRepo.getHistoricalEventById({ where: { id } }),
    );
  }

  @WithSpan(tracerName, 'historical_event.get_by_id_and_author', {
    'operation.type': 'read',
  })
  async getAuthorEventById(id: string, authorId: string) {
    const event = recordOperationTiming(
      tracerName,
      'prisma.findUnique.with.author',
      () =>
        this.eventPgRepo.getHistoricalEventById({
          where: {
            id,
            authorId,
          },
        }),
    ) as any;
    return event ?? undefined;
  }

  @WithSpan(tracerName, 'historical_event.update', {
    'operation.type': 'update',
  })
  async updateEvent({ id, ...payload }: UpdateHistoricalEventRequest) {
    const found = await this.getEventById({ id });
    if (!found) {
      const exception = new RpcException({
        message: 'Sự kiện lịch sử không tồn tại',
        code: status.NOT_FOUND,
      });
      trace.getActiveSpan()?.recordException(exception);
      throw exception;
    }

    await recordOperationTiming(tracerName, 'prisma.update', () =>
      this.eventPgRepo.updateHistoricalEvent(id, payload),
    );

    const updatedEvent = new HistoricalEventUpdatedEvent({
      ...payload,
      fromDateType: toEventDateType(payload.fromDateType),
      toDateType: toEventDateType(payload.toDateType),
      id,
    });
    const routingKey = commonUtils.getRoutingKey(HistoricalEventUpdatedEvent);
    this.rmqClient
      .emit(routingKey, updatedEvent)
      .pipe(
        tap(() => {
          this.logger.log(`Emitted ${routingKey} for event ID: ${id}`);
        }),
      )
      .subscribe({
        error: (err) => {
          this.logger.error(
            `Failed to emit ${routingKey} for event ID: ${id}`,
            err,
          );
        },
      });

    return { id, success: true };
  }

  @WithSpan(tracerName, 'historical_event.delete', {
    'operation.type': 'delete',
  })
  async deleteEvent({ id }: DeleteHistoricalEventRequest) {
    const span = trace.getActiveSpan();
    const event = await this.getEventById({ id });
    if (!event) {
      const exception = new RpcException({
        message: 'Sự kiện lịch sử không tồn tại',
        code: status.NOT_FOUND,
      });
      span?.recordException(exception);
      throw exception;
    }

    await recordOperationTiming(tracerName, 'prisma.delete', () =>
      this.eventPgRepo.deleteHistoricalEvent(id),
    );

    const deletedEvent = new HistoricalEventDeletedEvent({ id });
    const routingKey = commonUtils.getRoutingKey(HistoricalEventDeletedEvent);
    this.rmqClient
      .emit(routingKey, deletedEvent)
      .pipe(
        tap(() => {
          this.logger.log(`Emitted ${routingKey} for event ID: ${id}`);
        }),
      )
      .subscribe({
        error: (err) => {
          this.logger.error(
            `Failed to emit ${routingKey} for event ID: ${id}`,
            err,
          );
        },
      });

    return { id, success: true };
  }
}
