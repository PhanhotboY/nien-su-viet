import { Controller, Logger } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';

import { HistoricalEventService } from './historical-event.service';
import {
  type CreateHistoricalEventRequest,
  CreateHistoricalEventResponse,
  type DeleteHistoricalEventRequest,
  DeleteHistoricalEventResponse,
  type UpdateHistoricalEventRequest,
  UpdateHistoricalEventResponse,
} from '@phanhotboy/genproto/historical_event_service/historical_events';

@Controller()
export class HistoricalEventController {
  private readonly logger = new Logger(HistoricalEventController.name);

  constructor(
    private readonly historicalEventService: HistoricalEventService,
  ) {}

  @GrpcMethod('HistoricalEventService', 'CreateEvent')
  async createEvent(
    payload: CreateHistoricalEventRequest,
  ): Promise<CreateHistoricalEventResponse> {
    const result = await this.historicalEventService.createEvent(payload);
    return { data: result };
  }

  @GrpcMethod('HistoricalEventService', 'UpdateEvent')
  async updateEvent(
    payload: UpdateHistoricalEventRequest,
  ): Promise<UpdateHistoricalEventResponse> {
    const result = await this.historicalEventService.updateEvent(payload);

    return { data: result };
  }

  @GrpcMethod('HistoricalEventService', 'DeleteEvent')
  async deleteEvent(
    request: DeleteHistoricalEventRequest,
  ): Promise<DeleteHistoricalEventResponse> {
    this.logger.log(`Deleting historical event with id: ${request.id}`);
    const result = await this.historicalEventService.deleteEvent(request);
    return { data: result };
  }
}
