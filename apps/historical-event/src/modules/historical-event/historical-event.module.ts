import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { RmqModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';
import { HistoricalEventsPgRepo } from './infrastructure/persistence/postgresql/historical-events.repo';
import { IHistoricalEventDbRepo } from '@historical-event/modules/historical-event/domain/repo/historical-event.db.repo';

@Module({
  imports: [RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE })],
  controllers: [HistoricalEventController],
  providers: [
    {
      provide: IHistoricalEventDbRepo,
      useClass: HistoricalEventsPgRepo,
    },
    HistoricalEventService,
  ],
  exports: [HistoricalEventService],
})
export class HistoricalEventModule {}
