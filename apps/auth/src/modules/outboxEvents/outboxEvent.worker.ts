import { Injectable, Logger } from '@nestjs/common';
import { Cron, CronExpression } from '@nestjs/schedule';

import { OutboxEventService } from './outboxEvent.service';
import { RmqService } from '@phanhotboy/nsv-common';

@Injectable()
export class OutboxEventWorker {
  constructor(
    private readonly outboxEventService: OutboxEventService,
    private readonly rmq: RmqService,
    private readonly logger: Logger,
  ) {}

  @Cron(CronExpression.EVERY_5_SECONDS)
  async publishPendingEvents() {
    this.logger.log('Publishing pending events from outbox...');
    const pendingEvents = await this.outboxEventService.getPendingEvents({
      limit: 10,
    });

    for (const event of pendingEvents) {
      try {
        await this.rmq.publish(event.eventType, event.payload?.toString());

        await this.outboxEventService.markEventAsProcessed(event.id);
        // TODO: Handle failed to publish event after successfully publish
      } catch (error) {
        console.error(`Failed to publish event ${event.id}:`, error);
        await this.outboxEventService.markEventAsFailed(event.id, error);
      }
    }
  }
}
