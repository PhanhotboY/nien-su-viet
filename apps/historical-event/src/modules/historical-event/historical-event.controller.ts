import {
  Body,
  Controller,
  Delete,
  Get,
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
} from '@phanhotboy/nsv-common';

@Controller('historical-events')
export class HistoricalEventController {
  constructor(
    private readonly historicalEventService: HistoricalEventService,
  ) {}

  @Get()
  @Permissions({ historicalEvent: ['read'] })
  @Serialize(HistoricalEventBriefResponseDto)
  getAllHistoricalEvents(@Query() query: HistoricalEventQueryDto) {
    return this.historicalEventService.getEvents(query);
  }

  @Post()
  @Permissions({ historicalEvent: ['create'] })
  createHistoricalEvent(
    @Body() event: HistoricalEventBaseCreateDto,
    @CurrentUser('userId') authorId: string,
  ) {
    return this.historicalEventService.createEvent(authorId, event);
  }

  @Get(':id/preview')
  @Permissions({ historicalEvent: ['read'] })
  @Serialize(HistoricalEventPreviewResponseDto)
  getHistoricalEventPreviewById(@Param('id') id: string) {
    return this.historicalEventService.getEventById(id);
  }

  @Get(':id')
  @Permissions({ historicalEvent: ['read'] })
  @Serialize(HistoricalEventDetailResponseDto)
  getHistoricalEventById(@Param('id') id: string) {
    return this.historicalEventService.getEventById(id);
  }

  @Put(':id')
  @Permissions({ historicalEvent: ['update'] })
  @Serialize(HistoricalEventBaseDto)
  updateHistoricalEvent(
    @Param('id') id: string,
    @Body() event: HistoricalEventBaseUpdateDto,
  ) {
    return this.historicalEventService.updateEvent(id, event);
  }

  @Delete(':id')
  @Permissions({ historicalEvent: ['delete'] })
  deleteHistoricalEvent(
    @Param('id') id: string,
    @CurrentUser('userId') userId: string,
  ) {
    return this.historicalEventService.deleteEvent(id, userId);
  }
}
