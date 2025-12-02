import {
  Body,
  Controller,
  Delete,
  Get,
  Inject,
  Param,
  Post,
  Put,
  Query,
} from '@nestjs/common';

import { HistoricalEventService } from './historical-event.service';
import {
  HistoricalEventBaseCreateDto,
  HistoricalEventBaseDto,
  HistoricalEventBaseUpdateDto,
  HistoricalEventBriefResponseDto,
  HistoricalEventDetailResponseDto,
  HistoricalEventPreviewResponseDto,
  HistoricalEventQueryDto,
  Permissions,
  CurrentUser,
  Serialize,
  RedisService,
  type RedisServiceType,
  Public,
} from '@phanhotboy/nsv-common';

@Controller('historical-events')
export class HistoricalEventController {
  private routePath = '/api/v1/historical-events*';

  constructor(
    private readonly historicalEventService: HistoricalEventService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
  ) {}

  @Get(':id/preview')
  @Public()
  @Serialize(HistoricalEventPreviewResponseDto)
  getHistoricalEventPreviewById(@Param('id') id: string) {
    return this.historicalEventService.getEventPreviewById(id);
  }

  @Get(':id')
  @Public()
  @Serialize(HistoricalEventDetailResponseDto)
  getHistoricalEventById(@Param('id') id: string) {
    return this.historicalEventService.getEventById(id);
  }

  @Get()
  @Public()
  @Serialize(HistoricalEventBriefResponseDto)
  getAllHistoricalEvents(@Query() query: HistoricalEventQueryDto) {
    return this.historicalEventService.getEvents(query);
  }

  @Post()
  @Permissions({ historicalEvent: ['create'] })
  async createHistoricalEvent(
    @Body() event: HistoricalEventBaseCreateDto,
    @CurrentUser('id') authorId: string,
  ) {
    const res = await this.historicalEventService.createEvent(authorId, event);
    await this.redis.mdel(this.routePath);
    return res;
  }

  @Put(':id')
  @Permissions({ historicalEvent: ['update'] })
  @Serialize(HistoricalEventBaseDto)
  async updateHistoricalEvent(
    @Param('id') id: string,
    @Body() event: HistoricalEventBaseUpdateDto,
  ) {
    const res = await this.historicalEventService.updateEvent(id, event);
    await this.redis.mdel(this.routePath);
    return res;
  }

  @Delete(':id')
  @Permissions({ historicalEvent: ['delete'] })
  async deleteHistoricalEvent(
    @Param('id') id: string,
    @CurrentUser('userId') userId: string,
  ) {
    const res = await this.historicalEventService.deleteEvent(id, userId);
    await this.redis.mdel(this.routePath);
    return res;
  }
}
