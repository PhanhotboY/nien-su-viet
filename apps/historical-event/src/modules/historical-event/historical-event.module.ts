import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { RMQ, RmqModule } from '@phanhotboy/nsv-common';

@Module({
  imports: [RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE })],
  controllers: [HistoricalEventController],
  providers: [HistoricalEventService],
})
export class HistoricalEventModule {}
