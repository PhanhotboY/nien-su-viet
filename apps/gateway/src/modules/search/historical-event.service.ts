import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import type { ClientGrpc } from '@nestjs/microservices';
import { firstValueFrom, timeout, catchError, throwError } from 'rxjs';

import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  HistoricalEventServiceClient,
  EventDateType,
} from '@phanhotboy/genproto/search_service/historical_events';
import { HistoricalEventQueryDto } from './dto';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import { GRPC_SERVICE, HISTORICAL_EVENT } from '@phanhotboy/constants';

@Injectable()
export class SearchHistoricalEventService {
  private readonly serviceName = 'Historical Event Service';
  private readonly microserviceErrorHandler: MicroserviceErrorHandler;
  private heventClient: HistoricalEventServiceClient;

  constructor(
    @Inject(GRPC_SERVICE.SEARCH.NAME)
    private readonly client: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(logger);
    this.heventClient = this.client.getService<HistoricalEventServiceClient>(
      'HistoricalEventService',
    );
  }

  async getEvents(query: HistoricalEventQueryDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.heventClient
            .getAllEvents({
              ...query,
              categoryIds: query.categoryIds || [],
              /** Filter by created date range */
              createdAtFrom: TimestampUtil.toTimestamp(query.createdAtFrom),
              createdAtTo: TimestampUtil.toTimestamp(query.createdAtTo),
              /** Filter by updated date range */
              updatedAtFrom: TimestampUtil.toTimestamp(query.updatedAtFrom),
              updatedAtTo: TimestampUtil.toTimestamp(query.updatedAtTo),
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
}
