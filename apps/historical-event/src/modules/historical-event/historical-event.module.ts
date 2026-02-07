import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { RmqModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';
import { UserModule } from '../user';

@Module({
  imports: [
    RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE }),
    UserModule,
  ],
  controllers: [HistoricalEventController],
  providers: [HistoricalEventService],
})
export class HistoricalEventModule {}
