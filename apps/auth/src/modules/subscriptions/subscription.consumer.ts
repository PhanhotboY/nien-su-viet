import { Controller } from '@nestjs/common';
import { Ctx, EventPattern, RmqContext } from '@nestjs/microservices';

import {
  getRoutingKey,
  ParsedMessage,
  RmqService,
} from '@phanhotboy/nsv-common';
import { SubscriptionCreatedEvent } from '@auth/events/subscriptionCreated.event';
import { SubscriptionHandler } from './subscription.handler';
import { RetrySystem } from '@phanhotboy/nsv-common/util/retry.util';

@Controller()
export class SubscriptionConsumer {
  constructor(
    private readonly subscriptionHandler: SubscriptionHandler,
    private readonly retrier: RetrySystem,
    private readonly rmqService: RmqService,
  ) {}

  @EventPattern(getRoutingKey(SubscriptionCreatedEvent.name))
  async handleSubscriptionCreatedEvent(
    @Ctx() context: RmqContext,
    @ParsedMessage(SubscriptionCreatedEvent) message: SubscriptionCreatedEvent,
  ) {
    try {
      await this.retrier.execute(
        () => this.subscriptionHandler.handleSubscriptionCreatedEvent(message),
        context,
      );
      await this.rmqService.ack(context);
    } catch (err) {
      // move the message to dead letter queue if it fails after retries
      console.error(
        'Failed to handle subscription created event after retries:',
        err,
      );
      await this.rmqService.nack(context);
    }
  }
}
