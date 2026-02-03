import { Controller, Logger } from '@nestjs/common';
import { MessagePattern, Payload, Transport } from '@nestjs/microservices';

import { HistoricalEventService } from './historical-event.service';
import {
  HistoricalEventBaseCreateDto,
  HistoricalEventBaseUpdateDto,
  HistoricalEventQueryDto,
} from './dto';
import { HISTORICAL_EVENT_MESSAGE_PATTERN } from '@phanhotboy/constants/historical-event.message-pattern';

@Controller('historical-events')
export class HistoricalEventController {
  private readonly logger = new Logger(HistoricalEventController.name);

  constructor(
    private readonly historicalEventService: HistoricalEventService,
  ) {}

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_EVENT_PREVIEW_BY_ID)
  getHistoricalEventPreviewById(@Payload('id') id: string) {
    this.logger.log(`Getting historical event preview by id: ${id}`);
    return this.historicalEventService.getEventPreviewById(id);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_EVENT_BY_ID)
  getHistoricalEventById(@Payload('id') id: string) {
    return this.historicalEventService.getEventById(id);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.GET_ALL_EVENTS)
  getAllHistoricalEvents(@Payload('query') query: HistoricalEventQueryDto) {
    return this.historicalEventService.getEvents(query);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.CREATE_EVENT)
  async createHistoricalEvent(
    @Payload('authorId') authorId: string,
    @Payload('payload') event: HistoricalEventBaseCreateDto,
  ) {
    return this.historicalEventService.createEvent(authorId, event);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.UPDATE_EVENT, {
    transport: Transport.TCP,
  })
  async updateHistoricalEvent(
    @Payload('id') id: string,
    @Payload('payload') event: HistoricalEventBaseUpdateDto,
  ) {
    return this.historicalEventService.updateEvent(id, event);
  }

  @MessagePattern(HISTORICAL_EVENT_MESSAGE_PATTERN.DELETE_EVENT, {
    transport: Transport.TCP,
  })
  async deleteHistoricalEvent(
    @Payload('id') id: string,
    @Payload('authorId') authorId: string,
  ) {
    this.logger.log(`Deleting historical event with id: ${id}`);
    await this.historicalEventService.deleteEvent(id, authorId);
  }
}
