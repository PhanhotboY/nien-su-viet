import { describe, it, expect, beforeEach, beforeAll, afterAll } from 'vitest';
import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import { ScheduleModule } from '@nestjs/schedule';
import { randomUUID } from 'node:crypto';
import { GenericContainer, StartedTestContainer } from 'testcontainers';
import { execSync } from 'node:child_process';

import { AuthModule, AuthService } from '@auth/modules/auth';
import {
  CommonModule,
  ConfigService,
  getDLQName,
  getRoutingKey,
  RmqService,
} from '@phanhotboy/nsv-common';
import { PrismaModule } from '@auth/database';
import { SubscriptionModule } from '@auth/modules/subscriptions/subscription.module';
import { SubscriptionCreatedEvent } from '@auth/events/subscriptionCreated.event';
import { configuration } from '@auth/config/configuration';
import { SubscriptionStatus } from '@phanhotboy/genproto/billing_service/subscriptions';
import { SubscriptionHandler } from '@auth/modules/subscriptions/subscription.handler';
import { Config } from '@auth/config';
import { AuthHelper } from '../helper/auth.helper';
import { OutboxEventService } from '@auth/modules/outboxEvents/outboxEvent.service';
import { Client } from 'pg';
import { ROLES } from '@phanhotboy/nsv-common/lib';
import { ProcessedEventService } from '@auth/modules/processedEvents/processEvent.service';
import { OutboxEventStatus } from '@auth-prisma';
import { ProcessedEventModule } from '@auth/modules/processedEvents/processedEvent.module';

describe('SubscriptionCreatedEventHandler', () => {
  let app: INestApplication;
  let redisContainer: StartedTestContainer | undefined;
  let rmqContainer: StartedTestContainer | undefined;
  let pgContainer: StartedTestContainer | undefined;

  beforeAll(async () => {
    redisContainer = await new GenericContainer('redis:alpine')
      .withExposedPorts(6379)
      .start();

    rmqContainer = await new GenericContainer('rabbitmq:4-management-alpine')
      .withExposedPorts(5672, 15672)
      .start();

    pgContainer = await new GenericContainer('postgres:alpine')
      .withEnvironment({
        POSTGRES_USER: 'testuser',
        POSTGRES_PASSWORD: 'testpassword',
        POSTGRES_DB: 'testdb',
      })
      .withExposedPorts(5432)
      .start();

    const dbUrl = `postgresql://testuser:testpassword@${pgContainer?.getHost()}:${pgContainer?.getMappedPort(5432)}/testdb`;
    execSync('pnpm prisma migrate deploy', {
      stdio: 'inherit',
      env: {
        ...process.env,
        DATABASE_URL: dbUrl,
      },
    });

    const client = new Client({ connectionString: dbUrl });
    await client.connect();
    const result = await client.query(`
      SELECT tablename, schemaname FROM pg_catalog.pg_tables WHERE schemaname NOT IN ('pg_catalog', 'information_schema');
    `);
    console.log('Created tables: ', result.rows);
    await client.end();

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
                url: `redis://${redisContainer?.getHost()}:${redisContainer?.getMappedPort(6379)}`,
              },
              rmq: {
                urls: [
                  `amqp://guest:guest@${rmqContainer?.getHost()}:${rmqContainer?.getMappedPort(
                    5672,
                  )}`,
                ],
              },
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

    app = moduleFixture.createNestApplication();

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
  }, 30000); // Increase timeout for container startup

  afterAll(async () => {
    await Promise.all([
      redisContainer?.stop(),
      rmqContainer?.stop(),
      pgContainer?.stop(),
    ]);
  });

  it('should successfully handle SubscriptionCreatedEvent', async () => {
    const authService = app.get(AuthService);
    const outboxService = app.get(OutboxEventService);
    const processedEventsService = app.get(ProcessedEventService);

    const authHelper = app.get(AuthHelper);
    const adminCredentials = {
      email: 'admin@example.com',
      password: 'adminpassword',
      role: ROLES.ADMIN,
    };
    const credentials = {
      email: 'test@example.com',
      password: 'password123',
      role: ROLES.USER,
    };

    const { user: admin } = await authHelper.registerUser(adminCredentials);
    const { user } = await authHelper.registerUser(credentials);
    const authHeaders = await authHelper.getAuthHeaders(adminCredentials);

    await Promise.all([
      authService.api.createOrganization({
        body: {
          name: 'Test Organization',
          slug: 'premium',
        },
        headers: authHeaders,
      }),
      authService.api.createOrganization({
        body: {
          name: 'Default Organization',
          slug: 'default',
        },
        headers: authHeaders,
      }),
    ]);

    const subscriptionCreatedEvent = new SubscriptionCreatedEvent({
      eventId: randomUUID(),
      pattern: getRoutingKey(SubscriptionCreatedEvent.name),
      occurredAt: new Date(),
      data: {
        subscriptionId: randomUUID(),
        userId: user.id,
        planId: randomUUID(),
        status: SubscriptionStatus.SUBSCRIPTION_STATUS_PENDING,
        currentPeriodStart: new Date(),
        currentPeriodEnd: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // 30 days later
        cancelAtPeriodEnd: false,
      },
    });

    const subscriptionHandler = app.get(SubscriptionHandler);
    await subscriptionHandler.handleSubscriptionCreatedEvent(
      subscriptionCreatedEvent,
    );

    const { members: premiumMembers } = await authService.api.listMembers({
      query: {
        organizationSlug: 'premium',
      },
      headers: authHeaders,
    });

    expect(premiumMembers.some((member) => member.userId === user.id)).toBe(
      true,
    );
    expect(
      premiumMembers.find((member) => member.userId === user.id)?.role,
    ).toBe(ROLES.ORGANIZATION.MEMBER);

    const outboxEvents = await outboxService.findByEventType(
      getRoutingKey(SubscriptionCreatedEvent.name),
    );
    expect(
      outboxEvents.find((event) => event.status === OutboxEventStatus.PENDING),
    ).toBeDefined();

    const processedEvent = await processedEventsService.isEventProcessed(
      subscriptionCreatedEvent.eventId,
    );
    expect(processedEvent).toBe(true);
  });
});
