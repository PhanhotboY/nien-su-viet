import { Inject, Injectable, NotFoundException } from '@nestjs/common';

import { PrismaService } from '@historical-event/database';
import {
  RedisService,
  UtilService,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import {
  HistoricalEventQueryDto,
  HistoricalEventBriefResponseDto,
  HistoricalEventDetailResponseDto,
  HistoricalEventBaseCreateDto,
  HistoricalEventBaseUpdateDto,
} from '@phanhotboy/nsv-common/dto/historical-event';

@Injectable()
export class HistoricalEventService {
  private readonly cachePrefix = 'historical-event';
  private readonly cacheKey: string;

  constructor(
    private readonly prisma: PrismaService,
    private readonly util: UtilService,
    @Inject(RedisService)
    private readonly redisService: RedisServiceType,
  ) {
    this.cacheKey = this.util.genCacheKey(this.cachePrefix);
  }

  async createEvent(
    authorId: string,
    payload: HistoricalEventBaseCreateDto,
  ): Promise<HistoricalEventBriefResponseDto> {
    const role = await this.prisma.historicalEvent.create({
      data: { ...payload, authorId },
    });

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return role;
  }

  async getEvents(query: HistoricalEventQueryDto) {
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

    const options: Parameters<typeof this.prisma.historicalEvent.findMany>[0] =
      {
        where: {},
        skip: (page - 1) * limit,
        take: limit,
        orderBy: {
          [sortBy]: sortOrder,
        },
      };

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
          this.prisma.historicalEvent.findMany(options),
          this.prisma.historicalEvent.count({ where: options.where }),
        ]);

        return {
          data: events,
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

  async getEventById(id: string): Promise<HistoricalEventDetailResponseDto> {
    const options = {
      where: { id },
      include: {
        author: true,
        categories: true,
      },
    };

    return await this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
        notFoundMessage: 'Sự kiện lịch sử không tồn tại',
      },
      () => this.prisma.historicalEvent.findUnique(options),
    );
  }

  async getAuthorEventById(
    id: string,
    authorId: string,
  ): Promise<HistoricalEventDetailResponseDto> {
    const options = {
      where: { id, authorId },
      include: {
        author: true,
        categories: true,
      },
    } satisfies Parameters<typeof this.prisma.historicalEvent.findUnique>[0];

    return this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
        notFoundMessage: 'Sự kiện lịch sử không tồn tại',
      },
      () => this.prisma.historicalEvent.findUnique(options),
    );
  }

  async updateEvent(id: string, payload: HistoricalEventBaseUpdateDto) {
    const found = await this.getEventById(id);
    const cleanPayload =
      this.util.removeNestedUndefined<HistoricalEventBaseUpdateDto>(payload);
    if (this.util.isEmptyObj(cleanPayload)) {
      return found;
    }

    const updated = await this.prisma.historicalEvent.update({
      where: { id },
      data: cleanPayload,
    });

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return updated;
  }

  async deleteEvent(id: string, userId: string) {
    const event = await this.getAuthorEventById(id, userId);
    if (!event) {
      throw new NotFoundException('Sự kiện lịch sử không tồn tại');
    }

    await this.prisma.historicalEvent.delete({ where: { id } });

    // Clear cache
    await this.redisService.del(this.cacheKey);

    return { success: true };
  }
}
