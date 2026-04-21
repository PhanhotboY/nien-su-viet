import { Controller, Logger } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';

import { HistoricalEventService } from './historical-event.service';
import type {
  CreateHistoricalEventRequest,
  DeleteHistoricalEventRequest,
  GetAllHistoricalEventsRequest,
  GetHistoricalEventPreviewRequest,
  GetHistoricalEventRequest,
  UpdateHistoricalEventRequest,
} from '@phanhotboy/genproto/historical_event_service/historical_events';

@Controller()
export class HistoricalEventController {
  private readonly logger = new Logger(HistoricalEventController.name);

  constructor(
    private readonly historicalEventService: HistoricalEventService,
  ) {}

  @GrpcMethod('HistoricalEventService', 'GetEventPreview')
  getEventPreview(request: GetHistoricalEventPreviewRequest) {
    this.logger.log(`Getting historical event preview by id: ${request.id}`);
    return this.historicalEventService.getEventPreviewById(request);
  }

  @GrpcMethod('HistoricalEventService', 'GetEvent')
  getEvent(request: GetHistoricalEventRequest) {
    this.logger.log(`Getting historical event by id: ${request.id}`);
    return this.historicalEventService.getEventById(request);
  }

  @GrpcMethod('HistoricalEventService', 'GetAllEvents')
  getAllEvents(query: GetAllHistoricalEventsRequest) {
    return this.historicalEventService.getEvents(query);
  }

  @GrpcMethod('HistoricalEventService', 'CreateEvent')
  createEvent(payload: CreateHistoricalEventRequest) {
    return this.historicalEventService.createEvent(payload);
  }

  @GrpcMethod('HistoricalEventService', 'UpdateEvent')
  updateEvent(payload: UpdateHistoricalEventRequest) {
    return this.historicalEventService.updateEvent(payload);
  }

  @GrpcMethod('HistoricalEventService', 'DeleteEvent')
  deleteEvent(request: DeleteHistoricalEventRequest) {
    this.logger.log(`Deleting historical event with id: ${request.id}`);
    return this.historicalEventService.deleteEvent(request);
  }
}
