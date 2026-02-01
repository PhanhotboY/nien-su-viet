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
} from './dto';
import {
  Permissions,
  CurrentUser,
  Serialize,
  Public,
  RedisService,
  type RedisServiceType,
  ConfigService,
} from '@phanhotboy/nsv-common';
import { ThrottleMap } from '@gateway/common/decorators';
import { RATE_LIMIT } from '@gateway/config';

@Controller('historical-events')
export class HistoricalEventController {
  private readonly routePath: 'api/v1/historical-events';

  constructor(
    private readonly historicalEventService: HistoricalEventService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
    private readonly config: ConfigService,
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
  @ThrottleMap(RATE_LIMIT.INTERNAL)
  @Permissions({ historicalEvent: ['create'] })
  async createHistoricalEvent(
    @Body() event: HistoricalEventBaseCreateDto,
    @CurrentUser('id') authorId: string,
  ) {
    await this.redis.mdel(this.routePath);
    return this.historicalEventService.createEvent(authorId, event);
  }

  @Put(':id')
  @ThrottleMap(RATE_LIMIT.INTERNAL)
  @Permissions({ historicalEvent: ['update'] })
  @Serialize(HistoricalEventBaseDto)
  async updateHistoricalEvent(
    @Param('id') id: string,
    @Body() event: HistoricalEventBaseUpdateDto,
  ) {
    return this.historicalEventService.updateEvent(id, event);
  }

  @Delete(':id')
  @ThrottleMap(RATE_LIMIT.INTERNAL)
  @Permissions({ historicalEvent: ['delete'] })
  async deleteHistoricalEvent(
    @Param('id') id: string,
    @CurrentUser('userId') userId: string,
  ) {
    return this.historicalEventService.deleteEvent(id, userId);
  }
}
