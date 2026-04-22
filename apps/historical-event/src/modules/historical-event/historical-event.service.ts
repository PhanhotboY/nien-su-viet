import { Inject, Injectable, Logger } from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';
import { trace } from '@opentelemetry/api';
import { isUUID } from 'class-validator';

import { PrismaService } from '@historical-event/database';
import {
  getExcerpt,
  RedisService,
  UtilService,
  WithSpan,
  recordOperationTiming,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import { Prisma } from '@historical-event-prisma';
import {
  type CreateHistoricalEventRequest,
  CreateHistoricalEventResponse,
  DeleteHistoricalEventRequest,
  DeleteHistoricalEventResponse,
  type GetAllHistoricalEventsRequest,
  GetAllHistoricalEventsResponse,
  type GetHistoricalEventPreviewRequest,
  GetHistoricalEventPreviewResponse,
  type GetHistoricalEventRequest,
  GetHistoricalEventResponse,
  UpdateHistoricalEventRequest,
  UpdateHistoricalEventResponse,
} from '@phanhotboy/genproto/historical_event_service/historical_events';
import { UserService } from '../user';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import {
  toEventDateType,
  toGrpcEventDateType,
} from '@historical-event/helper/dateType.helper';

const tracerName = 'historical-event-service';
@Injectable()
export class HistoricalEventService {
  private readonly cachePrefix = 'historical-event';
  private readonly cacheKey: string;

  constructor(
    private readonly prisma: PrismaService,
    private readonly util: UtilService,
    @Inject(RedisService)
    private readonly redisService: RedisServiceType,
    private readonly userService: UserService,
    private readonly logger: Logger,
  ) {
    this.cacheKey = this.util.genCacheKey(this.cachePrefix);
  }

  @WithSpan(tracerName, 'historical_event.create', {
    'operation.type': 'create',
  })
  async createEvent(
    payload: CreateHistoricalEventRequest,
  ): Promise<CreateHistoricalEventResponse> {
    await this.userService.findUserById(payload.authorId);

    const event = await recordOperationTiming(tracerName, 'prisma.create', () =>
      this.prisma.historicalEvent.create({
        data: payload as any,
      }),
    );

    // Clear cache
    await recordOperationTiming(tracerName, 'redis.cache.delete', () =>
      this.redisService.del(this.cacheKey),
    );

    return { data: { id: event.id, success: true } };
  }

  @WithSpan(tracerName, 'historical_event.list', { 'operation.type': 'list' })
  async getEvents(
    query: GetAllHistoricalEventsRequest,
  ): Promise<GetAllHistoricalEventsResponse> {
    const {
      page = 1,
      limit = 10,
      categoryIds,
      fromDay,
      fromMonth,
      fromYear,
      toDay,
      toMonth,
      toYear,
      sortOrder = 'desc',
      sortBy = 'fromYear',
    } = query;

    const options = {
      where: {} as Prisma.HistoricalEventWhereInput,
      skip: (page - 1) * limit,
      take: limit,
      orderBy: {
        [sortBy]: sortOrder,
      },
      include: {
        author: true,
      },
    } satisfies Parameters<typeof this.prisma.historicalEvent.findMany>[0];

    if (categoryIds && categoryIds.length > 0) {
      options!.where!.categories = {
        some: { categoryId: { in: categoryIds } },
      };
    }

    const hasFromYear = fromYear !== undefined;
    const hasFromMonth = fromMonth !== undefined;
    const hasFromDay = fromDay !== undefined;
    const hasToYear = toYear !== undefined;
    const hasToMonth = toMonth !== undefined;
    const hasToDay = toDay !== undefined;

    if (hasFromYear) {
      options!.where!.fromYear = fromYear;
    }

    if (hasFromMonth && hasFromYear) {
      options!.where!.fromMonth = fromMonth;
    }

    if (hasFromDay && hasFromMonth && hasFromYear) {
      options!.where!.fromDay = fromDay;
    }
    if (hasToYear) {
      options!.where!.toYear = toYear;
    }

    if (hasToMonth && hasToYear) {
      options!.where!.toMonth = toMonth;
    }

    if (hasToDay && hasToMonth && hasToYear) {
      options!.where!.toDay = toDay;
    }

    return this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
      },
      async () => {
        const [events, total] = await Promise.all([
          recordOperationTiming(tracerName, 'prisma.findMany', () =>
            this.prisma.historicalEvent.findMany(options),
          ),
          recordOperationTiming(tracerName, 'prisma.count', () =>
            this.prisma.historicalEvent.count({ where: options.where }),
          ),
        ]);

        return {
          data: events as any,
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
  async getEventById({
    id,
  }: GetHistoricalEventRequest): Promise<GetHistoricalEventResponse> {
    const span = trace.getActiveSpan();

    if (isUUID(id, '4') === false) {
      const exception = new RpcException({
        message: 'ID không hợp lệ',
        statusCode: 400,
      });
      span?.recordException(exception);
      throw exception;
    }

    const options = {
      where: { id },
      include: {
        author: true,
        categories: { include: { category: true, event: false } },
      },
    } satisfies Parameters<typeof this.prisma.historicalEvent.findUnique>[0];

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
          () => this.prisma.historicalEvent.findUnique(options),
        );
        return {
          data: event
            ? {
                ...event,
                thumbnail: event.thumbnail ?? undefined,
                fromDateType: toGrpcEventDateType(event.fromDateType),
                fromYear: event.fromYear ?? undefined,
                fromMonth: event.fromMonth ?? undefined,
                fromDay: event.fromDay ?? undefined,
                toDateType: toGrpcEventDateType(event.toDateType),
                toYear: event.toYear ?? undefined,
                toMonth: event.toMonth ?? undefined,
                toDay: event.toDay ?? undefined,
                categories: [],
                createdAt: TimestampUtil.toTimestamp(event.createdAt),
                updatedAt: TimestampUtil.toTimestamp(event.updatedAt),
              }
            : undefined,
        };
      },
    );
  }

  @WithSpan(tracerName, 'historical_event.get_preview', {
    'operation.type': 'read',
  })
  async getEventPreviewById({
    id,
  }: GetHistoricalEventPreviewRequest): Promise<GetHistoricalEventPreviewResponse> {
    const { data: event } = await this.getEventById({ id });

    const excerpt = getExcerpt(event!.content, 1000);

    return {
      data: {
        ...event!,
        excerpt,
      },
    };
  }

  @WithSpan(tracerName, 'historical_event.get_by_id_and_author', {
    'operation.type': 'read',
  })
  async getAuthorEventById(
    id: string,
    authorId: string,
  ): Promise<GetHistoricalEventResponse> {
    const options = {
      where: { id, authorId },
      include: {
        author: true,
        categories: {
          include: { category: true, event: false },
        },
      },
    } satisfies Parameters<typeof this.prisma.historicalEvent.findUnique>[0];

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
          () => this.prisma.historicalEvent.findUnique(options),
        ) as any;
        return { data: event ?? undefined };
      },
    );
  }

  async updateEvent(
    payload: UpdateHistoricalEventRequest,
  ): Promise<UpdateHistoricalEventResponse> {
    const found = await this.getEventById({ id: payload.id });
    const cleanPayload =
      this.util.removeNestedUndefined<UpdateHistoricalEventRequest>(payload);
    if (this.util.isEmptyObj(cleanPayload)) {
      return { data: { id: payload.id, success: true } };
    }

    const updated = await this.prisma.historicalEvent.update({
      where: { id: payload.id },
      data: {
        ...cleanPayload,
        toDateType: toEventDateType(payload.toDateType),
        fromDateType: toEventDateType(payload.fromDateType),
      },
    });

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return { data: { id: updated.id, success: true } };
  }

  async deleteEvent({
    id,
    authorId,
  }: DeleteHistoricalEventRequest): Promise<DeleteHistoricalEventResponse> {
    const { data: event } = await this.getAuthorEventById(id, authorId);
    if (!event) {
      throw new RpcException({
        message: 'Sự kiện lịch sử không tồn tại',
        statusCode: 404,
      });
    }

    await this.prisma.historicalEvent.delete({ where: { id } });

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return { data: { id, success: true } };
  }
}
