import { Controller, Logger } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';

import { HistoricalEventService } from './historical-event.service';
import {
  type CreateHistoricalEventRequest,
  CreateHistoricalEventResponse,
  type DeleteHistoricalEventRequest,
  DeleteHistoricalEventResponse,
  type GetAllHistoricalEventsRequest,
  GetAllHistoricalEventsResponse,
  type GetHistoricalEventPreviewRequest,
  GetHistoricalEventPreviewResponse,
  type GetHistoricalEventRequest,
  GetHistoricalEventResponse,
  HistoricalEventDetail,
  type UpdateHistoricalEventRequest,
  UpdateHistoricalEventResponse,
} from '@phanhotboy/genproto/historical_event_service/historical_events';
import { toGrpcEventDateType } from '@historical-event/helper/dateType.helper';
import { UtilService } from '@phanhotboy/nsv-common';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import { HistoricalEvent } from '@historical-event-prisma';
import { HistoricalEventDbEntity } from './domain/entity/historical-event.db.entity';

@Controller()
export class HistoricalEventController {
  private readonly logger = new Logger(HistoricalEventController.name);

  constructor(
    private readonly historicalEventService: HistoricalEventService,
  ) {}

  @GrpcMethod('HistoricalEventService', 'GetEventPreview')
  async getEventPreview(
    request: GetHistoricalEventPreviewRequest,
  ): Promise<GetHistoricalEventPreviewResponse> {
    this.logger.log(`Getting historical event preview by id: ${request.id}`);
    const result =
      await this.historicalEventService.getEventPreviewById(request);
    return {
      data: { ...this.toGrpcResponse(result), excerpt: result.excerpt },
    };
  }

  @GrpcMethod('HistoricalEventService', 'GetEvent')
  async getEvent(
    request: GetHistoricalEventRequest,
  ): Promise<GetHistoricalEventResponse> {
    this.logger.log(`Getting historical event by id: ${request.id}`);
    const result = await this.historicalEventService.getEventById(request);
    return {
      data: result && this.toGrpcResponse(result),
    };
  }

  @GrpcMethod('HistoricalEventService', 'GetAllEvents')
  async getAllEvents(
    query: GetAllHistoricalEventsRequest,
  ): Promise<GetAllHistoricalEventsResponse> {
    const result = await this.historicalEventService.getEvents(query);
    return {
      data: result.events.map((event) => this.toGrpcResponse(event)),
      pagination: result.pagination,
    };
  }

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

  private toGrpcResponse(
    event: HistoricalEventDbEntity,
  ): HistoricalEventDetail {
    return {
      ...event,
      toDateType: toGrpcEventDateType(event.toDateType),
      fromDateType: toGrpcEventDateType(event.fromDateType),
      createdAt: TimestampUtil.toTimestamp(event.createdAt),
      updatedAt: TimestampUtil.toTimestamp(event.updatedAt),
      author: (event as any).author || undefined,
      categories: (event as any).categories || [],
    };
  }
}
