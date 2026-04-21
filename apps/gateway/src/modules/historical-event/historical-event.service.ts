import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import type { ClientGrpc } from '@nestjs/microservices';
import { firstValueFrom, timeout, catchError, throwError } from 'rxjs';

import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  HistoricalEventServiceClient,
  EventDateType,
} from '@phanhotboy/genproto/historical_event_service/historical_events';
import {
  HistoricalEventBaseCreateDto,
  HistoricalEventBaseDto,
  HistoricalEventBaseUpdateDto,
  HistoricalEventQueryDto,
} from './dto';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import { HISTORICAL_EVENT } from '@phanhotboy/constants';

@Injectable()
export class HistoricalEventService {
  private readonly serviceName = 'Historical Event Service';
  private readonly microserviceErrorHandler: MicroserviceErrorHandler;
  private heventClient: HistoricalEventServiceClient;

  constructor(
    @Inject('HISTORICAL_EVENT_SERVICE')
    private readonly client: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(logger);
    this.heventClient = this.client.getService<HistoricalEventServiceClient>(
      'HistoricalEventService',
    );
  }

  async createEvent(authorId: string, payload: HistoricalEventBaseCreateDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .createEvent({
              ...payload,
              authorId,
              fromDateType: toEventDateType(payload.fromDateType),
              toDateType: toEventDateType(payload.toDateType),
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'create historical event',
      this.serviceName,
    );
  }

  async getEvents(query: HistoricalEventQueryDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .getAllEvents({
              page: query.page,
              limit: query.limit,
              search: query.search,
              sortBy: query.sortBy,
              /** "asc" or "desc" */
              sortOrder: query.sortOrder,
              authorId: query.authorId,
              /** Filter by category */
              categoryIds: query.categoryIds || [],
              /** Filter by from date range */
              fromYear: query.fromYear,
              fromMonth: query.fromMonth,
              fromDay: query.fromDay,
              /** Filter by to date range */
              toYear: query.toYear,
              toMonth: query.toMonth,
              toDay: query.toDay,
              /** Search specific year */
              searchYear: query.searchYear,
              /** Filter by created date range */
              createdAtFrom:
                query.createdAtFrom &&
                TimestampUtil.toTimestamp(query.createdAtFrom),
              createdAtTo:
                query.createdAtTo &&
                TimestampUtil.toTimestamp(query.createdAtTo),
              /** Filter by updated date range */
              updatedAtFrom:
                query.updatedAtFrom &&
                TimestampUtil.toTimestamp(query.updatedAtFrom),
              updatedAtTo:
                query.updatedAtTo &&
                TimestampUtil.toTimestamp(query.updatedAtTo),
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get events',
      this.serviceName,
    );
  }

  async getEventById(id: string) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient.getEvent({ id }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'get event by id',
      this.serviceName,
    );
  }

  async getEventPreviewById(id: string) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient.getEventPreview({ id }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'get event preview by id',
      this.serviceName,
    );
  }

  async updateEvent(id: string, payload: HistoricalEventBaseUpdateDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .updateEvent({
              id,
              name: payload.name,
              thumbnail: payload.thumbnail,
              fromDateType: toEventDateType(payload.fromDateType),
              fromDay: payload.fromDay,
              fromMonth: payload.fromMonth,
              fromYear: payload.fromYear,
              toDateType: toEventDateType(payload.toDateType),
              toDay: payload.toDay,
              toMonth: payload.toMonth,
              toYear: payload.toYear,
              content: payload.content,
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'update event',
      this.serviceName,
    );
  }

  async deleteEvent(id: string, authorId: string) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient.deleteEvent({ id, authorId }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'delete event',
      this.serviceName,
    );
  }
}

function toEventDateType(
  dateType?: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>,
): EventDateType {
  switch (dateType) {
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT:
      return EventDateType.EXACT;
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE:
      return EventDateType.APPROXIMATE;
    default:
      return EventDateType.EVENT_DATE_TYPE_UNSPECIFIED;
  }
}
