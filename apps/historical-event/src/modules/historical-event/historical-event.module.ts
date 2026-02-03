import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { RmqModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';

@Module({
  imports: [RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE })],
  controllers: [HistoricalEventController],
  providers: [HistoricalEventService],
})
export class HistoricalEventModule {}
