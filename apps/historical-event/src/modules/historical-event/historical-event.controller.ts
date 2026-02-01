import { Controller, Logger } from '@nestjs/common';
import { MessagePattern, Payload } from '@nestjs/microservices';

import { HistoricalEventService } from './historical-event.service';
import {
  HistoricalEventBaseCreateDto,
  HistoricalEventBaseUpdateDto,
  HistoricalEventQueryDto,
} from './dto';
import { Permissions, CurrentUser, Public } from '@phanhotboy/nsv-common';
import { HISTORICAL_EVENT_MESSAGE_PATTERN } from '@phanhotboy/constants/historical-event.message-pattern';

@Controller('historical-events')
export class HistoricalEventController {
  private readonly logger = new Logger(HistoricalEventController.name);

  constructor(
    private readonly historicalEventService: HistoricalEventService,
  ) {}

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_EVENT_PREVIEW_BY_ID)
  @Public()
  getHistoricalEventPreviewById(@Payload('id') id: string) {
    this.logger.log(`Getting historical event preview by id: ${id}`);
    return this.historicalEventService.getEventPreviewById(id);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_EVENT_BY_ID)
  @Public()
  getHistoricalEventById(@Payload('id') id: string) {
    return this.historicalEventService.getEventById(id);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_ALL_EVENTS)
  @Public()
  getAllHistoricalEvents(query: HistoricalEventQueryDto) {
    this.logger.log(
      `Getting all historical events with query: ${JSON.stringify(query)}`,
    );
    return this.historicalEventService.getEvents(query);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.CREATE_EVENT)
  @Permissions({ historicalEvent: ['create'] })
  async createHistoricalEvent(
    event: HistoricalEventBaseCreateDto,
    @CurrentUser('id') authorId: string,
  ) {
    return this.historicalEventService.createEvent(authorId, event);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.UPDATE_EVENT)
  @Permissions({ historicalEvent: ['update'] })
  async updateHistoricalEvent(
    @Payload('id') id: string,
    @Payload() event: HistoricalEventBaseUpdateDto,
  ) {
    return this.historicalEventService.updateEvent(id, event);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.DELETE_EVENT)
  @Permissions({ historicalEvent: ['delete'] })
  async deleteHistoricalEvent(
    @Payload('id') id: string,
    @CurrentUser('userId') userId: string,
  ) {
    return this.historicalEventService.deleteEvent(id, userId);
  }
}
