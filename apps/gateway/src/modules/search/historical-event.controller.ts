import { Controller, Get, Inject, Param, Query } from '@nestjs/common';
import { Throttle } from '@nestjs/throttler';

import { SearchHistoricalEventService } from './historical-event.service';
import {
  HistoricalEventBriefResponseDto,
  HistoricalEventDetailResponseDto,
  HistoricalEventPreviewResponseDto,
  HistoricalEventQueryDto,
} from './dto';
import {
  RedisService,
  type RedisServiceType,
  ConfigService,
  Serialize,
  ApiOkSerializedResponse,
  ApiOkSerializedPaginatedResponse,
} from '@phanhotboy/nsv-common';
import { RATE_LIMIT } from '@gateway/config';
import { Public, Permissions, CurrentUser } from '@gateway/common/decorators';

@Controller('search/historical-events')
export class SearchHistoricalEventController {
  constructor(
    private readonly historicalEventService: SearchHistoricalEventService,
    private readonly config: ConfigService,
  ) {}

  @Get(':id/preview')
  @Public()
  @Serialize(HistoricalEventPreviewResponseDto)
  @ApiOkSerializedResponse(HistoricalEventPreviewResponseDto)
  getHistoricalEventPreviewById(@Param('id') id: string) {
    return this.historicalEventService.getEventPreviewById(id);
  }

  @Get(':id')
  @Public()
  @Serialize(HistoricalEventDetailResponseDto)
  @ApiOkSerializedResponse(HistoricalEventDetailResponseDto)
  getHistoricalEventById(@Param('id') id: string) {
    return this.historicalEventService.getEventById(id);
  }

  @Get()
  @Public()
  @Serialize(HistoricalEventBriefResponseDto)
  @ApiOkSerializedPaginatedResponse(HistoricalEventBriefResponseDto)
  getAllHistoricalEvents(@Query() query: HistoricalEventQueryDto) {
    return this.historicalEventService.getEvents(query) as any;
  }
}
