import { Module } from '@nestjs/common';
import { ScheduleModule } from '@nestjs/schedule';

import { CommonModule } from '@phanhotboy/nsv-common';
import { configuration } from './config/configuration';
import { AuthModule } from './modules/auth';
import { PrismaModule } from './database';
import { SubscriptionModule } from './modules/subscriptions/subscription.module';
import { OutboxEventModule } from './modules/outboxEvents/outboxEvent.module';

@Module({
  imports: [
    CommonModule.forRoot({
      configuration,
      cachePrefix: 'auth-service',
      global: true,
    }),
    PrismaModule.forRoot(),
    ScheduleModule.forRoot(),
    AuthModule,
    SubscriptionModule,
    OutboxEventModule,
  ],
})
export class AppModule {}
