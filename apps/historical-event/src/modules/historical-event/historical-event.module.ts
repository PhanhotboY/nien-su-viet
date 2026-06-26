import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { RmqModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';
import { UserModule } from '../user';

@Module({
  imports: [RmqModule.register(), UserModule],
  controllers: [HistoricalEventController],
  providers: [HistoricalEventService],
  exports: [HistoricalEventService],
})
export class HistoricalEventModule {}
