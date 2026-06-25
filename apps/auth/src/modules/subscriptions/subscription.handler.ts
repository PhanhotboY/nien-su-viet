import { Injectable } from '@nestjs/common';
import { RmqContext, RpcException } from '@nestjs/microservices';

import { AuthService } from '../auth';
import { PrismaService } from '@auth/database';
import { OutboxEventStatus } from '@auth-prisma';
import { SubscriptionCreatedEvent } from '@auth/events/subscriptionCreated.event';
import { ORGANIZATION } from '@phanhotboy/constants/organization.constant';
import { getRoutingKey, RmqService } from '@phanhotboy/nsv-common';
import { OutboxEventService } from '../outboxEvents/outboxEvent.service';

@Injectable()
export class SubscriptionHandler {
  constructor(
    private readonly authService: AuthService,
    private readonly prisma: PrismaService,
  ) {}

  async handleSubscriptionCreatedEvent(event: SubscriptionCreatedEvent) {
    const subscription = event.data;
    if (!subscription) {
      throw new RpcException(
        'Subscription data is missing in the event payload',
      );
    }

    const [defaultOrg, premiumOrg] = await Promise.all([
      this.authService.getOrganizationBySlug(ORGANIZATION.DEFAULT.SLUG),
      this.authService.getOrganizationBySlug(ORGANIZATION.PREMIUM.SLUG),
    ]);
    if (!defaultOrg || !premiumOrg) {
      throw new RpcException('Organization not found');
    }

    try {
      const res = await this.prisma.$transaction([
        // Fail if the event has already been processed to ensure idempotency
        this.prisma.processedEvent.create({
          data: {
            consumerName: 'SubscriptionConsumer',
            messageId: event.eventId,
            processedAt: new Date(),
          },
        }),
        // remove user from default org
        this.prisma.organization.update({
          where: { id: defaultOrg.id },
          data: {
            members: {
              disconnect: { id: subscription.userId },
            },
          },
        }),
        // add user to premium org
        this.prisma.organization.update({
          where: { id: premiumOrg.id },
          data: {
            members: {
              connect: { id: subscription.userId },
            },
          },
        }),
        // add outbox event -> published by outbox event worker
        this.prisma.outboxEvent.create({
          data: {
            eventType: getRoutingKey(SubscriptionCreatedEvent.name),
            aggregateType: 'user',
            aggregateId: subscription.userId,
            payload: JSON.stringify(event.data),
            status: OutboxEventStatus.PENDING,
          },
        }),
      ]);
    } catch (error) {
      console.error('Error handling subscription created event:', error);
    }
  }
}
