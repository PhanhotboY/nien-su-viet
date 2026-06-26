import { Module } from '@nestjs/common';

import { RmqModule } from '@phanhotboy/nsv-common';
import { OutboxEventService } from './outboxEvent.service';
import { OutboxEventWorker } from './outboxEvent.worker';

@Module({
  imports: [RmqModule.register()],
  providers: [OutboxEventService, OutboxEventWorker],
  exports: [OutboxEventService],
})
export class OutboxEventModule {}
