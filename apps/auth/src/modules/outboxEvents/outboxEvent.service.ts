import { Injectable } from '@nestjs/common';

import { PrismaService } from '@auth/database';
import { OutboxEventStatus } from '@auth-prisma';
import { calculateRetryDelay } from '@phanhotboy/nsv-common';

@Injectable()
export class OutboxEventService {
  constructor(private readonly prisma: PrismaService) {}

  async addEventToOutbox(event: any) {
    return this.prisma.outboxEvent.create({
      data: {
        aggregateType: event.aggregateType,
        aggregateId: event.aggregateId,
        eventType: event.eventType,
        payload: event.payload,
        status: OutboxEventStatus.PENDING,
      },
    });
  }

  async findByEventType(eventType: string) {
    return this.prisma.outboxEvent.findMany({
      where: {
        eventType: eventType,
      },
    });
  }

  async findById(eventId: string) {
    return this.prisma.outboxEvent.findUnique({
      where: {
        id: eventId,
      },
    });
  }

  async getPendingEvents({ limit = 10 }: { limit?: number } = {}) {
    const now = new Date();
    return this.prisma.outboxEvent.findMany({
      where: {
        OR: [
          { status: OutboxEventStatus.PENDING },
          { status: OutboxEventStatus.FAILED, nextRetryAt: { lte: now } },
        ],
      },
      orderBy: {
        createdAt: 'asc',
      },
      take: limit,
      select: {
        id: true,
        eventType: true,
        payload: true,
        retryCount: true,
      },
    });
  }

  async markEventAsProcessed(eventId: string) {
    return this.prisma.outboxEvent.update({
      where: {
        id: eventId,
      },
      data: {
        status: OutboxEventStatus.PUBLISHED,
        publishedAt: new Date(),
      },
    });
  }

  async markEventAsFailed(eventId: string, error: any) {
    const event = await this.findById(eventId);
    if (!event) {
      return;
    }

    const retryCount = event.retryCount + 1;
    let status: OutboxEventStatus = OutboxEventStatus.FAILED;
    let nextRetryAt = event.nextRetryAt;

    if (retryCount >= 10) {
      status = OutboxEventStatus.DEAD;
    } else {
      nextRetryAt = new Date(Date.now() + calculateRetryDelay({ retryCount }));
    }

    return this.prisma.outboxEvent.update({
      where: {
        id: eventId,
      },
      data: {
        status: status as any, // TODO: fix type
        retryCount: retryCount,
        lastError: error.message,
        nextRetryAt: nextRetryAt,
      },
    });
  }
}
