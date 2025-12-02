import { Injectable } from '@nestjs/common';
import { RmqOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '../providers';

export const RABBITMQ_OPTIONS = 'RABBITMQ_OPTIONS';

@Injectable()
export class RmqService {
  constructor(private readonly config: ConfigService) {}

  private getRabbitMQUri(): string {
    const rmqUrl = this.config.get('RABBITMQ_URL');
    if (!rmqUrl) {
      throw new Error('Missing RabbitMQ configuration in .env file.');
    }
    return rmqUrl;
  }

  /**
   * Get RMQ options for microservice consumer
   */
  getOptions(queue: string, noAck = false): RmqOptions {
    return {
      transport: Transport.RMQ,
      options: {
        urls: [this.getRabbitMQUri()],
        queue,
        noAck,
        wildcards: true,
        exchange: 'events',
        exchangeType: 'topic',
        persistent: true,
        queueOptions: {
          durable: true,
        },
      },
    };
  }

  ack(context: any) {
    const channel = context.getChannelRef();
    const originalMsg = context.getMessage();
    channel.ack(originalMsg);
  }
}
