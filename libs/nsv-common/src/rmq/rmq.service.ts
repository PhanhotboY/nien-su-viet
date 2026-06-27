import { Inject, Injectable } from '@nestjs/common';
import {
  ClientProxy,
  RmqContext,
  RmqOptions,
  Transport,
} from '@nestjs/microservices';
import { ConfigService } from '../providers';
import {
  getQueueName,
  getRoutingKey,
  getDLXName,
  getDeadRoutingKey,
  keysToSnakeCase,
} from '../util';
import { RMQ } from '@phanhotboy/constants';
import { AmqpConnectionManager } from 'amqp-connection-manager';

export const RABBITMQ_OPTIONS = 'RABBITMQ_OPTIONS';

@Injectable()
export class RmqService {
  constructor(
    private readonly config: ConfigService,
    @Inject(RMQ.TOPIC_EVENTS_EXCHANGE) private readonly rmq: ClientProxy,
  ) {}

  private getRabbitMQUri(): string {
    const rmqUrl = this.config.get('rabbitmq');
    if (!rmqUrl) {
      throw new Error('Missing RabbitMQ configuration in .env file.');
    }
    return rmqUrl;
  }

  /**
   * Get RMQ options for microservice consumer
   */
  getOptions(eventName: string, noAck = false): RmqOptions {
    return {
      transport: Transport.RMQ,
      options: {
        urls: [this.getRabbitMQUri()],
        queue: getQueueName(eventName),
        noAck,
        wildcards: true,
        exchange: 'events',
        exchangeType: 'topic',
        persistent: true,
        routingKey: getRoutingKey(eventName),
        queueOptions: {
          durable: true,
          deadLetterExchange: getDLXName('events'), // Dead letter exchange name
          deadLetterRoutingKey: getDeadRoutingKey(eventName), // Dead letter routing key
          messageTtl: 60000, // Time to live for messages in the queue (e.g., 60 seconds)
        },
      },
    };
  }

  async ack(context: RmqContext) {
    const channel = context.getChannelRef();
    const originalMsg = context.getMessage();
    await channel.ack(originalMsg);
  }

  async nack(context: RmqContext, requeue = false) {
    const channel = context.getChannelRef();
    const originalMsg = context.getMessage();
    await channel.nack(originalMsg, false, requeue);
  }

  // Publish an raw event object to RabbitMQ
  async publish(eventName: string, data: any) {
    const channel = this.rmq
      .unwrap<AmqpConnectionManager>()
      .createChannel({ json: true });

    if (typeof data === 'string') {
      data = JSON.parse(data);
    }
    const routingKey = getRoutingKey(eventName);
    await channel.publish('events', routingKey, {
      ...keysToSnakeCase(data), // for Go compatability
      pattern: routingKey,
    });
    channel.close();
  }

  // Emit an event to RabbitMQ using the ClientProxy
  emit(pattern: string, data: any) {
    this.rmq.emit(pattern, data);
  }

  async declareQueue(options: RmqOptions['options']) {
    const queueName = options?.queue,
      exchangeName = options?.exchange,
      routingKey = options?.routingKey,
      exchangeType = options?.exchangeType;
    if (!queueName || !exchangeName || !routingKey || !exchangeType) {
      throw new Error(
        'Missing required parameters for declaring queue, exchange, or routing key.',
      );
    }

    await this.rmq.connect();
    const channel = this.rmq.unwrap<AmqpConnectionManager>().createChannel();
    await channel.assertQueue(queueName, {
      ...options?.queueOptions,
    });
    await channel.assertExchange(exchangeName, exchangeType, {
      durable: true,
    });
    await channel.bindQueue(queueName, exchangeName, routingKey);
    await channel.close();
  }
}
