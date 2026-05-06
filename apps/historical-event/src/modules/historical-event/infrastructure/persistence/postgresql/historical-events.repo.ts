import { HistoricalEvent } from '@historical-event-prisma';
import { PrismaService } from '@historical-event/database';
import { HistoricalEventDbEntity } from '@historical-event/modules/historical-event/domain/entity/historical-event.db.entity';
import { IHistoricalEventDbRepo } from '@historical-event/modules/historical-event/domain/repo/historical-event.db.repo';
import { toEventDateType } from '@historical-event/helper/dateType.helper';
import {
  type CreateHistoricalEventRequest,
  type UpdateHistoricalEventRequest,
} from '@phanhotboy/genproto/historical_event_service/historical_events';
import { Injectable } from '@nestjs/common';
import {
  isEmptyObj,
  removeNestedUndefined,
} from '@phanhotboy/nsv-common/util/object.util';

type HistoricalEventListQuery = Parameters<
  PrismaService['historicalEvent']['findMany']
>[0];
type HistoricalEventUniqueQuery = Parameters<
  PrismaService['historicalEvent']['findUnique']
>[0];

@Injectable()
export class HistoricalEventsPgRepo implements IHistoricalEventDbRepo {
  constructor(private readonly prisma: PrismaService) {}

  async createHistoricalEvent(payload: CreateHistoricalEventRequest) {
    const event = await this.prisma.historicalEvent.create({
      data: {
        ...payload,
        toDateType: toEventDateType(payload.toDateType),
        fromDateType: toEventDateType(payload.fromDateType),
      },
    });

    return this.mapToDbEntity(event);
  }

  async getHistoricalEventById(options: HistoricalEventUniqueQuery) {
    const event = await this.prisma.historicalEvent.findUnique({
      ...options,
    });
    if (!event) {
      return null;
    }

    return this.mapToDbEntity(event);
  }

  async updateHistoricalEvent(
    id: string,
    payload: Omit<UpdateHistoricalEventRequest, 'id'>,
  ) {
    const cleanPayload =
      removeNestedUndefined<UpdateHistoricalEventRequest>(payload);
    if (isEmptyObj(cleanPayload)) {
      return id;
    }
    await this.prisma.historicalEvent.update({
      where: { id },
      data: {
        ...cleanPayload,
        toDateType: toEventDateType(payload.toDateType),
        fromDateType: toEventDateType(payload.fromDateType),
      },
    });

    return id;
  }

  async deleteHistoricalEvent(id: string) {
    await this.prisma.historicalEvent.delete({ where: { id } });
  }

  async searchHistoricalEvents(options: HistoricalEventListQuery) {
    const events = await this.prisma.historicalEvent.findMany(options);

    return events.map((event) => {
      return this.mapToDbEntity(event);
    });
  }

  async countHistoricalEvents(
    query: HistoricalEventListQuery,
  ): Promise<number> {
    return this.prisma.historicalEvent.count({ where: query?.where });
  }

  private mapToDbEntity(event: HistoricalEvent): HistoricalEventDbEntity {
    if (!event) {
      return null as any;
    }

    return new HistoricalEventDbEntity({
      ...event,
      thumbnail: event.thumbnail ?? undefined,
      fromDay: event.fromDay ?? undefined,
      fromMonth: event.fromMonth ?? undefined,
      toDay: event.toDay ?? undefined,
      toMonth: event.toMonth ?? undefined,
      toYear: event.toYear ?? undefined,
    });
  }
}
