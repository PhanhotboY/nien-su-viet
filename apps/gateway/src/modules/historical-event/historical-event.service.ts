import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { firstValueFrom, timeout, catchError, throwError } from 'rxjs';

import {
  HistoricalEventQueryDto,
  HistoricalEventBriefResponseDto,
  HistoricalEventDetailResponseDto,
  HistoricalEventBaseCreateDto,
  HistoricalEventBaseUpdateDto,
  HistoricalEventPreviewResponseDto,
} from './dto';
import { TCP_SERVICE } from '@phanhotboy/constants/tcp-service.constant';
import { HISTORICAL_EVENT_MESSAGE_PATTERN } from '@phanhotboy/constants/historical-event.message-pattern';
import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import { PaginatedResponseDto } from '@phanhotboy/nsv-common';

@Injectable()
export class HistoricalEventService {
  private readonly serviceName: 'Historical Event Service';

  constructor(
    @Inject(TCP_SERVICE.HISTORICAL_EVENT.NAME)
    private readonly heventClient: ClientProxy,
  ) {}

  async createEvent(authorId: string, payload: HistoricalEventBaseCreateDto) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .send(HISTORICAL_EVENT_MESSAGE_PATTERN.CREATE_EVENT, {
              authorId,
              payload,
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

  async getEvents(
    query: HistoricalEventQueryDto,
  ): Promise<PaginatedResponseDto<HistoricalEventBriefResponseDto>> {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .send(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_ALL_EVENTS, query)
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get events',
      this.serviceName,
    );
  }

  async getEventById(id: string): Promise<HistoricalEventDetailResponseDto> {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .send(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_EVENT_BY_ID, id)
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get event by id',
      this.serviceName,
    );
  }

  async getEventPreviewById(
    id: string,
  ): Promise<HistoricalEventPreviewResponseDto> {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .send(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_EVENT_PREVIEW_BY_ID, id)
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get event preview by id',
      this.serviceName,
    );
  }

  async updateEvent(id: string, payload: HistoricalEventBaseUpdateDto) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .send(HISTORICAL_EVENT_MESSAGE_PATTERN.UPDATE_EVENT, {
              id,
              payload,
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

  async deleteEvent(id: string, userId: string) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .send(HISTORICAL_EVENT_MESSAGE_PATTERN.DELETE_EVENT, { id, userId })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'delete event',
      this.serviceName,
    );
  }
}
