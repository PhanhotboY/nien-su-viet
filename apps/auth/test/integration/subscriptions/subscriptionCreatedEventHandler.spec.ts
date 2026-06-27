import { describe, it, expect, beforeEach, beforeAll, afterAll } from 'vitest';
import { INestApplication } from '@nestjs/common';
import { randomUUID } from 'node:crypto';

import { AuthService } from '@auth/modules/auth';
import { getRoutingKey } from '@phanhotboy/nsv-common';
import { SubscriptionCreatedEvent } from '@auth/events/subscriptionCreated.event';
import { SubscriptionStatus } from '@phanhotboy/genproto/billing_service/subscriptions';
import { SubscriptionHandler } from '@auth/modules/subscriptions/subscription.handler';
import { AuthHelper } from '../helper/auth.helper';
import { OutboxEventService } from '@auth/modules/outboxEvents/outboxEvent.service';
import { ROLES } from '@phanhotboy/nsv-common/lib';
import { ProcessedEventService } from '@auth/modules/processedEvents/processEvent.service';
import { OutboxEventStatus } from '@auth-prisma';
import { InfrastructureHelper } from '../helper/infrastructure.helper';
import { createTestingAppModule } from '../helper/app.helper';
import { UserWithRole } from 'better-auth/plugins';

describe('SubscriptionCreatedEventHandler', () => {
  let app: INestApplication;
  let infraHelper: InfrastructureHelper;
  let user: UserWithRole;
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

  beforeAll(async () => {
    infraHelper = new InfrastructureHelper();
    await infraHelper.startInfrastructure();

    app = await createTestingAppModule(infraHelper);

    const authHelper = app.get(AuthHelper);

    await authHelper.registerUser(adminCredentials);
    user = (await authHelper.registerUser(credentials)).user;
    const authHeaders = await authHelper.getAuthHeaders(adminCredentials);

    const authService = app.get(AuthService);
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
  }, 30000); // Increase timeout for container startup

  afterAll(async () => {
    await infraHelper.stopInfrastructure();
  });

  it('should successfully handle SubscriptionCreatedEvent', async () => {
    const authHelper = app.get(AuthHelper);
    const authService = app.get(AuthService);
    const outboxService = app.get(OutboxEventService);
    const processedEventsService = app.get(ProcessedEventService);
    const authHeaders = await authHelper.getAuthHeaders(adminCredentials);

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
