import { Inject, Injectable, Logger } from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';
import { trace } from '@opentelemetry/api';
import { isUUID } from 'class-validator';

import {
  getExcerpt,
  RedisService,
  UtilService,
  WithSpan,
  recordOperationTiming,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import type {
  CreateHistoricalEventRequest,
  DeleteHistoricalEventRequest,
  GetAllHistoricalEventsRequest,
  GetHistoricalEventPreviewRequest,
  GetHistoricalEventRequest,
  UpdateHistoricalEventRequest,
} from '@phanhotboy/genproto/historical_event_service/historical_events';
import { UserService } from '../user';
import { HistoricalEventsEsRepo } from '@historical-event/modules/historical-event/infrastructure/persistence/elasticsearch/historical-events.repo';
import HistoricalEventsPgRepo from '@historical-event/modules/historical-event/infrastructure/persistence/postgresql/historical-events.repo';
import { requestToListQuery, uniqueQueryBuilder } from './helper';
import { status } from '@grpc/grpc-js';

const tracerName = 'historical-event-service';
@Injectable()
export class HistoricalEventService {
  private readonly serviceName = 'historical-events';
  private readonly cacheKey: string;

  constructor(
    private readonly util: UtilService,
    @Inject(RedisService)
    private readonly redisService: RedisServiceType,
    private readonly userService: UserService,
    private readonly logger: Logger,
    private readonly eventPgRepo: HistoricalEventsPgRepo,
    private readonly eventEsRepo: HistoricalEventsEsRepo,
  ) {
    this.cacheKey = this.util.genCacheKey(this.serviceName);
  }

  @WithSpan(tracerName, 'historical_event.create', {
    'operation.type': 'create',
  })
  async createEvent(payload: CreateHistoricalEventRequest) {
    await this.userService.findUserById(payload.authorId);

    const event = await recordOperationTiming(tracerName, 'prisma.create', () =>
      this.eventPgRepo.createHistoricalEvent(payload),
    );

    // Clear cache + Index to Elasticsearch
    const [cacheRes, indexRes] = await Promise.allSettled([
      recordOperationTiming(tracerName, 'redis.cache.delete', () =>
        this.redisService.del(this.cacheKey),
      ),
      recordOperationTiming(tracerName, 'elasticsearch.index', () =>
        this.eventEsRepo.index(event),
      ),
    ]);
    if (cacheRes.status === 'rejected') {
      this.logger.error(
        `Failed to clear cache for key ${this.cacheKey} after creating event with id ${event.id}`,
        cacheRes.reason,
      );
    }
    if (indexRes.status === 'rejected') {
      this.logger.error(
        `Failed to index event with id ${event.id} to Elasticsearch after creation`,
        indexRes.reason,
      );
    }

    return { id: event.id, success: true };
  }

  @WithSpan(tracerName, 'historical_event.list', { 'operation.type': 'list' })
  async getEvents(query: GetAllHistoricalEventsRequest) {
    const page = query.page || 1;
    const limit = query.limit || 10;
    const options = requestToListQuery({ ...query, page, limit });

    return this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
      },
      async () => {
        const [events, total] = await Promise.all([
          recordOperationTiming(tracerName, 'prisma.findMany', () =>
            this.eventPgRepo.searchHistoricalEvents(options),
          ),
          recordOperationTiming(tracerName, 'prisma.count', () =>
            this.eventPgRepo.countHistoricalEvents(options),
          ),
        ]);

        return {
          events,
          pagination: {
            total,
            page,
            limit,
            totalPages: Math.ceil(total / limit),
          },
        };
      },
    );
  }

  @WithSpan(tracerName, 'historical_event.get_by_id', {
    'operation.type': 'read',
  })
  async getEventById({ id }: GetHistoricalEventRequest) {
    const span = trace.getActiveSpan();

    if (isUUID(id, '4') === false) {
      const exception = new RpcException({
        message: 'ID không hợp lệ',
        code: status.INVALID_ARGUMENT,
      });
      span?.recordException(exception);
      throw exception;
    }

    const options = uniqueQueryBuilder(id);

    return await this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
        notFoundMessage: 'Sự kiện lịch sử không tồn tại',
      },
      async () => {
        const event = await recordOperationTiming(
          tracerName,
          'prisma.findUnique',
          () => this.eventPgRepo.getHistoricalEventById(options),
        );
        return (
          event && {
            ...event,
            author: (event as any).author,
            categories: (event as any).categories,
          }
        );
      },
    );
  }

  @WithSpan(tracerName, 'historical_event.get_preview', {
    'operation.type': 'read',
  })
  async getEventPreviewById({ id }: GetHistoricalEventPreviewRequest) {
    const event = await this.getEventById({ id });

    const excerpt = getExcerpt(event!.content, 1000);

    return {
      ...event!,
      excerpt,
    };
  }

  @WithSpan(tracerName, 'historical_event.get_by_id_and_author', {
    'operation.type': 'read',
  })
  async getAuthorEventById(id: string, authorId: string) {
    const options = uniqueQueryBuilder(id, authorId);

    return this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
        notFoundMessage: 'Sự kiện lịch sử không tồn tại',
      },
      async () => {
        const event = recordOperationTiming(
          tracerName,
          'prisma.findUnique.with.author',
          () => this.eventPgRepo.getHistoricalEventById(options),
        ) as any;
        return event ?? undefined;
      },
    );
  }

  @WithSpan(tracerName, 'historical_event.update', {
    'operation.type': 'update',
  })
  async updateEvent({ id, ...payload }: UpdateHistoricalEventRequest) {
    const found = await this.getEventById({ id });

    await this.eventPgRepo.updateHistoricalEvent(id, payload);

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return { id, success: true };
  }

  @WithSpan(tracerName, 'historical_event.delete', {
    'operation.type': 'delete',
  })
  async deleteEvent({ id, authorId }: DeleteHistoricalEventRequest) {
    const span = trace.getActiveSpan();
    const { data: event } = await this.getAuthorEventById(id, authorId);
    if (!event) {
      const exception = new RpcException({
        message: 'Sự kiện lịch sử không tồn tại',
        code: status.NOT_FOUND,
      });
      span?.recordException(exception);
      throw exception;
    }

    await this.eventPgRepo.deleteHistoricalEvent(id);

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return { id, success: true };
  }
}
