import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { TCP_SERVICE } from '@phanhotboy/constants/tcp-service.constant';
import { catchError, firstValueFrom, throwError, timeout } from 'rxjs';
import {
  HistoricalEventBriefResponseDto,
  HistoricalEventQueryDto,
} from '../historical-event/dto';
import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import { HISTORICAL_EVENT_MESSAGE_PATTERN } from '@phanhotboy/constants/historical-event.message-pattern';
import { PaginatedResponseDto } from '@phanhotboy/nsv-common';

@Injectable()
export class AuthService {
  constructor() {}

  // async getEvents(
  //   query: HistoricalEventQueryDto,
  // ): Promise<PaginatedResponseDto<HistoricalEventBriefResponseDto>> {
  //   return MicroserviceErrorHandler.handleAsyncCall(
  //     () =>
  //       firstValueFrom(
  //         this.heventClient
  //           .send(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_ALL_EVENTS, query)
  //           .pipe(
  //             timeout(10000),
  //             catchError((error) => throwError(() => error)),
  //           ),
  //       ),
  //     'get events',
  //     'get events',
  //   );
  // }
}
