import { Module } from '@nestjs/common';
import { SubscriptionConsumer } from './subscription.consumer';
import { SubscriptionHandler } from './subscription.handler';
import { RmqModule } from '@phanhotboy/nsv-common';
import { AuthModule } from '../auth';

@Module({
  imports: [RmqModule.register(), AuthModule],
  controllers: [SubscriptionConsumer],
  providers: [SubscriptionHandler],
})
export class SubscriptionModule {}
