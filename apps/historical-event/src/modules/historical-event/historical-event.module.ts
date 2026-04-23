import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { RmqModule, SearchModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';
import { UserModule } from '../user';
import HistoricalEventsPgRepo from './infrastructure/persistence/postgresql/historical-events.repo';
import { HistoricalEventsEsRepo } from './infrastructure/persistence/elasticsearch/historical-events.repo';

@Module({
  imports: [
    RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE }),
    UserModule,
    SearchModule,
  ],
  controllers: [HistoricalEventController],
  providers: [
    HistoricalEventsPgRepo,
    HistoricalEventsEsRepo,
    HistoricalEventService,
  ],
  exports: [HistoricalEventService],
})
export class HistoricalEventModule {}
