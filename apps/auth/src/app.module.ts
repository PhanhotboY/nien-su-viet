import { Module } from '@nestjs/common';
import { CommonModule } from '@phanhotboy/nsv-common';

import { configuration } from './config/configuration';
import { AuthModule } from './auth';
import { PrismaModule } from './database';
import { SubscriptionConsumer } from './infrastructure/messaging/subscription.consumer';

@Module({
  imports: [
    CommonModule.forRoot({
      configuration,
      cachePrefix: 'auth-service',
      global: true,
    }),
    PrismaModule.forRoot(),
    AuthModule,
  ],
  providers: [SubscriptionConsumer],
})
export class AppModule {}
