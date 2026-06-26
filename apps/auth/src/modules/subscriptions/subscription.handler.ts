import { Injectable, Logger } from '@nestjs/common';
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
    private readonly logger: Logger,
  ) {}

  async handleSubscriptionCreatedEvent(event: SubscriptionCreatedEvent) {
    const subscription = event.data;
    if (!subscription) {
      throw new RpcException(
        'Subscription data is missing in the event payload',
      );
    }

    const [defaultOrg, premiumOrg, member] = await Promise.all([
      this.authService.getOrganizationBySlug(ORGANIZATION.DEFAULT.SLUG),
      this.authService.getOrganizationBySlug(ORGANIZATION.PREMIUM.SLUG),
      this.prisma.member.findFirst({ where: { userId: subscription.userId } }),
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
        // Upsert the member to ensure they are associated with the premium organization
        this.prisma.member.upsert({
          where: {
            id: member?.id || '',
            userId: subscription.userId,
          },
          create: {
            id: crypto.randomUUID(),
            userId: subscription.userId,
            organizationId: premiumOrg.id,
            createdAt: new Date(),
          },
          update: {
            organizationId: premiumOrg.id,
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
      this.logger.log('Subscription created event handled successfully');
    } catch (error) {
      this.logger.error('Error handling subscription created event:', error);
      throw new RpcException('Failed to handle subscription created event');
    }
  }
}
