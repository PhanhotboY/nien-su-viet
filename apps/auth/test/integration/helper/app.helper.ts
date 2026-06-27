import { configuration } from '@auth/config';
import { PrismaModule } from '@auth/database';
import { AuthModule } from '@auth/modules/auth';
import { ProcessedEventModule } from '@auth/modules/processedEvents/processedEvent.module';
import { SubscriptionModule } from '@auth/modules/subscriptions/subscription.module';
import { ScheduleModule } from '@nestjs/schedule';
import { Test, TestingModule } from '@nestjs/testing';
import { CommonModule, getDLQName, RmqService } from '@phanhotboy/nsv-common';
import { AuthHelper } from './auth.helper';
import { OutboxEventService } from '@auth/modules/outboxEvents/outboxEvent.service';
import { InfrastructureHelper } from './infrastructure.helper';
import { SubscriptionCreatedEvent } from '@auth/events/subscriptionCreated.event';

export async function createTestingAppModule(
  infraHelper: InfrastructureHelper,
) {
  const dbUrl = infraHelper.getPgConnectionStr();
  const moduleFixture: TestingModule = await Test.createTestingModule({
    imports: [
      CommonModule.forRoot({
        configuration: () =>
          ({
            ...configuration('apps/auth/.env.test')(),
            db: {
              directUrl: dbUrl,
              url: dbUrl,
            },
            redis: {
              url: infraHelper.getRedisConnectionStr(),
            },
            rabbitmq: infraHelper.getRmqConnectionStr(),
          }) as ReturnType<ReturnType<typeof configuration>>,
        cachePrefix: 'auth-service',
        global: true,
      }),
      PrismaModule.forRoot(),
      ScheduleModule.forRoot(),
      AuthModule,
      SubscriptionModule,
      ProcessedEventModule,
    ],
    providers: [AuthHelper, OutboxEventService],
  }).compile();

  const app = moduleFixture.createNestApplication();

  const rmqService = app.get(RmqService);
  [SubscriptionCreatedEvent].forEach((event) => {
    const rmqOptions = rmqService.getOptions(event.name);
    app.connectMicroservice(rmqOptions);
    // Declare the dead letter exchange and queue for the event
    rmqService.declareQueue({
      ...rmqOptions.options,
      exchange: rmqOptions.options?.queueOptions?.deadLetterExchange!,
      queue: getDLQName(rmqOptions.options?.queue!),
      routingKey: rmqOptions.options?.queueOptions?.deadLetterRoutingKey!,
      queueOptions: {
        durable: true,
      },
    });
  });
  app.setGlobalPrefix('/api/v1');

  const port = process.env.NODE_PORT || 3000;

  app.enableShutdownHooks();
  await app.startAllMicroservices();
  await app.listen(port);

  return app;
}
