import { PrismaService } from '@auth/database';
import { Injectable } from '@nestjs/common';

@Injectable()
export class ProcessedEventService {
  constructor(private readonly prisma: PrismaService) {}

  async addProcessedEvent(eventId: string, consumerName: string) {
    return this.prisma.processedEvent.create({
      data: {
        messageId: eventId,
        processedAt: new Date(),
        consumerName: consumerName,
      },
    });
  }

  async markEventAsProcessed(eventId: string) {
    return this.prisma.processedEvent.update({
      where: {
        id: eventId,
      },
      data: {
        processedAt: new Date(),
      },
    });
  }

  async isEventProcessed(eventId: string): Promise<boolean> {
    const event = await this.prisma.processedEvent.findFirst({
      where: {
        messageId: eventId,
        processedAt: {
          lte: new Date(),
        },
      },
    });
    return event !== null;
  }
}
